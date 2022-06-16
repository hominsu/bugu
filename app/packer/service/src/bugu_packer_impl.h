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
// Created by Homin Su on 2022/6/16.
//

#ifndef BUGU_PACKER_SERVICE_SRC_BUGU_PACKER_IMPL_H_
#define BUGU_PACKER_SERVICE_SRC_BUGU_PACKER_IMPL_H_

#include "api/packer/service/v1/cpp/bugu_packer.grpc.pb.h"

#include <grpc++/grpc++.h>
#include <memory_resource>
#include <memory>
#include <thread>

namespace bugu {

class XThreadPool;

class BuguPackerImpl final : public bugu_packer::service::v1::BuguPacker::Service {
 private:
  ::bugu::XThreadPool *thread_pool_;  ///< 线程池
  ::std::shared_ptr<::std::pmr::memory_resource> memory_resource_;  ///< 内存池

 public:
  explicit BuguPackerImpl(::bugu::XThreadPool *_thread_pool,
                            ::std::shared_ptr<::std::pmr::memory_resource> _memory_resource)
      : thread_pool_(_thread_pool), memory_resource_(::std::move(_memory_resource)) {};

  ::grpc::Status Packer(::grpc::ServerContext *_ctx,
                          const ::bugu_packer::service::v1::PackerRequest *_request,
                          ::bugu_packer::service::v1::PackerReply *_response) override;
};

} // namespace bugu

#endif //BUGU_PACKER_SERVICE_SRC_BUGU_PACKER_IMPL_H_
