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

#include "credentials.h"

#include <fstream>

::std::string bugu::Credentials::GetFileContents(const std::string &_path) {
  ::std::ifstream file_stream(_path);
  if (!file_stream.good()) {
    ::std::cerr << "Open Cert File Failed" << ::std::endl;
    exit(-1);
  }
  ::std::string contents;
  contents.assign((::std::istreambuf_iterator<char>(file_stream)), ::std::istreambuf_iterator<char>());
  file_stream.close();
  return contents;
}

::std::shared_ptr<::grpc::ServerCredentials> bugu::Credentials::GetServerCredentials(const ::std::string &_root_cert_dir,
                                                                                     const ::std::string &_server_key_dir,
                                                                                     const ::std::string &_server_cert_dir) {
  auto root_cert = GetFileContents(_root_cert_dir);
  auto key_str = GetFileContents(_server_key_dir);
  auto cert_str = GetFileContents(_server_cert_dir);
  auto x509KeyPair = ::grpc::SslServerCredentialsOptions::PemKeyCertPair{key_str, cert_str};

  ::grpc::SslServerCredentialsOptions cred_option;
  cred_option.pem_root_certs = root_cert;
  cred_option.pem_key_cert_pairs.push_back(x509KeyPair);

  return ::grpc::SslServerCredentials(cred_option);
}

::std::shared_ptr<::grpc::ChannelCredentials> bugu::Credentials::GetClientCredentials(const ::std::string &_root_cert_dir,
                                                                                      const ::std::string &_client_key_dir,
                                                                                      const ::std::string &_client_cert_dir) {
  auto root_cert = GetFileContents(_root_cert_dir);
  auto key_str = GetFileContents(_client_key_dir);
  auto cert_str = GetFileContents(_client_cert_dir);

  ::grpc::SslCredentialsOptions cred_option;
  cred_option.pem_root_certs = root_cert;
  cred_option.pem_private_key = key_str;
  cred_option.pem_cert_chain = cert_str;

  return ::grpc::SslCredentials(cred_option);
}
