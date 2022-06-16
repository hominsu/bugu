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

#include "bugu_packer/bugu_packer.h"
#include "bugu_packer_impl.h"
#include "data.h"
#include "packer_task.h"
#include "thread_pool/x_thread_pool.h"

#include <cstring>
#include <string>

::grpc::Status bugu::BuguPackerImpl::Packer(::grpc::ServerContext *_ctx,
                                            const ::bugu_packer::service::v1::PackerRequest *_request,
                                            ::bugu_packer::service::v1::PackerReply *_response) {
  auto size = _request->size();

  auto data = ::bugu::Data::Make(memory_resource_);
  auto buffer = static_cast<char *>(data->New(size));

  ::std::size_t read_bytes;
  ::std::size_t total_read_bytes = 0;

  // Copy the data from request to buffer
  for (auto &d: _request->data()) {
    read_bytes = d.size();
    memcpy(buffer + total_read_bytes, d.data(), read_bytes);
    total_read_bytes += read_bytes;
  }
  BUGU_ASSERT(total_read_bytes == size && "total_read_bytes != size");

  // create the packer task and append to the thread pool, block to wait return
  auto packer_task = ::std::make_shared<PackerTask>(data, memory_resource_);
  thread_pool_->AddTask(packer_task);
  auto ret = packer_task->get_return();

  _response->add_data(ret->data(), ret->size());
  _response->set_size(ret->size());

//  return {::grpc::StatusCode::UNIMPLEMENTED, "method Packer not implemented"};
  return ::grpc::Status::OK;
}
