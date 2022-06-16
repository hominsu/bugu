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

#ifndef BUGU_PACKER_SERVICE_SRC_PACKER_TASK_H_
#define BUGU_PACKER_SERVICE_SRC_PACKER_TASK_H_

#include "thread_pool/x_task.h"

#include <fstream>
#include <memory>
#include <memory_resource>
#include <string>
#include <utility>

namespace bugu {

class Data;

class PackerTask : public XTask<::std::shared_ptr<Data>> {
 private:
  ::std::shared_ptr<Data> data_;
  ::std::shared_ptr<::std::pmr::memory_resource> memory_resource_;  ///< 内存池

 public:
  explicit PackerTask(::std::shared_ptr<Data> _data,
                        ::std::shared_ptr<::std::pmr::memory_resource> _memory_resource)
      : data_(std::move(_data)), memory_resource_(std::move(_memory_resource)) {};
  ~PackerTask() override = default;

 private:
  /**
   * 线程入口函数
   */
  void Main() final;
};

} // namespace bugu

#endif //BUGU_PACKER_SERVICE_SRC_PACKER_TASK_H_
