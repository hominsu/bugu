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
// Created by Homin Su on 2022/5/17.
//

#ifndef BUGU_OBFUSION_SERVICE_SRC_THREAD_X_THREAD_H_
#define BUGU_OBFUSION_SERVICE_SRC_THREAD_X_THREAD_H_

#include <mutex>
#include <thread>
#include <shared_mutex>
#include <functional>

namespace bugu {
/**
 * @brief 线程基类
 * @details Start() 启动服务，Stop() 关闭服务
 */
class XThread {
 private:
  ::std::thread thread_;  ///< 线程句柄
  bool is_running_ = false;  ///< 当前线程运行状态
  mutable ::std::shared_mutex is_running_mutex_;  ///< 线程运行状态互斥量

 public:
  virtual void Start();
  virtual void Wait();
  virtual void Stop();
  virtual void StopWith(::std::function<void()> &_do);
  virtual void ThreadSleep(::std::chrono::milliseconds _time);

  bool IsRunning() const;

 private:
  void SetIsRunning(bool is_running);

  /**
   * @brief 该纯虚函数必须在子类中实现，用于线程函数的主函数
   */
  virtual void Main() = 0;
};

} // namespace bugu

#endif //BUGU_OBFUSION_SERVICE_SRC_THREAD_X_THREAD_H_
