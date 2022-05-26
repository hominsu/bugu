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
// Created by HominSu on 2022/5/16.
//

#include "bugu_obfusion_server.h"
#include "thread_pool/x_thread_pool.h"

#include <cstdio>
#include <csignal>

void handler(int signal) {
  fprintf(stdout, "receive the signal: %d", signal);
  ::bugu::BuguObfusionServer::Get()->Stop();
  ::bugu::XThreadPool::Get()->Stop();
}

int main() {
  auto thread_pool = ::bugu::XThreadPool::Get();
  thread_pool->Init(::std::thread::hardware_concurrency());

  auto memory_resource = ::std::make_shared<std::pmr::synchronized_pool_resource>();

  auto server = ::bugu::BuguObfusionServer::Get();
  server->Init("127.0.0.1", thread_pool, memory_resource);
  server->Start();

  signal(SIGINT, handler);
  signal(SIGQUIT, handler);

  // TODO: block to wait the exit

  return 0;
}
