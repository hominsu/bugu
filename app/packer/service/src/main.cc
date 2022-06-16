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

#include "bugu_packer_server.h"
#include "conf/config.h"
#include "thread_pool/x_thread_pool.h"
#include "utils/interrupt_sleeper.h"

#include <cstdio>
#include <csignal>

#include <memory_resource>
#include <memory>
#include <mutex>

::bugu::InterruptSleeper interrupt_sleeper;

void handler(int signal) {
  fprintf(stdout, "terminate with signal: %" PRId32 "\n", signal);
  interrupt_sleeper.interrupt();
}

int main() {
  auto bootstrap = ::std::make_shared<config::Bootstrap>();

  ::bugu::Config conf;

  conf.Load("/data/conf/config.json");
  conf.Scan(bootstrap.get());

  // Init the threadpool and memory-resource
  auto thread_pool = ::bugu::XThreadPool::Get();
  thread_pool->Init(::std::thread::hardware_concurrency());
  auto memory_resource = ::std::make_shared<::std::pmr::synchronized_pool_resource>();

  // Init Grpc server
  auto server = ::bugu::BuguPackerServer::Get();
  server->Init(bootstrap->server().grpc().addr(), thread_pool, memory_resource);
  server->Start();

  // capture the int and term signal
  struct sigaction sa{};
  sa.sa_handler = handler;
  sigemptyset(&sa.sa_mask);
  sa.sa_flags = 0;
  sigaction(SIGINT, &sa, nullptr);
  sigaction(SIGTERM, &sa, nullptr);

  // block to wait the exit
  interrupt_sleeper.wait();

  ::bugu::BuguPackerServer::Get()->Stop();
  ::bugu::XThreadPool::Get()->Stop();

  return 0;
}
