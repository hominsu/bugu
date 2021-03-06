// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: v1/bugu.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BuguClient is the client API for Bugu service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BuguClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
	GetCurrentUser(ctx context.Context, in *GetCurrentUserRequest, opts ...grpc.CallOption) (*GetCurrentUserReply, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserReply, error)
	GetFileMeta(ctx context.Context, in *GetFileMetaRequest, opts ...grpc.CallOption) (*GetFileMetaReply, error)
	GetFileMetaByUserId(ctx context.Context, in *GetFileMetaByUserIdRequest, opts ...grpc.CallOption) (*GetFileMetaByUserIdReply, error)
	DeleteFileMetadata(ctx context.Context, in *DeleteFileMetadataRequest, opts ...grpc.CallOption) (*DeleteFileMetadataReply, error)
	Detect(ctx context.Context, in *DetectRequest, opts ...grpc.CallOption) (*DetectReply, error)
	Confusion(ctx context.Context, in *ConfusionRequest, opts ...grpc.CallOption) (*ConfusionReply, error)
	Packer(ctx context.Context, in *PackerRequest, opts ...grpc.CallOption) (*PackerReply, error)
	GetArtifactMetadata(ctx context.Context, in *GetArtifactMetadataRequest, opts ...grpc.CallOption) (*GetArtifactMetadataReply, error)
	GetArtifactMetadataByFileId(ctx context.Context, in *GetArtifactMetadataByFileIdRequest, opts ...grpc.CallOption) (*GetArtifactMetadataByFileIdReply, error)
	DeleteArtifactMetadata(ctx context.Context, in *DeleteArtifactMetadataRequest, opts ...grpc.CallOption) (*DeleteArtifactMetadataReply, error)
}

type buguClient struct {
	cc grpc.ClientConnInterface
}

func NewBuguClient(cc grpc.ClientConnInterface) BuguClient {
	return &buguClient{cc}
}

func (c *buguClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error) {
	out := new(RegisterReply)
	err := c.cc.Invoke(ctx, "/bugu.service.v1.Bugu/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buguClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, "/bugu.service.v1.Bugu/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buguClient) GetCurrentUser(ctx context.Context, in *GetCurrentUserRequest, opts ...grpc.CallOption) (*GetCurrentUserReply, error) {
	out := new(GetCurrentUserReply)
	err := c.cc.Invoke(ctx, "/bugu.service.v1.Bugu/GetCurrentUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buguClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserReply, error) {
	out := new(UpdateUserReply)
	err := c.cc.Invoke(ctx, "/bugu.service.v1.Bugu/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buguClient) GetFileMeta(ctx context.Context, in *GetFileMetaRequest, opts ...grpc.CallOption) (*GetFileMetaReply, error) {
	out := new(GetFileMetaReply)
	err := c.cc.Invoke(ctx, "/bugu.service.v1.Bugu/GetFileMeta", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buguClient) GetFileMetaByUserId(ctx context.Context, in *GetFileMetaByUserIdRequest, opts ...grpc.CallOption) (*GetFileMetaByUserIdReply, error) {
	out := new(GetFileMetaByUserIdReply)
	err := c.cc.Invoke(ctx, "/bugu.service.v1.Bugu/GetFileMetaByUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buguClient) DeleteFileMetadata(ctx context.Context, in *DeleteFileMetadataRequest, opts ...grpc.CallOption) (*DeleteFileMetadataReply, error) {
	out := new(DeleteFileMetadataReply)
	err := c.cc.Invoke(ctx, "/bugu.service.v1.Bugu/DeleteFileMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buguClient) Detect(ctx context.Context, in *DetectRequest, opts ...grpc.CallOption) (*DetectReply, error) {
	out := new(DetectReply)
	err := c.cc.Invoke(ctx, "/bugu.service.v1.Bugu/Detect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buguClient) Confusion(ctx context.Context, in *ConfusionRequest, opts ...grpc.CallOption) (*ConfusionReply, error) {
	out := new(ConfusionReply)
	err := c.cc.Invoke(ctx, "/bugu.service.v1.Bugu/Confusion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buguClient) Packer(ctx context.Context, in *PackerRequest, opts ...grpc.CallOption) (*PackerReply, error) {
	out := new(PackerReply)
	err := c.cc.Invoke(ctx, "/bugu.service.v1.Bugu/Packer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buguClient) GetArtifactMetadata(ctx context.Context, in *GetArtifactMetadataRequest, opts ...grpc.CallOption) (*GetArtifactMetadataReply, error) {
	out := new(GetArtifactMetadataReply)
	err := c.cc.Invoke(ctx, "/bugu.service.v1.Bugu/GetArtifactMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buguClient) GetArtifactMetadataByFileId(ctx context.Context, in *GetArtifactMetadataByFileIdRequest, opts ...grpc.CallOption) (*GetArtifactMetadataByFileIdReply, error) {
	out := new(GetArtifactMetadataByFileIdReply)
	err := c.cc.Invoke(ctx, "/bugu.service.v1.Bugu/GetArtifactMetadataByFileId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buguClient) DeleteArtifactMetadata(ctx context.Context, in *DeleteArtifactMetadataRequest, opts ...grpc.CallOption) (*DeleteArtifactMetadataReply, error) {
	out := new(DeleteArtifactMetadataReply)
	err := c.cc.Invoke(ctx, "/bugu.service.v1.Bugu/DeleteArtifactMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BuguServer is the server API for Bugu service.
// All implementations must embed UnimplementedBuguServer
// for forward compatibility
type BuguServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	GetCurrentUser(context.Context, *GetCurrentUserRequest) (*GetCurrentUserReply, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserReply, error)
	GetFileMeta(context.Context, *GetFileMetaRequest) (*GetFileMetaReply, error)
	GetFileMetaByUserId(context.Context, *GetFileMetaByUserIdRequest) (*GetFileMetaByUserIdReply, error)
	DeleteFileMetadata(context.Context, *DeleteFileMetadataRequest) (*DeleteFileMetadataReply, error)
	Detect(context.Context, *DetectRequest) (*DetectReply, error)
	Confusion(context.Context, *ConfusionRequest) (*ConfusionReply, error)
	Packer(context.Context, *PackerRequest) (*PackerReply, error)
	GetArtifactMetadata(context.Context, *GetArtifactMetadataRequest) (*GetArtifactMetadataReply, error)
	GetArtifactMetadataByFileId(context.Context, *GetArtifactMetadataByFileIdRequest) (*GetArtifactMetadataByFileIdReply, error)
	DeleteArtifactMetadata(context.Context, *DeleteArtifactMetadataRequest) (*DeleteArtifactMetadataReply, error)
	mustEmbedUnimplementedBuguServer()
}

// UnimplementedBuguServer must be embedded to have forward compatible implementations.
type UnimplementedBuguServer struct {
}

func (UnimplementedBuguServer) Register(context.Context, *RegisterRequest) (*RegisterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedBuguServer) Login(context.Context, *LoginRequest) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedBuguServer) GetCurrentUser(context.Context, *GetCurrentUserRequest) (*GetCurrentUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentUser not implemented")
}
func (UnimplementedBuguServer) UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedBuguServer) GetFileMeta(context.Context, *GetFileMetaRequest) (*GetFileMetaReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFileMeta not implemented")
}
func (UnimplementedBuguServer) GetFileMetaByUserId(context.Context, *GetFileMetaByUserIdRequest) (*GetFileMetaByUserIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFileMetaByUserId not implemented")
}
func (UnimplementedBuguServer) DeleteFileMetadata(context.Context, *DeleteFileMetadataRequest) (*DeleteFileMetadataReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFileMetadata not implemented")
}
func (UnimplementedBuguServer) Detect(context.Context, *DetectRequest) (*DetectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Detect not implemented")
}
func (UnimplementedBuguServer) Confusion(context.Context, *ConfusionRequest) (*ConfusionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Confusion not implemented")
}
func (UnimplementedBuguServer) Packer(context.Context, *PackerRequest) (*PackerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Packer not implemented")
}
func (UnimplementedBuguServer) GetArtifactMetadata(context.Context, *GetArtifactMetadataRequest) (*GetArtifactMetadataReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArtifactMetadata not implemented")
}
func (UnimplementedBuguServer) GetArtifactMetadataByFileId(context.Context, *GetArtifactMetadataByFileIdRequest) (*GetArtifactMetadataByFileIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArtifactMetadataByFileId not implemented")
}
func (UnimplementedBuguServer) DeleteArtifactMetadata(context.Context, *DeleteArtifactMetadataRequest) (*DeleteArtifactMetadataReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArtifactMetadata not implemented")
}
func (UnimplementedBuguServer) mustEmbedUnimplementedBuguServer() {}

// UnsafeBuguServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BuguServer will
// result in compilation errors.
type UnsafeBuguServer interface {
	mustEmbedUnimplementedBuguServer()
}

func RegisterBuguServer(s grpc.ServiceRegistrar, srv BuguServer) {
	s.RegisterService(&Bugu_ServiceDesc, srv)
}

func _Bugu_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuguServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bugu.service.v1.Bugu/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuguServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bugu_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuguServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bugu.service.v1.Bugu/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuguServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bugu_GetCurrentUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrentUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuguServer).GetCurrentUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bugu.service.v1.Bugu/GetCurrentUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuguServer).GetCurrentUser(ctx, req.(*GetCurrentUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bugu_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuguServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bugu.service.v1.Bugu/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuguServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bugu_GetFileMeta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFileMetaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuguServer).GetFileMeta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bugu.service.v1.Bugu/GetFileMeta",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuguServer).GetFileMeta(ctx, req.(*GetFileMetaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bugu_GetFileMetaByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFileMetaByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuguServer).GetFileMetaByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bugu.service.v1.Bugu/GetFileMetaByUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuguServer).GetFileMetaByUserId(ctx, req.(*GetFileMetaByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bugu_DeleteFileMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFileMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuguServer).DeleteFileMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bugu.service.v1.Bugu/DeleteFileMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuguServer).DeleteFileMetadata(ctx, req.(*DeleteFileMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bugu_Detect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuguServer).Detect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bugu.service.v1.Bugu/Detect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuguServer).Detect(ctx, req.(*DetectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bugu_Confusion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfusionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuguServer).Confusion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bugu.service.v1.Bugu/Confusion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuguServer).Confusion(ctx, req.(*ConfusionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bugu_Packer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PackerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuguServer).Packer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bugu.service.v1.Bugu/Packer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuguServer).Packer(ctx, req.(*PackerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bugu_GetArtifactMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArtifactMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuguServer).GetArtifactMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bugu.service.v1.Bugu/GetArtifactMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuguServer).GetArtifactMetadata(ctx, req.(*GetArtifactMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bugu_GetArtifactMetadataByFileId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArtifactMetadataByFileIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuguServer).GetArtifactMetadataByFileId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bugu.service.v1.Bugu/GetArtifactMetadataByFileId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuguServer).GetArtifactMetadataByFileId(ctx, req.(*GetArtifactMetadataByFileIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bugu_DeleteArtifactMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteArtifactMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuguServer).DeleteArtifactMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bugu.service.v1.Bugu/DeleteArtifactMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuguServer).DeleteArtifactMetadata(ctx, req.(*DeleteArtifactMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Bugu_ServiceDesc is the grpc.ServiceDesc for Bugu service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Bugu_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bugu.service.v1.Bugu",
	HandlerType: (*BuguServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Bugu_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Bugu_Login_Handler,
		},
		{
			MethodName: "GetCurrentUser",
			Handler:    _Bugu_GetCurrentUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _Bugu_UpdateUser_Handler,
		},
		{
			MethodName: "GetFileMeta",
			Handler:    _Bugu_GetFileMeta_Handler,
		},
		{
			MethodName: "GetFileMetaByUserId",
			Handler:    _Bugu_GetFileMetaByUserId_Handler,
		},
		{
			MethodName: "DeleteFileMetadata",
			Handler:    _Bugu_DeleteFileMetadata_Handler,
		},
		{
			MethodName: "Detect",
			Handler:    _Bugu_Detect_Handler,
		},
		{
			MethodName: "Confusion",
			Handler:    _Bugu_Confusion_Handler,
		},
		{
			MethodName: "Packer",
			Handler:    _Bugu_Packer_Handler,
		},
		{
			MethodName: "GetArtifactMetadata",
			Handler:    _Bugu_GetArtifactMetadata_Handler,
		},
		{
			MethodName: "GetArtifactMetadataByFileId",
			Handler:    _Bugu_GetArtifactMetadataByFileId_Handler,
		},
		{
			MethodName: "DeleteArtifactMetadata",
			Handler:    _Bugu_DeleteArtifactMetadata_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/bugu.proto",
}
