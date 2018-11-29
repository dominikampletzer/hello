// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloworld.proto

package helloworld

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type UUIDRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UUIDRequest) Reset()         { *m = UUIDRequest{} }
func (m *UUIDRequest) String() string { return proto.CompactTextString(m) }
func (*UUIDRequest) ProtoMessage()    {}
func (*UUIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{0}
}

func (m *UUIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UUIDRequest.Unmarshal(m, b)
}
func (m *UUIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UUIDRequest.Marshal(b, m, deterministic)
}
func (m *UUIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UUIDRequest.Merge(m, src)
}
func (m *UUIDRequest) XXX_Size() int {
	return xxx_messageInfo_UUIDRequest.Size(m)
}
func (m *UUIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UUIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UUIDRequest proto.InternalMessageInfo

type UUIDReply struct {
	Uuid                 int64    `protobuf:"varint,1,opt,name=Uuid,proto3" json:"Uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UUIDReply) Reset()         { *m = UUIDReply{} }
func (m *UUIDReply) String() string { return proto.CompactTextString(m) }
func (*UUIDReply) ProtoMessage()    {}
func (*UUIDReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{1}
}

func (m *UUIDReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UUIDReply.Unmarshal(m, b)
}
func (m *UUIDReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UUIDReply.Marshal(b, m, deterministic)
}
func (m *UUIDReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UUIDReply.Merge(m, src)
}
func (m *UUIDReply) XXX_Size() int {
	return xxx_messageInfo_UUIDReply.Size(m)
}
func (m *UUIDReply) XXX_DiscardUnknown() {
	xxx_messageInfo_UUIDReply.DiscardUnknown(m)
}

var xxx_messageInfo_UUIDReply proto.InternalMessageInfo

func (m *UUIDReply) GetUuid() int64 {
	if m != nil {
		return m.Uuid
	}
	return 0
}

// The request message containing the user's name.
type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{2}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{3}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*UUIDRequest)(nil), "helloworld.UUIDRequest")
	proto.RegisterType((*UUIDReply)(nil), "helloworld.UUIDReply")
	proto.RegisterType((*HelloRequest)(nil), "helloworld.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "helloworld.HelloReply")
}

func init() { proto.RegisterFile("helloworld.proto", fileDescriptor_17b8c58d586b62f2) }

var fileDescriptor_17b8c58d586b62f2 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x48, 0xcd, 0xc9,
	0xc9, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88,
	0x28, 0xf1, 0x72, 0x71, 0x87, 0x86, 0x7a, 0xba, 0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x28,
	0xc9, 0x73, 0x71, 0x42, 0xb8, 0x05, 0x39, 0x95, 0x42, 0x42, 0x5c, 0x2c, 0xa1, 0xa5, 0x99, 0x29,
	0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x60, 0xb6, 0x92, 0x12, 0x17, 0x8f, 0x07, 0x48, 0x37,
	0x54, 0x03, 0x48, 0x4d, 0x5e, 0x62, 0x6e, 0x2a, 0x58, 0x0d, 0x67, 0x10, 0x98, 0xad, 0xa4, 0xc6,
	0xc5, 0x05, 0x55, 0x03, 0x32, 0x45, 0x82, 0x8b, 0x3d, 0x37, 0xb5, 0xb8, 0x38, 0x31, 0x1d, 0xa6,
	0x08, 0xc6, 0x35, 0x3a, 0xcd, 0xc8, 0xc5, 0xee, 0x5e, 0x94, 0x9a, 0x5a, 0x92, 0x5a, 0x24, 0x64,
	0xc7, 0xc5, 0x11, 0x9c, 0x58, 0x09, 0xd6, 0x26, 0x24, 0xa1, 0x87, 0xe4, 0x64, 0x64, 0xdb, 0xa4,
	0xc4, 0xb0, 0xc8, 0x14, 0xe4, 0x54, 0x2a, 0x31, 0x08, 0x39, 0x73, 0xf1, 0xc2, 0xf4, 0x3b, 0xa6,
	0x27, 0x66, 0xe6, 0x91, 0x65, 0x88, 0x35, 0x17, 0xbb, 0x7b, 0x6a, 0x09, 0x28, 0x00, 0x84, 0xc4,
	0x91, 0x15, 0x21, 0x85, 0x90, 0x94, 0x28, 0xa6, 0x04, 0x58, 0xb3, 0x93, 0x01, 0x97, 0x74, 0x66,
	0xbe, 0x5e, 0x7a, 0x51, 0x41, 0xb2, 0x5e, 0x6a, 0x45, 0x62, 0x6e, 0x41, 0x4e, 0x6a, 0x31, 0x92,
	0x52, 0x27, 0x7e, 0xb0, 0x4d, 0xe1, 0x20, 0x76, 0x00, 0x28, 0x16, 0x02, 0x18, 0x93, 0xd8, 0xc0,
	0xd1, 0x61, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x01, 0x0b, 0x6e, 0xa2, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	SayHelloAgain(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	GetUUID(ctx context.Context, in *UUIDRequest, opts ...grpc.CallOption) (*UUIDReply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/helloworld.Greeter/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) SayHelloAgain(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/helloworld.Greeter/SayHelloAgain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) GetUUID(ctx context.Context, in *UUIDRequest, opts ...grpc.CallOption) (*UUIDReply, error) {
	out := new(UUIDReply)
	err := c.cc.Invoke(ctx, "/helloworld.Greeter/GetUUID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	SayHelloAgain(context.Context, *HelloRequest) (*HelloReply, error)
	GetUUID(context.Context, *UUIDRequest) (*UUIDReply, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_SayHelloAgain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHelloAgain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHelloAgain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHelloAgain(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_GetUUID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UUIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).GetUUID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/GetUUID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).GetUUID(ctx, req.(*UUIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
		{
			MethodName: "SayHelloAgain",
			Handler:    _Greeter_SayHelloAgain_Handler,
		},
		{
			MethodName: "GetUUID",
			Handler:    _Greeter_GetUUID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld.proto",
}
