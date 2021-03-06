// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: conveyor.proto

package conveyor

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

type OK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *OK) Reset() {
	*x = OK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conveyor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OK) ProtoMessage() {}

func (x *OK) ProtoReflect() protoreflect.Message {
	mi := &file_conveyor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OK.ProtoReflect.Descriptor instead.
func (*OK) Descriptor() ([]byte, []int) {
	return file_conveyor_proto_rawDescGZIP(), []int{0}
}

func (x *OK) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type Position struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float32 `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float32 `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
	Z float32 `protobuf:"fixed32,3,opt,name=z,proto3" json:"z,omitempty"`
}

func (x *Position) Reset() {
	*x = Position{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conveyor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Position) ProtoMessage() {}

func (x *Position) ProtoReflect() protoreflect.Message {
	mi := &file_conveyor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Position.ProtoReflect.Descriptor instead.
func (*Position) Descriptor() ([]byte, []int) {
	return file_conveyor_proto_rawDescGZIP(), []int{1}
}

func (x *Position) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Position) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Position) GetZ() float32 {
	if x != nil {
		return x.Z
	}
	return 0
}

type Engines struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Green  int32 `protobuf:"varint,1,opt,name=green,proto3" json:"green,omitempty"`
	Red    int32 `protobuf:"varint,2,opt,name=red,proto3" json:"red,omitempty"`
	Black  int32 `protobuf:"varint,3,opt,name=black,proto3" json:"black,omitempty"`
	Orange int32 `protobuf:"varint,4,opt,name=orange,proto3" json:"orange,omitempty"`
}

func (x *Engines) Reset() {
	*x = Engines{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conveyor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Engines) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Engines) ProtoMessage() {}

func (x *Engines) ProtoReflect() protoreflect.Message {
	mi := &file_conveyor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Engines.ProtoReflect.Descriptor instead.
func (*Engines) Descriptor() ([]byte, []int) {
	return file_conveyor_proto_rawDescGZIP(), []int{2}
}

func (x *Engines) GetGreen() int32 {
	if x != nil {
		return x.Green
	}
	return 0
}

func (x *Engines) GetRed() int32 {
	if x != nil {
		return x.Red
	}
	return 0
}

func (x *Engines) GetBlack() int32 {
	if x != nil {
		return x.Black
	}
	return 0
}

func (x *Engines) GetOrange() int32 {
	if x != nil {
		return x.Orange
	}
	return 0
}

var File_conveyor_proto protoreflect.FileDescriptor

var file_conveyor_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x79, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x79, 0x6f, 0x72, 0x22, 0x1e, 0x0a, 0x02, 0x4f, 0x4b,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x34, 0x0a, 0x08, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x7a,
	0x22, 0x5f, 0x0a, 0x07, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x67,
	0x72, 0x65, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x67, 0x72, 0x65, 0x65,
	0x6e, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x72, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x61,
	0x6e, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x72, 0x61, 0x6e, 0x67,
	0x65, 0x32, 0x7a, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x79, 0x6f, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x0b, 0x4e, 0x65, 0x77, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x12, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x79, 0x6f, 0x72, 0x2e, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0c, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x79,
	0x6f, 0x72, 0x2e, 0x4f, 0x4b, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0f, 0x4e, 0x65, 0x77, 0x45, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x63, 0x6f, 0x6e,
	0x76, 0x65, 0x79, 0x6f, 0x72, 0x2e, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x73, 0x1a, 0x0c, 0x2e,
	0x63, 0x6f, 0x6e, 0x76, 0x65, 0x79, 0x6f, 0x72, 0x2e, 0x4f, 0x4b, 0x22, 0x00, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_conveyor_proto_rawDescOnce sync.Once
	file_conveyor_proto_rawDescData = file_conveyor_proto_rawDesc
)

func file_conveyor_proto_rawDescGZIP() []byte {
	file_conveyor_proto_rawDescOnce.Do(func() {
		file_conveyor_proto_rawDescData = protoimpl.X.CompressGZIP(file_conveyor_proto_rawDescData)
	})
	return file_conveyor_proto_rawDescData
}

var file_conveyor_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_conveyor_proto_goTypes = []interface{}{
	(*OK)(nil),       // 0: conveyor.OK
	(*Position)(nil), // 1: conveyor.Position
	(*Engines)(nil),  // 2: conveyor.Engines
}
var file_conveyor_proto_depIdxs = []int32{
	1, // 0: conveyor.ConveyorService.NewPosition:input_type -> conveyor.Position
	2, // 1: conveyor.ConveyorService.NewEnginesState:input_type -> conveyor.Engines
	0, // 2: conveyor.ConveyorService.NewPosition:output_type -> conveyor.OK
	0, // 3: conveyor.ConveyorService.NewEnginesState:output_type -> conveyor.OK
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_conveyor_proto_init() }
func file_conveyor_proto_init() {
	if File_conveyor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_conveyor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OK); i {
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
		file_conveyor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Position); i {
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
		file_conveyor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Engines); i {
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
			RawDescriptor: file_conveyor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_conveyor_proto_goTypes,
		DependencyIndexes: file_conveyor_proto_depIdxs,
		MessageInfos:      file_conveyor_proto_msgTypes,
	}.Build()
	File_conveyor_proto = out.File
	file_conveyor_proto_rawDesc = nil
	file_conveyor_proto_goTypes = nil
	file_conveyor_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ConveyorServiceClient is the client API for ConveyorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConveyorServiceClient interface {
	NewPosition(ctx context.Context, in *Position, opts ...grpc.CallOption) (*OK, error)
	NewEnginesState(ctx context.Context, in *Engines, opts ...grpc.CallOption) (*OK, error)
}

type conveyorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConveyorServiceClient(cc grpc.ClientConnInterface) ConveyorServiceClient {
	return &conveyorServiceClient{cc}
}

func (c *conveyorServiceClient) NewPosition(ctx context.Context, in *Position, opts ...grpc.CallOption) (*OK, error) {
	out := new(OK)
	err := c.cc.Invoke(ctx, "/conveyor.ConveyorService/NewPosition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conveyorServiceClient) NewEnginesState(ctx context.Context, in *Engines, opts ...grpc.CallOption) (*OK, error) {
	out := new(OK)
	err := c.cc.Invoke(ctx, "/conveyor.ConveyorService/NewEnginesState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConveyorServiceServer is the server API for ConveyorService service.
type ConveyorServiceServer interface {
	NewPosition(context.Context, *Position) (*OK, error)
	NewEnginesState(context.Context, *Engines) (*OK, error)
}

// UnimplementedConveyorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedConveyorServiceServer struct {
}

func (*UnimplementedConveyorServiceServer) NewPosition(context.Context, *Position) (*OK, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewPosition not implemented")
}
func (*UnimplementedConveyorServiceServer) NewEnginesState(context.Context, *Engines) (*OK, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewEnginesState not implemented")
}

func RegisterConveyorServiceServer(s *grpc.Server, srv ConveyorServiceServer) {
	s.RegisterService(&_ConveyorService_serviceDesc, srv)
}

func _ConveyorService_NewPosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Position)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConveyorServiceServer).NewPosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/conveyor.ConveyorService/NewPosition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConveyorServiceServer).NewPosition(ctx, req.(*Position))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConveyorService_NewEnginesState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Engines)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConveyorServiceServer).NewEnginesState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/conveyor.ConveyorService/NewEnginesState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConveyorServiceServer).NewEnginesState(ctx, req.(*Engines))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConveyorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "conveyor.ConveyorService",
	HandlerType: (*ConveyorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewPosition",
			Handler:    _ConveyorService_NewPosition_Handler,
		},
		{
			MethodName: "NewEnginesState",
			Handler:    _ConveyorService_NewEnginesState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "conveyor.proto",
}
