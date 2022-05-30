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
// Created by Homin Su on 2022/5/30.
//

#ifndef BUGU_DETECT_SERVICE_SRC_BUGU_DETECT_SERVER_H_
#define BUGU_DETECT_SERVICE_SRC_BUGU_DETECT_SERVER_H_

#include "bugu_dectet/bugu_detect.h"
#include "thread/x_thread.h"

#include <grpc++/grpc++.h>

#include <condition_variable>
#include <memory>
#include <memory_resource>
#include <mutex>
#include <utility>

namespace bugu {

class XThreadPool;

class BuguDetectServer final : public XThread {
 private:
  bool init_flag_ = false;

  ::std::string local_address_{};  ///< rpc 服务地址：ip+端口
  ::std::unique_ptr<::grpc::Server> server_{};  ///< rpc 服务句柄，用智能指针管理

  ::bugu::XThreadPool *thread_pool_ = nullptr;  ///< 线程池
  ::std::shared_ptr<::std::pmr::memory_resource> memory_resource_;  ///< 内存池

  ::std::shared_mutex mutex_;

 public:
  ~BuguDetectServer();
  static BuguDetectServer *Get() {
    static BuguDetectServer s;
    return &s;
  }

 private:
  BuguDetectServer() = default;

 public:
  BuguDetectServer *Init(::std::string _local_address,
                         ::bugu::XThreadPool *_thread_pool,
                         ::std::shared_ptr<::std::pmr::memory_resource> _memory_resource);

  void Start() override;

  void Main() override;

  void Stop() override;

  void set_init_flag(bool _init_flag);

  bool init_flag();
};

} // namespace bugu

#endif //BUGU_DETECT_SERVICE_SRC_BUGU_DETECT_SERVER_H_
