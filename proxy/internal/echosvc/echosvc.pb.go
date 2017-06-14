// Code generated by protoc-gen-go. DO NOT EDIT.
// source: echosvc.proto

/*
Package echosvc is a generated protocol buffer package.

It is generated from these files:
	echosvc.proto

It has these top-level messages:
	Msg
*/
package echosvc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type Msg struct {
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
}

func (m *Msg) Reset()                    { *m = Msg{} }
func (m *Msg) String() string            { return proto.CompactTextString(m) }
func (*Msg) ProtoMessage()               {}
func (*Msg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Msg) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*Msg)(nil), "echosvc.Msg")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for EchoSvc service

type EchoSvcClient interface {
	Send(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Msg, error)
}

type echoSvcClient struct {
	cc *grpc.ClientConn
}

func NewEchoSvcClient(cc *grpc.ClientConn) EchoSvcClient {
	return &echoSvcClient{cc}
}

func (c *echoSvcClient) Send(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Msg, error) {
	out := new(Msg)
	err := grpc.Invoke(ctx, "/echosvc.EchoSvc/Send", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for EchoSvc service

type EchoSvcServer interface {
	Send(context.Context, *Msg) (*Msg, error)
}

func RegisterEchoSvcServer(s *grpc.Server, srv EchoSvcServer) {
	s.RegisterService(&_EchoSvc_serviceDesc, srv)
}

func _EchoSvc_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Msg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoSvcServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echosvc.EchoSvc/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoSvcServer).Send(ctx, req.(*Msg))
	}
	return interceptor(ctx, in, info, handler)
}

var _EchoSvc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "echosvc.EchoSvc",
	HandlerType: (*EchoSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _EchoSvc_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "echosvc.proto",
}

func init() { proto.RegisterFile("echosvc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 102 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x4d, 0xce, 0xc8,
	0x2f, 0x2e, 0x4b, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x24, 0xb9,
	0x98, 0x7d, 0x8b, 0xd3, 0x85, 0x84, 0xb8, 0x58, 0x4a, 0x52, 0x2b, 0x4a, 0x24, 0x18, 0x15, 0x18,
	0x35, 0x38, 0x83, 0xc0, 0x6c, 0x23, 0x5d, 0x2e, 0x76, 0xd7, 0xe4, 0x8c, 0xfc, 0xe0, 0xb2, 0x64,
	0x21, 0x25, 0x2e, 0x96, 0xe0, 0xd4, 0xbc, 0x14, 0x21, 0x1e, 0x3d, 0x98, 0x31, 0xbe, 0xc5, 0xe9,
	0x52, 0x28, 0xbc, 0x24, 0x36, 0xb0, 0xc9, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x45,
	0x07, 0x31, 0x6a, 0x00, 0x00, 0x00,
}