// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: detect.proto

#include "detect.pb.h"
#include "detect.grpc.pb.h"

#include <functional>
#include <grpcpp/impl/codegen/async_stream.h>
#include <grpcpp/impl/codegen/async_unary_call.h>
#include <grpcpp/impl/codegen/channel_interface.h>
#include <grpcpp/impl/codegen/client_unary_call.h>
#include <grpcpp/impl/codegen/client_callback.h>
#include <grpcpp/impl/codegen/message_allocator.h>
#include <grpcpp/impl/codegen/method_handler.h>
#include <grpcpp/impl/codegen/rpc_service_method.h>
#include <grpcpp/impl/codegen/server_callback.h>
#include <grpcpp/impl/codegen/server_callback_handlers.h>
#include <grpcpp/impl/codegen/server_context.h>
#include <grpcpp/impl/codegen/service_type.h>
#include <grpcpp/impl/codegen/sync_stream.h>
namespace bugu_detect {
namespace service {
namespace v1 {

static const char* BuguDetect_method_names[] = {
  "/bugu_detect.service.v1.BuguDetect/Detect",
};

std::unique_ptr< BuguDetect::Stub> BuguDetect::NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options) {
  (void)options;
  std::unique_ptr< BuguDetect::Stub> stub(new BuguDetect::Stub(channel, options));
  return stub;
}

BuguDetect::Stub::Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options)
  : channel_(channel), rpcmethod_Detect_(BuguDetect_method_names[0], options.suffix_for_stats(),::grpc::internal::RpcMethod::NORMAL_RPC, channel)
  {}

::grpc::Status BuguDetect::Stub::Detect(::grpc::ClientContext* context, const ::bugu_detect::service::v1::DetectRequest& request, ::bugu_detect::service::v1::DetectReply* response) {
  return ::grpc::internal::BlockingUnaryCall< ::bugu_detect::service::v1::DetectRequest, ::bugu_detect::service::v1::DetectReply, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), rpcmethod_Detect_, context, request, response);
}

void BuguDetect::Stub::async::Detect(::grpc::ClientContext* context, const ::bugu_detect::service::v1::DetectRequest* request, ::bugu_detect::service::v1::DetectReply* response, std::function<void(::grpc::Status)> f) {
  ::grpc::internal::CallbackUnaryCall< ::bugu_detect::service::v1::DetectRequest, ::bugu_detect::service::v1::DetectReply, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_Detect_, context, request, response, std::move(f));
}

void BuguDetect::Stub::async::Detect(::grpc::ClientContext* context, const ::bugu_detect::service::v1::DetectRequest* request, ::bugu_detect::service::v1::DetectReply* response, ::grpc::ClientUnaryReactor* reactor) {
  ::grpc::internal::ClientCallbackUnaryFactory::Create< ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(stub_->channel_.get(), stub_->rpcmethod_Detect_, context, request, response, reactor);
}

::grpc::ClientAsyncResponseReader< ::bugu_detect::service::v1::DetectReply>* BuguDetect::Stub::PrepareAsyncDetectRaw(::grpc::ClientContext* context, const ::bugu_detect::service::v1::DetectRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::internal::ClientAsyncResponseReaderHelper::Create< ::bugu_detect::service::v1::DetectReply, ::bugu_detect::service::v1::DetectRequest, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(channel_.get(), cq, rpcmethod_Detect_, context, request);
}

::grpc::ClientAsyncResponseReader< ::bugu_detect::service::v1::DetectReply>* BuguDetect::Stub::AsyncDetectRaw(::grpc::ClientContext* context, const ::bugu_detect::service::v1::DetectRequest& request, ::grpc::CompletionQueue* cq) {
  auto* result =
    this->PrepareAsyncDetectRaw(context, request, cq);
  result->StartCall();
  return result;
}

BuguDetect::Service::Service() {
  AddMethod(new ::grpc::internal::RpcServiceMethod(
      BuguDetect_method_names[0],
      ::grpc::internal::RpcMethod::NORMAL_RPC,
      new ::grpc::internal::RpcMethodHandler< BuguDetect::Service, ::bugu_detect::service::v1::DetectRequest, ::bugu_detect::service::v1::DetectReply, ::grpc::protobuf::MessageLite, ::grpc::protobuf::MessageLite>(
          [](BuguDetect::Service* service,
             ::grpc::ServerContext* ctx,
             const ::bugu_detect::service::v1::DetectRequest* req,
             ::bugu_detect::service::v1::DetectReply* resp) {
               return service->Detect(ctx, req, resp);
             }, this)));
}

BuguDetect::Service::~Service() {
}

::grpc::Status BuguDetect::Service::Detect(::grpc::ServerContext* context, const ::bugu_detect::service::v1::DetectRequest* request, ::bugu_detect::service::v1::DetectReply* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}


}  // namespace bugu_detect
}  // namespace service
}  // namespace v1

