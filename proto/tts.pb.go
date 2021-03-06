// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: tts.proto

package api

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Text struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input string `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
}

func (x *Text) Reset() {
	*x = Text{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Text) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Text) ProtoMessage() {}

func (x *Text) ProtoReflect() protoreflect.Message {
	mi := &file_tts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Text.ProtoReflect.Descriptor instead.
func (*Text) Descriptor() ([]byte, []int) {
	return file_tts_proto_rawDescGZIP(), []int{0}
}

func (x *Text) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

type Speech struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Output []byte `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *Speech) Reset() {
	*x = Speech{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Speech) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Speech) ProtoMessage() {}

func (x *Speech) ProtoReflect() protoreflect.Message {
	mi := &file_tts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Speech.ProtoReflect.Descriptor instead.
func (*Speech) Descriptor() ([]byte, []int) {
	return file_tts_proto_rawDescGZIP(), []int{1}
}

func (x *Speech) GetOutput() []byte {
	if x != nil {
		return x.Output
	}
	return nil
}

var File_tts_proto protoreflect.FileDescriptor

var file_tts_proto_rawDesc = []byte{
	0x0a, 0x09, 0x74, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1c, 0x0a, 0x04, 0x54,
	0x65, 0x78, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x20, 0x0a, 0x06, 0x53, 0x70, 0x65,
	0x65, 0x63, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x32, 0x20, 0x0a, 0x03, 0x54,
	0x54, 0x53, 0x12, 0x19, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x12, 0x05, 0x2e,
	0x54, 0x65, 0x78, 0x74, 0x1a, 0x07, 0x2e, 0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x42, 0x06, 0x5a,
	0x04, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tts_proto_rawDescOnce sync.Once
	file_tts_proto_rawDescData = file_tts_proto_rawDesc
)

func file_tts_proto_rawDescGZIP() []byte {
	file_tts_proto_rawDescOnce.Do(func() {
		file_tts_proto_rawDescData = protoimpl.X.CompressGZIP(file_tts_proto_rawDescData)
	})
	return file_tts_proto_rawDescData
}

var file_tts_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tts_proto_goTypes = []interface{}{
	(*Text)(nil),   // 0: Text
	(*Speech)(nil), // 1: Speech
}
var file_tts_proto_depIdxs = []int32{
	0, // 0: TTS.Process:input_type -> Text
	1, // 1: TTS.Process:output_type -> Speech
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tts_proto_init() }
func file_tts_proto_init() {
	if File_tts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Text); i {
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
		file_tts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Speech); i {
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
			RawDescriptor: file_tts_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tts_proto_goTypes,
		DependencyIndexes: file_tts_proto_depIdxs,
		MessageInfos:      file_tts_proto_msgTypes,
	}.Build()
	File_tts_proto = out.File
	file_tts_proto_rawDesc = nil
	file_tts_proto_goTypes = nil
	file_tts_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TTSClient is the client API for TTS service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TTSClient interface {
	Process(ctx context.Context, in *Text, opts ...grpc.CallOption) (*Speech, error)
}

type tTSClient struct {
	cc grpc.ClientConnInterface
}

func NewTTSClient(cc grpc.ClientConnInterface) TTSClient {
	return &tTSClient{cc}
}

func (c *tTSClient) Process(ctx context.Context, in *Text, opts ...grpc.CallOption) (*Speech, error) {
	out := new(Speech)
	err := c.cc.Invoke(ctx, "/TTS/Process", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TTSServer is the server API for TTS service.
type TTSServer interface {
	Process(context.Context, *Text) (*Speech, error)
}

// UnimplementedTTSServer can be embedded to have forward compatible implementations.
type UnimplementedTTSServer struct {
}

func (*UnimplementedTTSServer) Process(context.Context, *Text) (*Speech, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Process not implemented")
}

func RegisterTTSServer(s *grpc.Server, srv TTSServer) {
	s.RegisterService(&_TTS_serviceDesc, srv)
}

func _TTS_Process_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Text)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TTSServer).Process(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TTS/Process",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TTSServer).Process(ctx, req.(*Text))
	}
	return interceptor(ctx, in, info, handler)
}

var _TTS_serviceDesc = grpc.ServiceDesc{
	ServiceName: "TTS",
	HandlerType: (*TTSServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Process",
			Handler:    _TTS_Process_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tts.proto",
}
