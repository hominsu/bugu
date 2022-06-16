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

#ifndef BUGU_PACKER_SERVICE_SRC_UTILS_CREDENTIALS_H_
#define BUGU_PACKER_SERVICE_SRC_UTILS_CREDENTIALS_H_

#include <grpc++/grpc++.h>

#include <string>
#include <memory>

namespace bugu {

class Credentials {
 public:
  static ::std::string GetFileContents(const ::std::string &_path);

  static ::std::shared_ptr<::grpc::ServerCredentials> GetServerCredentials(const ::std::string &_root_cert_dir = "/cert/ca.crt",
                                                                           const ::std::string &_server_key_dir = "/cert/server.key",
                                                                           const ::std::string &_server_cert_dir = "/cert/server.pem");
  static ::std::shared_ptr<::grpc::ChannelCredentials> GetClientCredentials(const ::std::string &_root_cert_dir = "/cert/ca.crt",
                                                                            const ::std::string &_client_key_dir = "/cert/client.key",
                                                                            const ::std::string &_client_cert_dir = "/cert/client.pem");
};

} // namespace bugu

#endif //BUGU_PACKER_SERVICE_SRC_UTILS_CREDENTIALS_H_
