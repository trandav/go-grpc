// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: calculator/calculatorpb/calculator.proto

package calculatorpb

import (
	context "context"
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

type CalculatorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *CalculatorRequest) Reset() {
	*x = CalculatorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculatorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculatorRequest) ProtoMessage() {}

func (x *CalculatorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculatorRequest.ProtoReflect.Descriptor instead.
func (*CalculatorRequest) Descriptor() ([]byte, []int) {
	return file_calculator_calculatorpb_calculator_proto_rawDescGZIP(), []int{0}
}

func (x *CalculatorRequest) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *CalculatorRequest) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type CalculatorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sum int32 `protobuf:"varint,1,opt,name=sum,proto3" json:"sum,omitempty"`
}

func (x *CalculatorResponse) Reset() {
	*x = CalculatorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculatorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculatorResponse) ProtoMessage() {}

func (x *CalculatorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculatorResponse.ProtoReflect.Descriptor instead.
func (*CalculatorResponse) Descriptor() ([]byte, []int) {
	return file_calculator_calculatorpb_calculator_proto_rawDescGZIP(), []int{1}
}

func (x *CalculatorResponse) GetSum() int32 {
	if x != nil {
		return x.Sum
	}
	return 0
}

type CalculatorStreamingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
}

func (x *CalculatorStreamingRequest) Reset() {
	*x = CalculatorStreamingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculatorStreamingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculatorStreamingRequest) ProtoMessage() {}

func (x *CalculatorStreamingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculatorStreamingRequest.ProtoReflect.Descriptor instead.
func (*CalculatorStreamingRequest) Descriptor() ([]byte, []int) {
	return file_calculator_calculatorpb_calculator_proto_rawDescGZIP(), []int{2}
}

func (x *CalculatorStreamingRequest) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

type CalculatorStreamingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
}

func (x *CalculatorStreamingResponse) Reset() {
	*x = CalculatorStreamingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculatorStreamingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculatorStreamingResponse) ProtoMessage() {}

func (x *CalculatorStreamingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculatorStreamingResponse.ProtoReflect.Descriptor instead.
func (*CalculatorStreamingResponse) Descriptor() ([]byte, []int) {
	return file_calculator_calculatorpb_calculator_proto_rawDescGZIP(), []int{3}
}

func (x *CalculatorStreamingResponse) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

var File_calculator_calculatorpb_calculator_proto protoreflect.FileDescriptor

var file_calculator_calculatorpb_calculator_proto_rawDesc = []byte{
	0x0a, 0x28, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x63, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x2f, 0x0a, 0x11, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x78,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x22, 0x26, 0x0a, 0x12, 0x43, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x73, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x73, 0x75, 0x6d, 0x22,
	0x2a, 0x0a, 0x1a, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a,
	0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x22, 0x2b, 0x0a, 0x1b, 0x43,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x32, 0xcd, 0x01, 0x0a, 0x11, 0x43, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a,
	0x0a, 0x09, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x63, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61,
	0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6c, 0x0a, 0x17, 0x43, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x26, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e,
	0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x1b, 0x5a, 0x19, 0x2e, 0x2f, 0x63, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61,
	0x74, 0x6f, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_calculator_calculatorpb_calculator_proto_rawDescOnce sync.Once
	file_calculator_calculatorpb_calculator_proto_rawDescData = file_calculator_calculatorpb_calculator_proto_rawDesc
)

func file_calculator_calculatorpb_calculator_proto_rawDescGZIP() []byte {
	file_calculator_calculatorpb_calculator_proto_rawDescOnce.Do(func() {
		file_calculator_calculatorpb_calculator_proto_rawDescData = protoimpl.X.CompressGZIP(file_calculator_calculatorpb_calculator_proto_rawDescData)
	})
	return file_calculator_calculatorpb_calculator_proto_rawDescData
}

var file_calculator_calculatorpb_calculator_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_calculator_calculatorpb_calculator_proto_goTypes = []interface{}{
	(*CalculatorRequest)(nil),           // 0: calculator.CalculatorRequest
	(*CalculatorResponse)(nil),          // 1: calculator.CalculatorResponse
	(*CalculatorStreamingRequest)(nil),  // 2: calculator.CalculatorStreamingRequest
	(*CalculatorStreamingResponse)(nil), // 3: calculator.CalculatorStreamingResponse
}
var file_calculator_calculatorpb_calculator_proto_depIdxs = []int32{
	0, // 0: calculator.CalculatorService.Calculate:input_type -> calculator.CalculatorRequest
	2, // 1: calculator.CalculatorService.CalculatePrimeStreaming:input_type -> calculator.CalculatorStreamingRequest
	1, // 2: calculator.CalculatorService.Calculate:output_type -> calculator.CalculatorResponse
	3, // 3: calculator.CalculatorService.CalculatePrimeStreaming:output_type -> calculator.CalculatorStreamingResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_calculator_calculatorpb_calculator_proto_init() }
func file_calculator_calculatorpb_calculator_proto_init() {
	if File_calculator_calculatorpb_calculator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calculator_calculatorpb_calculator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculatorRequest); i {
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
		file_calculator_calculatorpb_calculator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculatorResponse); i {
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
		file_calculator_calculatorpb_calculator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculatorStreamingRequest); i {
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
		file_calculator_calculatorpb_calculator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculatorStreamingResponse); i {
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
			RawDescriptor: file_calculator_calculatorpb_calculator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calculator_calculatorpb_calculator_proto_goTypes,
		DependencyIndexes: file_calculator_calculatorpb_calculator_proto_depIdxs,
		MessageInfos:      file_calculator_calculatorpb_calculator_proto_msgTypes,
	}.Build()
	File_calculator_calculatorpb_calculator_proto = out.File
	file_calculator_calculatorpb_calculator_proto_rawDesc = nil
	file_calculator_calculatorpb_calculator_proto_goTypes = nil
	file_calculator_calculatorpb_calculator_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	// Unary
	Calculate(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error)
	CalculatePrimeStreaming(ctx context.Context, in *CalculatorStreamingRequest, opts ...grpc.CallOption) (CalculatorService_CalculatePrimeStreamingClient, error)
}

type calculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorServiceClient(cc grpc.ClientConnInterface) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Calculate(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error) {
	out := new(CalculatorResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Calculate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) CalculatePrimeStreaming(ctx context.Context, in *CalculatorStreamingRequest, opts ...grpc.CallOption) (CalculatorService_CalculatePrimeStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[0], "/calculator.CalculatorService/CalculatePrimeStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceCalculatePrimeStreamingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_CalculatePrimeStreamingClient interface {
	Recv() (*CalculatorStreamingResponse, error)
	grpc.ClientStream
}

type calculatorServiceCalculatePrimeStreamingClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceCalculatePrimeStreamingClient) Recv() (*CalculatorStreamingResponse, error) {
	m := new(CalculatorStreamingResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	// Unary
	Calculate(context.Context, *CalculatorRequest) (*CalculatorResponse, error)
	CalculatePrimeStreaming(*CalculatorStreamingRequest, CalculatorService_CalculatePrimeStreamingServer) error
}

// UnimplementedCalculatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (*UnimplementedCalculatorServiceServer) Calculate(context.Context, *CalculatorRequest) (*CalculatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Calculate not implemented")
}
func (*UnimplementedCalculatorServiceServer) CalculatePrimeStreaming(*CalculatorStreamingRequest, CalculatorService_CalculatePrimeStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method CalculatePrimeStreaming not implemented")
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_Calculate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Calculate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Calculate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Calculate(ctx, req.(*CalculatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_CalculatePrimeStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CalculatorStreamingRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).CalculatePrimeStreaming(m, &calculatorServiceCalculatePrimeStreamingServer{stream})
}

type CalculatorService_CalculatePrimeStreamingServer interface {
	Send(*CalculatorStreamingResponse) error
	grpc.ServerStream
}

type calculatorServiceCalculatePrimeStreamingServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceCalculatePrimeStreamingServer) Send(m *CalculatorStreamingResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Calculate",
			Handler:    _CalculatorService_Calculate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CalculatePrimeStreaming",
			Handler:       _CalculatorService_CalculatePrimeStreaming_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "calculator/calculatorpb/calculator.proto",
}