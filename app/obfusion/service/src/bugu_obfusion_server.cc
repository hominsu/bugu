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
// Created by Homin Su on 2022/5/19.
//

#include "bugu_obfusion_server.h"
#include "bugu_obfusion_impl.h"
#include "thread_pool/x_thread_pool.h"
//#include "utils/credentials.h"

#if BUGU_DEBUG
#include <string>
#endif
#include <utility>

bugu::BuguObfusionServer::~BuguObfusionServer() {
  if (IsRunning()) { Stop(); }
}

bugu::BuguObfusionServer *bugu::BuguObfusionServer::Init(::std::string _local_address,
                                                         ::bugu::XThreadPool *_thread_pool,
                                                         ::std::shared_ptr<::std::pmr::memory_resource> _memory_resource) {
  // The thread pool and Grpc service are initialized only once
  if (!init_flag()) {
    local_address_ = ::std::move(_local_address);
    thread_pool_ = _thread_pool;
    memory_resource_ = ::std::move(_memory_resource);
    set_init_flag(true);
  }

  return this;
}

void bugu::BuguObfusionServer::Start() {
  BUGU_ASSERT(init_flag() && "BuguObfusionServer should init first");
  XThread::Start();
}

void bugu::BuguObfusionServer::Stop() {
  ::std::function<void()> func = [this]() {
    server_->Shutdown();
  };
  StopWith(func);
}

void bugu::BuguObfusionServer::Main() {
  BuguObfusionImpl bugu_obfusion_service(thread_pool_, memory_resource_);

  ::grpc::ServerBuilder builder;
  builder.AddListeningPort(local_address_, ::grpc::InsecureServerCredentials());
  builder.RegisterService(&bugu_obfusion_service);

  ::std::unique_ptr<::grpc::Server> server(builder.BuildAndStart());
  server_ = ::std::move(server);

#if BUGU_DEBUG
  ::std::cout << "Server listening on " << local_address_ << ::std::endl;
#endif

  server_->Wait();

#if BUGU_DEBUG
  ::std::cout << "Rpc Service Shut Down" << ::std::endl;
#endif
}

void bugu::BuguObfusionServer::set_init_flag(bool _init_flag) {
  ::std::unique_lock<::std::shared_mutex> lock(mutex_);
  init_flag_ = _init_flag;
}

bool bugu::BuguObfusionServer::init_flag() {
  ::std::shared_lock<::std::shared_mutex> lock(mutex_);
  return init_flag_;
}

