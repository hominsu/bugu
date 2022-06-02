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

#ifndef BUGU_DETECT_SERVICE_SRC_THREAD_POOL_X_TASK_H_
#define BUGU_DETECT_SERVICE_SRC_THREAD_POOL_X_TASK_H_

#include <functional>
#include <future>

namespace bugu {
/**
 * @brief 线程池任务基类
 */
class XTaskBase {
 public:
  ::std::function<bool()> is_running = nullptr; ///< 线程池运行状态函数指针

 public:
  virtual ~XTaskBase() = default;
  virtual void Main() = 0;
};

/**
 * @brief 线程池任务模版类，ret_type 设定值类型
 * @tparam ret_type 值类型，不允许为 ::std::thread
 */
template<class ret_type,
    class = typename ::std::enable_if
        <!::std::is_same<ret_type, ::std::thread>::value>::type
>
class XTask : public XTaskBase {
 private:
  ::std::promise<ret_type> p_; ///< 接收返回值

 public:
  /**
   * @brief 设置 future 的值
   * @param ret_type int value
   */
  void set_return(ret_type &&_value) {
    p_.set_value(::std::forward<ret_type>(_value));
  }

  void set_return(const ret_type &_value) {
    p_.set_value(_value);
  }

  /**
   * @brief 阻塞等待 set_value
   * @return decltype(auto)
   */
  decltype(auto) get_return() {
    return p_.get_future().get();
  }
};
} // namespace bugu

#endif //BUGU_DETECT_SERVICE_SRC_THREAD_POOL_X_TASK_H_
