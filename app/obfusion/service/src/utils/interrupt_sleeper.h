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
// Created by Homin Su on 2022/5/26.
//

#ifndef BUGU_OBFUSION_SERVICE_SRC_UTILS_INTERRUPT_SLEEPER_H_
#define BUGU_OBFUSION_SERVICE_SRC_UTILS_INTERRUPT_SLEEPER_H_

#include <atomic>
#include <chrono>
#include <memory>
#include <mutex>

namespace bugu {

class InterruptSleeper {
 private:
  ::std::condition_variable cv_;
  ::std::mutex mutex_;
  ::std::atomic<bool> terminate_ = false;

 public:
  void wait() {
    ::std::unique_lock<::std::mutex> lock(mutex_);
    cv_.wait(lock, [&] { return terminate_.load(); });
  }

  template<typename R, typename P>
  bool wait_for(::std::chrono::duration<R, P> const &_time) {
    ::std::unique_lock<::std::mutex> lock(mutex_);
    return !cv_.wait_for(lock, _time, [&] { return terminate_.load(); });
  }

  void interrupt() {
    ::std::unique_lock<::std::mutex> lock(mutex_);
    terminate_.store(true);
    cv_.notify_all();
  }
};

} // namespace bugu

#endif //BUGU_OBFUSION_SERVICE_SRC_UTILS_INTERRUPT_SLEEPER_H_
