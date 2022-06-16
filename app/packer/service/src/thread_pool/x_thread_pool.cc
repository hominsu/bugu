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
#include "x_thread_pool.h"
#include "x_task.h"

#include <iostream>
#include <sstream>
#include <thread>

bugu::XThreadPool::~XThreadPool() {
  // 检查是否停止
  if (is_running_) {
    Stop();
  }
}

/**
 * @brief 初始化所有线程，并启动线程
 * @param _thread_nums 线程数量
 */
::std::size_t bugu::XThreadPool::Init(::std::size_t _thread_nums) {
  ::std::unique_lock<::std::shared_mutex> lock(mutex_);

  thread_nums_ = _thread_nums;

  BUGU_ASSERT(thread_nums_ > 0 && "thread_nums_ <= 0");
  if (thread_nums_ <= 0) {
    throw ::std::runtime_error("thread_nums_ <= 0");
  }

  BUGU_ASSERT(threads_.empty() && "threads should be empty");
  if (!threads_.empty()) {
    throw ::std::runtime_error("!threads_.empty()");
  }

  // 创建线程对象
  for (::std::size_t i = 0; i < thread_nums_; ++i) {
    threads_.push_back(::std::make_unique<::std::thread>(&bugu::XThreadPool::Run, this));
  }

  // 设置线程池运行状态
  is_running_ = true;

  return thread_nums_;
}

/**
 * @brief 线程池退出
 */
void bugu::XThreadPool::Stop() {
  // 设置退出状态
  is_running_ = false;

  // 通知全部线程
  cv_.notify_all();

  // 等待线程结束任务退出
  for (auto &th: threads_) {
    th->join();
  }

  // 独占锁
  ::std::unique_lock<::std::shared_mutex> lock(mutex_);

  // 清理线程池中的线程对象
  threads_.clear();
}

/**
 * @brief 线程池线程的入口函数
 */
void bugu::XThreadPool::Run() {
#ifdef BUGU_DEBUG
  ::std::stringstream str_info;
  str_info << "Run: " << ::std::this_thread::get_id() << ::std::endl;
  ::std::cout << str_info.str();
#endif

  while (is_running_) {
    // 获取任务
    auto task = GetTask();
    // 获取到空指针, continue之后然后继续获取或退出线程
    if (nullptr == task) {
      continue;
    }

    ++task_run_count_;  // 设置运行中的任务个数
    try {
      // 执行任务
      task->Main();
    } catch (::std::exception &e) {
      ::std::stringstream str_e;
      str_e << "Failure in thread " << ::std::this_thread::get_id() << ", Exception: " << e.what() << ::std::endl;
      ::std::cerr << str_e.str();
    } catch (...) {
      ::std::stringstream str_e;
      str_e << "Unknown failure in thread " << ::std::this_thread::get_id() << ::std::endl;
      ::std::cerr << str_e.str();
    }
    --task_run_count_;
#ifdef BUGU_DEBUG
    ::std::cout << "run: " << task_run_count_ << ::std::endl;
#endif
  }
}

/**
 * @brief 插入任务
 * @param _x_task 任务指针
 */
void bugu::XThreadPool::AddTask(::std::shared_ptr<XTaskBase> &&_x_task) {
  // 将任务插入到队列
  {
    // 独占锁
    ::std::unique_lock<::std::shared_mutex> lock(mutex_);
    // 将线程池运行状态函数的函数指针传入任务中
    _x_task->is_running = [this] {
      return is_running();
    };
    x_tasks_.push(_x_task);
  }

  // 通知一个线程取任务
  cv_.notify_one();
}

/**
 * @brief 获取任务指针
 * @return XTaskBase* 任务指针
 */
::std::shared_ptr<bugu::XTaskBase> bugu::XThreadPool::GetTask() {
  // 独占锁，防止抢占
  ::std::unique_lock<::std::shared_mutex> lock(mutex_);

  // 当任务队列为空就阻塞
  if (x_tasks_.empty()) {
    cv_.wait(lock);
  }

  // 退出
  if (!is_running_) {
    return nullptr;
  }

  // 防止多次通知
  if (x_tasks_.empty()) {
    return nullptr;
  }

  // 取出队头任务
  auto task = x_tasks_.front();
  x_tasks_.pop();
  return task;
}

/**
 * 获取线程池运行状态
 * @return bool 线程池运行状态
 */
bool bugu::XThreadPool::is_running() const {
  return is_running_;
}

/**
 * 获取线程池当前执行中的任务数量
 * @return int 任务数量
 */
int bugu::XThreadPool::task_run_count() {
  return task_run_count_;
}

