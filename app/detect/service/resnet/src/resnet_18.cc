// MIT License
//
// Copyright (c) 2022. HominSu
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//

//
// Created by Homin Su on 2022/5/31.
//

#include "bugu_dectet/bugu_detect.h"
#include "resnet/resnet_18.h"

#include "NvInfer.h"
#include "cuda_runtime_api.h"

#include <cmath>

#include <fstream>
#include <iostream>
#include <map>
#include <memory>
#include <sstream>
#include <exception>
#include <vector>
#include <chrono>

#define CHECK(status) \
    do {\
        auto ret = (status);\
        if (ret != 0)\
        {\
            ::std::clog << "Cuda failure: " << ret << ::std::endl;\
            abort();\
        }\
    } while (0)

using namespace nvinfer1;

bugu::Resnet::v18::Resnet::~Resnet() {
  if (is_prepared_) {
    // Release stream and gpu_buffers_
    cudaStreamDestroy(stream_);
    CHECK(cudaFree(gpu_buffers_[input_index_]));
    CHECK(cudaFree(gpu_buffers_[output_index_]));

    // Destroy the engine
    if (nullptr != context_) {
      context_->destroy();
      context_ = nullptr;
      std::cout << "context_->destroy();" << std::endl;
    }
    if (nullptr != engine_) {
      engine_->destroy();
      engine_ = nullptr;
      std::cout << "engine_->destroy();" << std::endl;
    }
    if (nullptr != runtime_) {
      runtime_->destroy();
      runtime_ = nullptr;
      std::cout << "runtime_->destroy();" << std::endl;
    }
  }
}

void bugu::Resnet::v18::Resnet::Init() {
  runtime_ = createInferRuntime(g_logger_);
  assert(runtime_ != nullptr);

  engine_ = runtime_->deserializeCudaEngine(trtModelStream_.get(), model_size_, nullptr);
  assert(engine_ != nullptr);

  context_ = engine_->createExecutionContext();
  assert(context_ != nullptr);

  // Pointers to input and output device buffers to pass to engine.
  // Engine requires exactly IEngine::getNbBindings() number of buffers.
  assert(engine_.getNbBindings() == 2);

  // In order to bind the buffers, we need to know the names of the input and output tensors.
  // Note that indices are guaranteed to be less than IEngine::getNbBindings()
  input_index_ = engine_.getBindingIndex(INPUT_BLOB_NAME);
  output_index_ = engine_.getBindingIndex(OUTPUT_BLOB_NAME);
  assert(input_index_ == 0);
  assert(output_index_ == 1);

  // Create GPU buffers on device
  CHECK(cudaMalloc(&gpu_buffers_[input_index_], BATCH_SIZE * 3 * INPUT_H * INPUT_W * sizeof(float)));
  CHECK(cudaMalloc(&gpu_buffers_[output_index_], BATCH_SIZE * OUTPUT_SIZE * sizeof(float)));

  // Check stream
  CHECK(cudaStreamCreate(&stream_));

  is_prepared_ = true;
}

void bugu::Resnet::v18::Resnet::GenEngine() {
  assert(!engine_name_.empty() && "engine name should not be empty");

  // create a model using the API directly and serialize it to a stream
  ::std::clog << "Gen Engine ..." << ::std::endl;

  IHostMemory *model_stream{nullptr};
  APIToModel(1, &model_stream);
  assert(model_stream != nullptr);

  ::std::ofstream p(engine_name_, ::std::ios::binary);
  if (!p) {
    throw ::std::runtime_error("could not open plan output file");
  }
  p.write(reinterpret_cast<const char *>(model_stream->data()), model_stream->size());

  model_stream->destroy();
}

void bugu::Resnet::v18::Resnet::LoadEngine() {
  // deserialize the .engine
  ::std::ifstream file(engine_name_, ::std::ios::binary);
  if (file.good()) {
    // get the file size
    file.seekg(0, ::std::ifstream::end);
    model_size_ = file.tellg();
    file.seekg(0, ::std::ifstream::beg);

    trtModelStream_ = ::std::shared_ptr<char[]>(new char[model_size_]());
    assert(trtModelStream_);

    // load the model data
    file.read(trtModelStream_.get(), model_size_);
    file.close();
  }
}

void bugu::Resnet::v18::Resnet::Inference(float *_input, float *_output) {
  // DMA input batch data to device, infer on the batch asynchronously, and DMA output back to host
  CHECK(cudaMemcpyAsync(gpu_buffers_[inputIndex],
                        _input,
                        BATCH_SIZE * 3 * INPUT_H * INPUT_W * sizeof(float),
                        cudaMemcpyHostToDevice,
                        stream_));

  context_.enqueue(_batch_size, gpu_buffers_, stream_, nullptr);

  CHECK(cudaMemcpyAsync(_output,
                        gpu_buffers_[outputIndex],
                        BATCH_SIZE * OUTPUT_SIZE * sizeof(float),
                        cudaMemcpyDeviceToHost,
                        stream_));

  cudaStreamSynchronize(stream_);
}

void bugu::Resnet::v18::Resnet::APIToModel(unsigned int _max_batch_size, IHostMemory **_model_stream) {
  // Create builder
  IBuilder *builder = createInferBuilder(g_logger_);
  IBuilderConfig *config = builder->createBuilderConfig();

  // Create model to populate the network, then set the outputs and create an engine
  ICudaEngine *engine = BuildEngine(_max_batch_size, builder, config, DataType::kFLOAT);
  assert(engine != nullptr);

  // Serialize the engine
  (*_model_stream) = engine->serialize();

  // Close everything down
  engine->destroy();
  builder->destroy();
  config->destroy();
}

void bugu::Resnet::v18::Resnet::BuildEngine(unsigned int _max_batch_size,
                                            IBuilder *_builder,
                                            IBuilderConfig *_config,
                                            DataType _data_type) {
  INetworkDefinition *network = _builder->createNetworkV2(0U);

  // Create input tensor of shape { 3, INPUT_H, INPUT_W } with name INPUT_BLOB_NAME
  ITensor *data = network->addInput(INPUT_BLOB_NAME, _data_type, Dims3{3, INPUT_H, INPUT_W});
  assert(data);

  ::std::map<::std::string, Weights> weight_map = loadWeights("../resnet18.wts");
  Weights empty_wts{DataType::kFLOAT, nullptr, 0};

  IConvolutionLayer *conv1 = network->addConvolutionNd(*data, 64, DimsHW{7, 7}, weight_map["conv1.weight"], empty_wts);
  assert(conv1);
  conv1->setStrideNd(DimsHW{2, 2});
  conv1->setPaddingNd(DimsHW{3, 3});

  IScaleLayer *bn1 = AddBatchNorm2d(network, weight_map, *conv1->getOutput(0), "bn1", 1e-5);

  IActivationLayer *relu1 = network->addActivation(*bn1->getOutput(0), ActivationType::kRELU);
  assert(relu1);

  IPoolingLayer *pool1 = network->addPoolingNd(*relu1->getOutput(0), PoolingType::kMAX, DimsHW{3, 3});
  assert(pool1);
  pool1->setStrideNd(DimsHW{2, 2});
  pool1->setPaddingNd(DimsHW{1, 1});

  IActivationLayer *relu2 = BasicBlock(network, weight_map, *pool1->getOutput(0), 64, 64, 1, "layer1.0.");
  IActivationLayer *relu3 = BasicBlock(network, weight_map, *relu2->getOutput(0), 64, 64, 1, "layer1.1.");

  IActivationLayer *relu4 = BasicBlock(network, weight_map, *relu3->getOutput(0), 64, 128, 2, "layer2.0.");
  IActivationLayer *relu5 = BasicBlock(network, weight_map, *relu4->getOutput(0), 128, 128, 1, "layer2.1.");

  IActivationLayer *relu6 = BasicBlock(network, weight_map, *relu5->getOutput(0), 128, 256, 2, "layer3.0.");
  IActivationLayer *relu7 = BasicBlock(network, weight_map, *relu6->getOutput(0), 256, 256, 1, "layer3.1.");

  IActivationLayer *relu8 = BasicBlock(network, weight_map, *relu7->getOutput(0), 256, 512, 2, "layer4.0.");
  IActivationLayer *relu9 = BasicBlock(network, weight_map, *relu8->getOutput(0), 512, 512, 1, "layer4.1.");

  IPoolingLayer *pool2 = network->addPoolingNd(*relu9->getOutput(0), PoolingType::kAVERAGE, DimsHW{7, 7});
  assert(pool2);
  pool2->setStrideNd(DimsHW{1, 1});

  IFullyConnectedLayer
      *fc1 = network->addFullyConnected(*pool2->getOutput(0), 1000, weight_map["fc.weight"], weight_map["fc.bias"]);
  assert(fc1);

  fc1->getOutput(0)->setName(OUTPUT_BLOB_NAME);
  ::std::clog << "set name out" << ::std::endl;
  network->markOutput(*fc1->getOutput(0));

  // Build engine
  _builder->setMaxBatchSize(_max_batch_size);
  _config->setMaxWorkspaceSize(1 << 20);
  ICudaEngine *engine = _builder->buildEngineWithConfig(*network, *_config);
  ::std::clog << "build out" << ::std::endl;

  // Don't need the network anymore
  network->destroy();

  // Release host memory
  for (auto &mem: weight_map) {
    free((void *) (mem.second.values));
  }

  return engine;
}

// Load weights from files shared with TensorRT samples.
// TensorRT weight files have a simple space delimited format:
// [type] [size] <data x size in hex>
::std::map<::std::string, Weights> bugu::Resnet::v18::Resnet::LoadWeights(const ::std::string _file_name) {
  ::std::clog << "Loading weights: " << _file_name << ::std::endl;
  ::std::map<::std::string, Weights> weight_map;

  // Open weights file
  ::std::ifstream input(_file_name);
  assert(input.is_open() && "Unable to load weight file.");

  // Read number of weight blobs
  int32_t count;
  input >> count;
  assert(count > 0 && "Invalid weight map file.");

  while (count--) {
    Weights wt{DataType::kFLOAT, nullptr, 0};
    uint32_t size;

    // Read name and type of blob
    ::std::string name;
    input >> name >> ::std::dec >> size;
    wt.type = DataType::kFLOAT;

    // Load blob
    auto *val = reinterpret_cast<uint32_t *>(malloc(sizeof(val) * size));
    for (uint32_t x = 0, y = size; x < y; ++x) {
      input >> ::std::hex >> val[x];
    }
    wt.values = val;

    wt.count = size;
    weight_map[name] = wt;
  }

  return weight_map;
}

IScaleLayer *bugu::Resnet::v18::Resnet::AddBatchNorm2d(INetworkDefinition *_network,
                                                       ::std::map<::std::string, Weights> &_weight_map,
                                                       ITensor &_input,
                                                       ::std::string _layer_name,
                                                       float _eps) {
  float *gamma = (float *) _weight_map[_layer_name + ".weight"].values;
  float *beta = (float *) _weight_map[_layer_name + ".bias"].values;
  float *mean = (float *) _weight_map[_layer_name + ".running_mean"].values;
  float *var = (float *) _weight_map[_layer_name + ".running_var"].values;
  int len = _weight_map[_layer_name + ".running_var"].count;
  ::std::clog << "len " << len << ::std::endl;

  auto scval = reinterpret_cast<float *>(malloc(sizeof(float) * len));
  for (int i = 0; i < len; i++) {
    scval[i] = gamma[i] / ::std::sqrt(var[i] + _eps);
  }
  Weights scale{DataType::kFLOAT, scval, len};

  auto shval = reinterpret_cast<float *>(malloc(sizeof(float) * len));
  for (int i = 0; i < len; i++) {
    shval[i] = beta[i] - mean[i] * gamma[i] / ::std::sqrt(var[i] + _eps);
  }
  Weights shift{DataType::kFLOAT, shval, len};

  auto pval = reinterpret_cast<float *>(malloc(sizeof(float) * len));
  for (int i = 0; i < len; i++) {
    pval[i] = 1.0;
  }
  Weights power{DataType::kFLOAT, pval, len};

  _weight_map[_layer_name + ".scale"] = scale;
  _weight_map[_layer_name + ".shift"] = shift;
  _weight_map[_layer_name + ".power"] = power;
  IScaleLayer *scale_1 = _network->addScale(_input, ScaleMode::kCHANNEL, shift, scale, power);
  assert(scale_1);
  return scale_1;
}

IActivationLayer *bugu::Resnet::v18::Resnet::BasicBlock(INetworkDefinition *_network,
                                                        ::std::map<::std::string, Weights> &_weight_map,
                                                        ITensor &_input,
                                                        int _in_ch,
                                                        int _out_ch,
                                                        int _stride,
                                                        ::std::string _layer_name) {
  Weights empty_wts{DataType::kFLOAT, nullptr, 0};

  IConvolutionLayer *conv1 =
      _network->addConvolutionNd(_input, _out_ch, DimsHW{3, 3}, _weight_map[_layer_name + "conv1.weight"], empty_wts);
  assert(conv1);
  conv1->setStrideNd(DimsHW{_stride, _stride});
  conv1->setPaddingNd(DimsHW{1, 1});

  IScaleLayer *bn1 = AddBatchNorm2d(_network, _weight_map, *conv1->getOutput(0), _layer_name + "bn1", 1e-5);

  IActivationLayer *relu1 = _network->addActivation(*bn1->getOutput(0), ActivationType::kRELU);
  assert(relu1);

  IConvolutionLayer *conv2 = _network->addConvolutionNd(*relu1->getOutput(0),
                                                        _out_ch,
                                                        DimsHW{3, 3},
                                                        _weight_map[_layer_name + "conv2.weight"],
                                                        empty_wts);
  assert(conv2);
  conv2->setPaddingNd(DimsHW{1, 1});

  IScaleLayer *bn2 = AddBatchNorm2d(_network, _weight_map, *conv2->getOutput(0), _layer_name + "bn2", 1e-5);

  IElementWiseLayer *ew1;
  if (_in_ch != _out_ch) {
    IConvolutionLayer *conv3 = _network->addConvolutionNd(_input,
                                                          _out_ch,
                                                          DimsHW{1, 1},
                                                          _weight_map[_layer_name + "downsample.0.weight"],
                                                          empty_wts);
    assert(conv3);
    conv3->setStrideNd(DimsHW{_stride, _stride});

    IScaleLayer *bn3 = AddBatchNorm2d(_network, _weight_map, *conv3->getOutput(0), _layer_name + "downsample.1", 1e-5);

    ew1 = _network->addElementWise(*bn3->getOutput(0), *bn2->getOutput(0), ElementWiseOperation::kSUM);
  } else {
    ew1 = _network->addElementWise(_input, *bn2->getOutput(0), ElementWiseOperation::kSUM);
  }

  IActivationLayer *relu2 = _network->addActivation(*ew1->getOutput(0), ActivationType::kRELU);
  assert(relu2);
  return relu2;
}

int main(int argc, char **argv) {
  if (argc != 2) {
    std::cerr << "arguments not right!" << std::endl;
    std::cerr << "./resnet18 -s   // serialize model to plan file" << std::endl;
    std::cerr << "./resnet18 -d   // deserialize plan file and run inference" << std::endl;
    return -1;
  }

  // create a model using the API directly and serialize it to a stream
  char *trtModelStream{nullptr};
  size_t size{0};

  if (std::string(argv[1]) == "-s") {
    IHostMemory *modelStream{nullptr};
    APIToModel(1, &modelStream);
    assert(modelStream != nullptr);

    std::ofstream p("resnet18.engine", std::ios::binary);
    if (!p) {
      std::cerr << "could not open plan output file" << std::endl;
      return -1;
    }
    p.write(reinterpret_cast<const char *>(modelStream->data()), modelStream->size());
    modelStream->destroy();
    return 1;
  } else if (std::string(argv[1]) == "-d") {
    std::ifstream file("resnet18.engine", std::ios::binary);
    if (file.good()) {
      file.seekg(0, file.end);
      size = file.tellg();
      file.seekg(0, file.beg);
      trtModelStream = new char[size];
      assert(trtModelStream);
      file.read(trtModelStream, size);
      file.close();
    }
  } else {
    return -1;
  }


  // Subtract mean from image
  static float data[3 * INPUT_H * INPUT_W];
  for (int i = 0; i < 3 * INPUT_H * INPUT_W; i++)
    data[i] = 1.0;

  IRuntime *runtime = createInferRuntime(gLogger);
  assert(runtime != nullptr);
  ICudaEngine *engine = runtime->deserializeCudaEngine(trtModelStream, size, nullptr);
  assert(engine != nullptr);
  IExecutionContext *context = engine->createExecutionContext();
  assert(context != nullptr);
  delete[] trtModelStream;

  // Run inference
  static float prob[OUTPUT_SIZE];
  for (int i = 0; i < 100; i++) {
    auto start = std::chrono::system_clock::now();
    doInference(*context, data, prob, 1);
    auto end = std::chrono::system_clock::now();
    std::cout << std::chrono::duration_cast<std::chrono::milliseconds>(end - start).count() << "ms" << std::endl;
  }

  // Destroy the engine
  context->destroy();
  engine->destroy();
  runtime->destroy();

  // Print histogram of the output distribution
  std::cout << "\nOutput:\n\n";
  for (unsigned int i = 0; i < 10; i++) {
    std::cout << prob[i] << ", ";
  }
  std::cout << std::endl;
  for (unsigned int i = 0; i < 10; i++) {
    std::cout << prob[OUTPUT_SIZE - 10 + i] << ", ";
  }
  std::cout << std::endl;

  return 0;
}
