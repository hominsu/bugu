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

#include "x_thread.h"

/**
 * @brief 线程开始函数
 * @details 该函数是多线程的开始函数，会将 Main 函数放入一个线程中运行
 */
void bugu::XThread::Start() {
  if (!this->IsRunning()) {
    this->SetIsRunning(true);
    this->thread_ = ::std::thread(&XThread::Main, this);
  }
}

/**
 * @brief 等待线程完成
 */
void bugu::XThread::Wait() {
  if (this->thread_.joinable()) {
    this->thread_.join();
  }
}

/**
 * @brief 停止线程
 */
void bugu::XThread::Stop() {
  if (this->IsRunning()) {
    this->SetIsRunning(false);
  }
  Wait();
}

void bugu::XThread::StopWith(::std::function<void()> &_do) {
  _do();
  Stop();
}

/**
 * @brief 休眠该线程若干毫秒
 * @details 接收一个 ::std::chrono::milliseconds 的时间戳，调用 ::std::this_thread::sleep_for() 休眠该线程指定时间
 * @param _time ::std::chrono::milliseconds 时间戳
 */
void bugu::XThread::ThreadSleep(::std::chrono::milliseconds _time) {
  ::std::this_thread::sleep_for(_time);
}

/**
 * @brief 获取 isRunning_ 状态
 * @return bool 返回值为 true，说明线程当前处于运行状态
 */
bool bugu::XThread::IsRunning() const {
  ::std::shared_lock<::std::shared_mutex> lock(is_running_mutex_);
  return is_running_;
}

/**
 * @brief 设置线程运行状态
 * @param is_running 运行状态
 */
void bugu::XThread::SetIsRunning(bool is_running) {
  ::std::unique_lock<::std::shared_mutex> lock(is_running_mutex_);
  is_running_ = is_running;
}
