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

#ifndef BUGU_OBFUSION_SERVICE_SRC_THREAD_POOL_X_THREAD_POOL_H_
#define BUGU_OBFUSION_SERVICE_SRC_THREAD_POOL_X_THREAD_POOL_H_

#include <atomic>
#include <condition_variable>
#include <mutex>
#include <queue>
#include <shared_mutex>
#include <thread>
#include <vector>

namespace bugu {
class XTaskBase;

/**
 * @brief 线程池
 */
class XThreadPool {
 private:
  static ::std::once_flag flag_;  ///< 函数单次执行标识
  ::std::size_t thread_nums_ = 0; ///< 线程数量
  ::std::vector<::std::unique_ptr<::std::thread>> threads_;  ///< 线程池线程
  ::std::queue<::std::shared_ptr<XTaskBase>> x_tasks_;   ///< 任务队列
  ::std::atomic<bool> is_running_ = false;      ///< 线程池运行状态
  ::std::atomic<int> task_run_count_ = 0; ///< 正在运行的任务数量，原子变量，线程安全

  mutable ::std::shared_mutex mutex_;
  ::std::condition_variable_any cv_;

 public:
  ~XThreadPool();

  // 禁止生成拷移动构造函数, 拷贝构造函数
  XThreadPool(XThreadPool &&_pool) = delete;
  XThreadPool(const XThreadPool &_pool) = delete;
  XThreadPool &operator=(const XThreadPool &_pool) = delete;

  /**
   * @brief 单件模式
   * @return XThreadPool*
   */
  static XThreadPool *Get() {
    static XThreadPool p;
    return &p;
  }

 private:
  XThreadPool() = default;

 public:
  /**
   * @brief 初始化所有线程，并启动线程
   */
  ::std::size_t Init(::std::size_t _thread_nums);

  /**
   * @brief 线程池退出
   */
  void Stop();

  /**
   * @brief 插入任务
   * @param _x_task 任务指针
   */
  void AddTask(::std::shared_ptr<XTaskBase> &&_x_task);

 private:
  /**
   * @brief 线程池线程的入口函数
   */
  void Run();

  /**
   * @brief 获取任务指针
   * @return XTaskBase* 任务指针
   */
  std::shared_ptr<XTaskBase> GetTask();

  // 获取器和设置器
 public:
  /**
   * 获取线程池运行状态
   * @return bool 线程池运行状态
   */
  bool is_running() const;

  /**
   * 获取线程池当前执行中的任务数量
   * @return int 任务数量
   */
  int task_run_count();
};
} // namespace bugu

#endif //BUGU_OBFUSION_SERVICE_SRC_THREAD_POOL_X_THREAD_POOL_H_
