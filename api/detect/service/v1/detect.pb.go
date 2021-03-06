// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: v1/detect.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Type int32

const (
	Type_TYPE_ADPOSHEL     Type = 0
	Type_TYPE_AGENT        Type = 1
	Type_TYPE_ALLAPLE      Type = 2
	Type_TYPE_AMONETIZE    Type = 3
	Type_TYPE_ANDROM       Type = 4
	Type_TYPE_AUTORUN      Type = 5
	Type_TYPE_BROWSE_FOX   Type = 6
	Type_TYPE_DINWOD       Type = 7
	Type_TYPE_ELEX         Type = 8
	Type_TYPE_EXPIRO       Type = 9
	Type_TYPE_FASONG       Type = 10
	Type_TYPE_HACK_KMS     Type = 11
	Type_TYPE_HLUX         Type = 12
	Type_TYPE_INJECTOR     Type = 13
	Type_TYPE_INSTALL_CORE Type = 14
	Type_TYPE_MULTI_Plug   Type = 15
	Type_TYPE_NEOREKLAMI   Type = 16
	Type_TYPE_NESHTA       Type = 17
	Type_TYPE_OTHER        Type = 18
	Type_TYPE_REGRUN       Type = 19
	Type_TYPE_SALITY       Type = 20
	Type_TYPE_SNARASITE    Type = 21
	Type_TYPE_STABTINKO    Type = 22
	Type_TYPE_VBA          Type = 23
	Type_TYPE_VBKRYPT      Type = 24
	Type_TYPE_VILSEL       Type = 25
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0:  "TYPE_ADPOSHEL",
		1:  "TYPE_AGENT",
		2:  "TYPE_ALLAPLE",
		3:  "TYPE_AMONETIZE",
		4:  "TYPE_ANDROM",
		5:  "TYPE_AUTORUN",
		6:  "TYPE_BROWSE_FOX",
		7:  "TYPE_DINWOD",
		8:  "TYPE_ELEX",
		9:  "TYPE_EXPIRO",
		10: "TYPE_FASONG",
		11: "TYPE_HACK_KMS",
		12: "TYPE_HLUX",
		13: "TYPE_INJECTOR",
		14: "TYPE_INSTALL_CORE",
		15: "TYPE_MULTI_Plug",
		16: "TYPE_NEOREKLAMI",
		17: "TYPE_NESHTA",
		18: "TYPE_OTHER",
		19: "TYPE_REGRUN",
		20: "TYPE_SALITY",
		21: "TYPE_SNARASITE",
		22: "TYPE_STABTINKO",
		23: "TYPE_VBA",
		24: "TYPE_VBKRYPT",
		25: "TYPE_VILSEL",
	}
	Type_value = map[string]int32{
		"TYPE_ADPOSHEL":     0,
		"TYPE_AGENT":        1,
		"TYPE_ALLAPLE":      2,
		"TYPE_AMONETIZE":    3,
		"TYPE_ANDROM":       4,
		"TYPE_AUTORUN":      5,
		"TYPE_BROWSE_FOX":   6,
		"TYPE_DINWOD":       7,
		"TYPE_ELEX":         8,
		"TYPE_EXPIRO":       9,
		"TYPE_FASONG":       10,
		"TYPE_HACK_KMS":     11,
		"TYPE_HLUX":         12,
		"TYPE_INJECTOR":     13,
		"TYPE_INSTALL_CORE": 14,
		"TYPE_MULTI_Plug":   15,
		"TYPE_NEOREKLAMI":   16,
		"TYPE_NESHTA":       17,
		"TYPE_OTHER":        18,
		"TYPE_REGRUN":       19,
		"TYPE_SALITY":       20,
		"TYPE_SNARASITE":    21,
		"TYPE_STABTINKO":    22,
		"TYPE_VBA":          23,
		"TYPE_VBKRYPT":      24,
		"TYPE_VILSEL":       25,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_detect_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_v1_detect_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_v1_detect_proto_rawDescGZIP(), []int{0}
}

type DetectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data [][]byte `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Size uint32   `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *DetectRequest) Reset() {
	*x = DetectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_detect_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetectRequest) ProtoMessage() {}

func (x *DetectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_detect_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetectRequest.ProtoReflect.Descriptor instead.
func (*DetectRequest) Descriptor() ([]byte, []int) {
	return file_v1_detect_proto_rawDescGZIP(), []int{0}
}

func (x *DetectRequest) GetData() [][]byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *DetectRequest) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type DetectReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type Type `protobuf:"varint,1,opt,name=type,proto3,enum=bugu_detect.service.v1.Type" json:"type,omitempty"`
}

func (x *DetectReply) Reset() {
	*x = DetectReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_detect_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetectReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetectReply) ProtoMessage() {}

func (x *DetectReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_detect_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetectReply.ProtoReflect.Descriptor instead.
func (*DetectReply) Descriptor() ([]byte, []int) {
	return file_v1_detect_proto_rawDescGZIP(), []int{1}
}

func (x *DetectReply) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_TYPE_ADPOSHEL
}

var File_v1_detect_proto protoreflect.FileDescriptor

var file_v1_detect_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x16, 0x62, 0x75, 0x67, 0x75, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x37, 0x0a, 0x0d, 0x44, 0x65, 0x74,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x22, 0x3f, 0x0a, 0x0b, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x30, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1c, 0x2e, 0x62, 0x75, 0x67, 0x75, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x2a, 0xdb, 0x03, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x11, 0x0a, 0x0d,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x44, 0x50, 0x4f, 0x53, 0x48, 0x45, 0x4c, 0x10, 0x00, 0x12,
	0x0e, 0x0a, 0x0a, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12,
	0x10, 0x0a, 0x0c, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x4c, 0x4c, 0x41, 0x50, 0x4c, 0x45, 0x10,
	0x02, 0x12, 0x12, 0x0a, 0x0e, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x4d, 0x4f, 0x4e, 0x45, 0x54,
	0x49, 0x5a, 0x45, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x4e,
	0x44, 0x52, 0x4f, 0x4d, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41,
	0x55, 0x54, 0x4f, 0x52, 0x55, 0x4e, 0x10, 0x05, 0x12, 0x13, 0x0a, 0x0f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x42, 0x52, 0x4f, 0x57, 0x53, 0x45, 0x5f, 0x46, 0x4f, 0x58, 0x10, 0x06, 0x12, 0x0f, 0x0a,
	0x0b, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x49, 0x4e, 0x57, 0x4f, 0x44, 0x10, 0x07, 0x12, 0x0d,
	0x0a, 0x09, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x4c, 0x45, 0x58, 0x10, 0x08, 0x12, 0x0f, 0x0a,
	0x0b, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x58, 0x50, 0x49, 0x52, 0x4f, 0x10, 0x09, 0x12, 0x0f,
	0x0a, 0x0b, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x41, 0x53, 0x4f, 0x4e, 0x47, 0x10, 0x0a, 0x12,
	0x11, 0x0a, 0x0d, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x48, 0x41, 0x43, 0x4b, 0x5f, 0x4b, 0x4d, 0x53,
	0x10, 0x0b, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x48, 0x4c, 0x55, 0x58, 0x10,
	0x0c, 0x12, 0x11, 0x0a, 0x0d, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x4a, 0x45, 0x43, 0x54,
	0x4f, 0x52, 0x10, 0x0d, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x53,
	0x54, 0x41, 0x4c, 0x4c, 0x5f, 0x43, 0x4f, 0x52, 0x45, 0x10, 0x0e, 0x12, 0x13, 0x0a, 0x0f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x4d, 0x55, 0x4c, 0x54, 0x49, 0x5f, 0x50, 0x6c, 0x75, 0x67, 0x10, 0x0f,
	0x12, 0x13, 0x0a, 0x0f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x45, 0x4f, 0x52, 0x45, 0x4b, 0x4c,
	0x41, 0x4d, 0x49, 0x10, 0x10, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x45,
	0x53, 0x48, 0x54, 0x41, 0x10, 0x11, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f,
	0x54, 0x48, 0x45, 0x52, 0x10, 0x12, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52,
	0x45, 0x47, 0x52, 0x55, 0x4e, 0x10, 0x13, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x53, 0x41, 0x4c, 0x49, 0x54, 0x59, 0x10, 0x14, 0x12, 0x12, 0x0a, 0x0e, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x53, 0x4e, 0x41, 0x52, 0x41, 0x53, 0x49, 0x54, 0x45, 0x10, 0x15, 0x12, 0x12, 0x0a, 0x0e,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x42, 0x54, 0x49, 0x4e, 0x4b, 0x4f, 0x10, 0x16,
	0x12, 0x0c, 0x0a, 0x08, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x56, 0x42, 0x41, 0x10, 0x17, 0x12, 0x10,
	0x0a, 0x0c, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x56, 0x42, 0x4b, 0x52, 0x59, 0x50, 0x54, 0x10, 0x18,
	0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x56, 0x49, 0x4c, 0x53, 0x45, 0x4c, 0x10,
	0x19, 0x32, 0x64, 0x0a, 0x0a, 0x42, 0x75, 0x67, 0x75, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x12,
	0x56, 0x0a, 0x06, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x12, 0x25, 0x2e, 0x62, 0x75, 0x67, 0x75,
	0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x62, 0x75, 0x67, 0x75, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x6f, 0x6d, 0x69, 0x6e, 0x73, 0x75, 0x2f, 0x62, 0x75,
	0x67, 0x75, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_v1_detect_proto_rawDescOnce sync.Once
	file_v1_detect_proto_rawDescData = file_v1_detect_proto_rawDesc
)

func file_v1_detect_proto_rawDescGZIP() []byte {
	file_v1_detect_proto_rawDescOnce.Do(func() {
		file_v1_detect_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_detect_proto_rawDescData)
	})
	return file_v1_detect_proto_rawDescData
}

var file_v1_detect_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_detect_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_detect_proto_goTypes = []interface{}{
	(Type)(0),             // 0: bugu_detect.service.v1.Type
	(*DetectRequest)(nil), // 1: bugu_detect.service.v1.DetectRequest
	(*DetectReply)(nil),   // 2: bugu_detect.service.v1.DetectReply
}
var file_v1_detect_proto_depIdxs = []int32{
	0, // 0: bugu_detect.service.v1.DetectReply.type:type_name -> bugu_detect.service.v1.Type
	1, // 1: bugu_detect.service.v1.BuguDetect.Detect:input_type -> bugu_detect.service.v1.DetectRequest
	2, // 2: bugu_detect.service.v1.BuguDetect.Detect:output_type -> bugu_detect.service.v1.DetectReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v1_detect_proto_init() }
func file_v1_detect_proto_init() {
	if File_v1_detect_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_detect_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetectRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_detect_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetectReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_detect_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_detect_proto_goTypes,
		DependencyIndexes: file_v1_detect_proto_depIdxs,
		EnumInfos:         file_v1_detect_proto_enumTypes,
		MessageInfos:      file_v1_detect_proto_msgTypes,
	}.Build()
	File_v1_detect_proto = out.File
	file_v1_detect_proto_rawDesc = nil
	file_v1_detect_proto_goTypes = nil
	file_v1_detect_proto_depIdxs = nil
}
