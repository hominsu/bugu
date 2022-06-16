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

#ifndef BUGU_PACKER_SERVICE_SRC_CONF_CONFIG_H_
#define BUGU_PACKER_SERVICE_SRC_CONF_CONFIG_H_

#include "bugu_packer/bugu_packer.h"
#include "conf.pb.h"

#include "google/protobuf/message.h"
#include "google/protobuf/util/json_util.h"

#include <cstdio>

#include <exception>

namespace bugu {

class Config {
 public:
  void Load(const char *_path);
  void Scan(::google::protobuf::Message *_message);

 private:
  void read(FILE *_in);

  ::std::string get_str() {
    return {buffer_.begin(), buffer_.end()};
  }

 private:
  std::vector<char> buffer_;
};

inline void Config::Load(const char *_path) {
  FILE *input;
#if BUGU_WINDOWS
  fopen_s(&input, _path, "r");
#elif BUGU_LINUX || BUGU_APPLE
  input = fopen(_path, "r");
#endif
  if (input == nullptr) { exit(EXIT_FAILURE); }
  read(input);
  fclose(input);
}

inline void Config::read(FILE *_in) {
  char buf[65536];
  if (!buffer_.empty()) { buffer_.clear(); }
  while (true) {
    size_t n = fread(buf, 1, sizeof(buf), _in);
    if (n == 0) { break; }
    buffer_.insert(buffer_.end(), buf, buf + n);
  }
}

void Config::Scan(::google::protobuf::Message *_message) {
  if (!google::protobuf::util::JsonStringToMessage(get_str(), _message).ok()) {
    throw ::std::runtime_error("Parse json data to message failed");
  }
}

} // namespace bugu

#endif //BUGU_PACKER_SERVICE_SRC_CONF_CONFIG_H_
