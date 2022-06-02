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

#ifndef BUGU_DETECT_SERVICE_RESNET_INCLUDE_RESNET_RESNET_18_H_
#define BUGU_DETECT_SERVICE_RESNET_INCLUDE_RESNET_RESNET_18_H_

#include "g_logging.h"

#include <map>
#include <memory>
#include <string>

namespace bugu::Resnet {

static const int INPUT_H = 224;
static const int INPUT_W = 224;
static const int OUTPUT_SIZE = 1000;
static const int BATCH_SIZE = 1;

inline namespace v18 {

class Resnet {
 public:
  static const int INPUT_H = ::bugu::Resnet::INPUT_H;
  static const int INPUT_W = ::bugu::Resnet::INPUT_W;
  static const int OUTPUT_SIZE = ::bugu::Resnet::OUTPUT_SIZE;

  const char *INPUT_BLOB_NAME = "data";
  const char *OUTPUT_BLOB_NAME = "prob";

 private:
  bool is_prepared_ = false;
  nvinfer1::IRuntime *runtime_{};
  nvinfer1::ICudaEngine *engine_{};
  nvinfer1::IExecutionContext *context_{};
  int input_index_{};
  int output_index_{};
  void *gpu_buffers_[2]{};
  cudaStream_t stream_{};

  Logger g_logger_;

  ::std::string engine_name_ = "resnet18.engine";
  long model_size_{};
  ::std::shared_ptr<char[]> trtModelStream_{};

 public:
  ~Resnet();
  static Resnet *Get() {
    static Resnet r;
    return &r;
  }
 private:
  Resnet() = default;

 public:
  void Init();
  void GenEngine();
  void LoadEngine();
  void Inference(float *_input, float *_output);

 private:
  void APIToModel(unsigned int _max_batch_size, IHostMemory **_model_stream);
  void BuildEngine(unsigned int _max_batch_size,
                   IBuilder *_builder,
                   IBuilderConfig *_config,
                   DataType _data_type);
  ::std::map<::std::string, Weights> LoadWeights(const ::std::string _file_name);
  IScaleLayer *AddBatchNorm2d(INetworkDefinition *_network,
                              ::std::map<::std::string, Weights> &_weight_map,
                              ITensor &_input,
                              ::std::string _layer_name,
                              float _eps);
  IActivationLayer *BasicBlock(INetworkDefinition *_network,
                               ::std::map<::std::string, Weights> &_weight_map,
                               ITensor &_input,
                               int _in_ch,
                               int _out_ch,
                               int _stride,
                               ::std::string _layer_name);

};

} // inline namespace v18

} // namespace bugu::Resnet

#endif //BUGU_DETECT_SERVICE_RESNET_INCLUDE_RESNET_RESNET_18_H_
