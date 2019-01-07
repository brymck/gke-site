// Code generated by protoc-gen-go. DO NOT EDIT.
// source: demo.proto

package gkesite

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GreetingRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetingRequest) Reset()         { *m = GreetingRequest{} }
func (m *GreetingRequest) String() string { return proto.CompactTextString(m) }
func (*GreetingRequest) ProtoMessage()    {}
func (*GreetingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{0}
}

func (m *GreetingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetingRequest.Unmarshal(m, b)
}
func (m *GreetingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetingRequest.Marshal(b, m, deterministic)
}
func (m *GreetingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetingRequest.Merge(m, src)
}
func (m *GreetingRequest) XXX_Size() int {
	return xxx_messageInfo_GreetingRequest.Size(m)
}
func (m *GreetingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetingRequest proto.InternalMessageInfo

func (m *GreetingRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Greeting struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Greeting) Reset()         { *m = Greeting{} }
func (m *Greeting) String() string { return proto.CompactTextString(m) }
func (*Greeting) ProtoMessage()    {}
func (*Greeting) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{1}
}

func (m *Greeting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Greeting.Unmarshal(m, b)
}
func (m *Greeting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Greeting.Marshal(b, m, deterministic)
}
func (m *Greeting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Greeting.Merge(m, src)
}
func (m *Greeting) XXX_Size() int {
	return xxx_messageInfo_Greeting.Size(m)
}
func (m *Greeting) XXX_DiscardUnknown() {
	xxx_messageInfo_Greeting.DiscardUnknown(m)
}

var xxx_messageInfo_Greeting proto.InternalMessageInfo

func (m *Greeting) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type SquareRequest struct {
	Number               float32  `protobuf:"fixed32,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SquareRequest) Reset()         { *m = SquareRequest{} }
func (m *SquareRequest) String() string { return proto.CompactTextString(m) }
func (*SquareRequest) ProtoMessage()    {}
func (*SquareRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{2}
}

func (m *SquareRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SquareRequest.Unmarshal(m, b)
}
func (m *SquareRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SquareRequest.Marshal(b, m, deterministic)
}
func (m *SquareRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SquareRequest.Merge(m, src)
}
func (m *SquareRequest) XXX_Size() int {
	return xxx_messageInfo_SquareRequest.Size(m)
}
func (m *SquareRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SquareRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SquareRequest proto.InternalMessageInfo

func (m *SquareRequest) GetNumber() float32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type SquareResponse struct {
	Number               float32  `protobuf:"fixed32,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SquareResponse) Reset()         { *m = SquareResponse{} }
func (m *SquareResponse) String() string { return proto.CompactTextString(m) }
func (*SquareResponse) ProtoMessage()    {}
func (*SquareResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{3}
}

func (m *SquareResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SquareResponse.Unmarshal(m, b)
}
func (m *SquareResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SquareResponse.Marshal(b, m, deterministic)
}
func (m *SquareResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SquareResponse.Merge(m, src)
}
func (m *SquareResponse) XXX_Size() int {
	return xxx_messageInfo_SquareResponse.Size(m)
}
func (m *SquareResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SquareResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SquareResponse proto.InternalMessageInfo

func (m *SquareResponse) GetNumber() float32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type CountRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CountRequest) Reset()         { *m = CountRequest{} }
func (m *CountRequest) String() string { return proto.CompactTextString(m) }
func (*CountRequest) ProtoMessage()    {}
func (*CountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{4}
}

func (m *CountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CountRequest.Unmarshal(m, b)
}
func (m *CountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CountRequest.Marshal(b, m, deterministic)
}
func (m *CountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CountRequest.Merge(m, src)
}
func (m *CountRequest) XXX_Size() int {
	return xxx_messageInfo_CountRequest.Size(m)
}
func (m *CountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CountRequest proto.InternalMessageInfo

type CountResponse struct {
	Count                int32    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CountResponse) Reset()         { *m = CountResponse{} }
func (m *CountResponse) String() string { return proto.CompactTextString(m) }
func (*CountResponse) ProtoMessage()    {}
func (*CountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{5}
}

func (m *CountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CountResponse.Unmarshal(m, b)
}
func (m *CountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CountResponse.Marshal(b, m, deterministic)
}
func (m *CountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CountResponse.Merge(m, src)
}
func (m *CountResponse) XXX_Size() int {
	return xxx_messageInfo_CountResponse.Size(m)
}
func (m *CountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CountResponse proto.InternalMessageInfo

func (m *CountResponse) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*GreetingRequest)(nil), "gkesite.GreetingRequest")
	proto.RegisterType((*Greeting)(nil), "gkesite.Greeting")
	proto.RegisterType((*SquareRequest)(nil), "gkesite.SquareRequest")
	proto.RegisterType((*SquareResponse)(nil), "gkesite.SquareResponse")
	proto.RegisterType((*CountRequest)(nil), "gkesite.CountRequest")
	proto.RegisterType((*CountResponse)(nil), "gkesite.CountResponse")
}

func init() { proto.RegisterFile("demo.proto", fileDescriptor_ca53982754088a9d) }

var fileDescriptor_ca53982754088a9d = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x4f, 0x4b, 0x33, 0x31,
	0x18, 0xc4, 0xdb, 0x97, 0xb7, 0xff, 0x1e, 0xdb, 0x8a, 0x41, 0xd7, 0xb2, 0x27, 0x09, 0x16, 0x7b,
	0xda, 0xc3, 0x7a, 0x54, 0xbc, 0x78, 0x58, 0x41, 0x41, 0xd8, 0x7e, 0x82, 0x6d, 0x1d, 0x96, 0xc5,
	0x6e, 0xd2, 0x26, 0x59, 0x3f, 0xbf, 0x98, 0xcd, 0x53, 0xb5, 0xc5, 0x5b, 0x66, 0x32, 0x33, 0xe4,
	0x47, 0x88, 0xde, 0x50, 0xeb, 0x64, 0x6b, 0xb4, 0xd3, 0x62, 0x50, 0xbe, 0xc3, 0x56, 0x0e, 0x72,
	0x4e, 0xa7, 0x99, 0x01, 0x5c, 0xa5, 0xca, 0x1c, 0xbb, 0x06, 0xd6, 0x09, 0x41, 0xff, 0x55, 0x51,
	0x63, 0xd6, 0xbd, 0xea, 0x2e, 0x46, 0xb9, 0x3f, 0xcb, 0x6b, 0x1a, 0x72, 0x4c, 0xcc, 0x68, 0x50,
	0xc3, 0xda, 0xa2, 0xe4, 0x08, 0x4b, 0x79, 0x43, 0x93, 0xe5, 0xae, 0x29, 0x0c, 0x78, 0x2a, 0xa2,
	0xbe, 0x6a, 0xea, 0x15, 0x8c, 0x4f, 0xfe, 0xcb, 0x83, 0x92, 0x0b, 0x9a, 0x72, 0xd0, 0x6e, 0xb5,
	0xb2, 0xf8, 0x33, 0x39, 0xa5, 0xf1, 0xa3, 0x6e, 0x94, 0x0b, 0x8b, 0x72, 0x4e, 0x93, 0xa0, 0x43,
	0xf1, 0x9c, 0x7a, 0xeb, 0x2f, 0xc3, 0xf7, 0x7a, 0x79, 0x2b, 0xd2, 0x17, 0x1a, 0x3f, 0x61, 0xb3,
	0xd1, 0x4b, 0x98, 0x8f, 0x6a, 0x0d, 0x71, 0x4f, 0x27, 0x19, 0xdc, 0x37, 0x42, 0x12, 0xf8, 0x93,
	0x03, 0xf8, 0xf8, 0xec, 0xe8, 0x46, 0x76, 0xd2, 0x57, 0xe6, 0xe2, 0xb9, 0x07, 0x1a, 0x65, 0x70,
	0xad, 0x27, 0xa2, 0x7d, 0xe5, 0x17, 0x7c, 0x7c, 0x79, 0xe4, 0xb7, 0x4f, 0x96, 0x9d, 0xf4, 0x39,
	0x50, 0xf1, 0xde, 0x1d, 0x0d, 0x33, 0x38, 0x6f, 0x89, 0x8b, 0x7d, 0xed, 0x27, 0x78, 0x1c, 0x1d,
	0xda, 0x3c, 0xb6, 0xea, 0xfb, 0x2f, 0xbd, 0xfd, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xee, 0xc6, 0xb5,
	0xaf, 0xe0, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloServiceClient interface {
	GetGreeting(ctx context.Context, in *GreetingRequest, opts ...grpc.CallOption) (*Greeting, error)
}

type helloServiceClient struct {
	cc *grpc.ClientConn
}

func NewHelloServiceClient(cc *grpc.ClientConn) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) GetGreeting(ctx context.Context, in *GreetingRequest, opts ...grpc.CallOption) (*Greeting, error) {
	out := new(Greeting)
	err := c.cc.Invoke(ctx, "/gkesite.HelloService/GetGreeting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServiceServer is the server API for HelloService service.
type HelloServiceServer interface {
	GetGreeting(context.Context, *GreetingRequest) (*Greeting, error)
}

func RegisterHelloServiceServer(s *grpc.Server, srv HelloServiceServer) {
	s.RegisterService(&_HelloService_serviceDesc, srv)
}

func _HelloService_GetGreeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).GetGreeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gkesite.HelloService/GetGreeting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).GetGreeting(ctx, req.(*GreetingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gkesite.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGreeting",
			Handler:    _HelloService_GetGreeting_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo.proto",
}

// SquareServiceClient is the client API for SquareService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SquareServiceClient interface {
	GetSquare(ctx context.Context, in *SquareRequest, opts ...grpc.CallOption) (*SquareResponse, error)
}

type squareServiceClient struct {
	cc *grpc.ClientConn
}

func NewSquareServiceClient(cc *grpc.ClientConn) SquareServiceClient {
	return &squareServiceClient{cc}
}

func (c *squareServiceClient) GetSquare(ctx context.Context, in *SquareRequest, opts ...grpc.CallOption) (*SquareResponse, error) {
	out := new(SquareResponse)
	err := c.cc.Invoke(ctx, "/gkesite.SquareService/GetSquare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SquareServiceServer is the server API for SquareService service.
type SquareServiceServer interface {
	GetSquare(context.Context, *SquareRequest) (*SquareResponse, error)
}

func RegisterSquareServiceServer(s *grpc.Server, srv SquareServiceServer) {
	s.RegisterService(&_SquareService_serviceDesc, srv)
}

func _SquareService_GetSquare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SquareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SquareServiceServer).GetSquare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gkesite.SquareService/GetSquare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SquareServiceServer).GetSquare(ctx, req.(*SquareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SquareService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gkesite.SquareService",
	HandlerType: (*SquareServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSquare",
			Handler:    _SquareService_GetSquare_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo.proto",
}

// CountServiceClient is the client API for CountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CountServiceClient interface {
	GetCount(ctx context.Context, in *CountRequest, opts ...grpc.CallOption) (*CountResponse, error)
}

type countServiceClient struct {
	cc *grpc.ClientConn
}

func NewCountServiceClient(cc *grpc.ClientConn) CountServiceClient {
	return &countServiceClient{cc}
}

func (c *countServiceClient) GetCount(ctx context.Context, in *CountRequest, opts ...grpc.CallOption) (*CountResponse, error) {
	out := new(CountResponse)
	err := c.cc.Invoke(ctx, "/gkesite.CountService/GetCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CountServiceServer is the server API for CountService service.
type CountServiceServer interface {
	GetCount(context.Context, *CountRequest) (*CountResponse, error)
}

func RegisterCountServiceServer(s *grpc.Server, srv CountServiceServer) {
	s.RegisterService(&_CountService_serviceDesc, srv)
}

func _CountService_GetCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountServiceServer).GetCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gkesite.CountService/GetCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountServiceServer).GetCount(ctx, req.(*CountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CountService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gkesite.CountService",
	HandlerType: (*CountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCount",
			Handler:    _CountService_GetCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo.proto",
}
