//
// Created by Homin Su on 2022/5/17.
//

#ifndef BUGU_OBFUSION_SERVICE_SRC_THREAD_X_THREAD_H_
#define BUGU_OBFUSION_SERVICE_SRC_THREAD_X_THREAD_H_

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
  std::thread thread_;  ///< 线程句柄
  bool is_running_ = false;  ///< 当前线程运行状态
  mutable std::shared_mutex is_running_mutex_;  ///< 线程运行状态互斥量

 public:
  virtual void Start();
  virtual void Wait();
  virtual void Stop();
  virtual void StopWith(std::function<void()> &_do);
  virtual void ThreadSleep(std::chrono::milliseconds _time);

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
