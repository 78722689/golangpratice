// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterClient interface {
	Say1(ctx context.Context, in *SayWhat, opts ...grpc.CallOption) (Greeter_Say1Client, error)
	Say2(ctx context.Context, opts ...grpc.CallOption) (Greeter_Say2Client, error)
	Say3(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	Say4(ctx context.Context, opts ...grpc.CallOption) (Greeter_Say4Client, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) Say1(ctx context.Context, in *SayWhat, opts ...grpc.CallOption) (Greeter_Say1Client, error) {
	stream, err := c.cc.NewStream(ctx, &Greeter_ServiceDesc.Streams[0], "/proto.Greeter/Say1", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterSay1Client{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Greeter_Say1Client interface {
	Recv() (*ReplyWhat, error)
	grpc.ClientStream
}

type greeterSay1Client struct {
	grpc.ClientStream
}

func (x *greeterSay1Client) Recv() (*ReplyWhat, error) {
	m := new(ReplyWhat)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterClient) Say2(ctx context.Context, opts ...grpc.CallOption) (Greeter_Say2Client, error) {
	stream, err := c.cc.NewStream(ctx, &Greeter_ServiceDesc.Streams[1], "/proto.Greeter/Say2", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterSay2Client{stream}
	return x, nil
}

type Greeter_Say2Client interface {
	Send(*SayWhat) error
	CloseAndRecv() (*ReplyWhat, error)
	grpc.ClientStream
}

type greeterSay2Client struct {
	grpc.ClientStream
}

func (x *greeterSay2Client) Send(m *SayWhat) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterSay2Client) CloseAndRecv() (*ReplyWhat, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ReplyWhat)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterClient) Say3(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/proto.Greeter/Say3", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) Say4(ctx context.Context, opts ...grpc.CallOption) (Greeter_Say4Client, error) {
	stream, err := c.cc.NewStream(ctx, &Greeter_ServiceDesc.Streams[2], "/proto.Greeter/Say4", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterSay4Client{stream}
	return x, nil
}

type Greeter_Say4Client interface {
	Send(*SayWhat) error
	Recv() (*ReplyWhat, error)
	grpc.ClientStream
}

type greeterSay4Client struct {
	grpc.ClientStream
}

func (x *greeterSay4Client) Send(m *SayWhat) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterSay4Client) Recv() (*ReplyWhat, error) {
	m := new(ReplyWhat)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreeterServer is the server API for Greeter service.
// All implementations must embed UnimplementedGreeterServer
// for forward compatibility
type GreeterServer interface {
	Say1(*SayWhat, Greeter_Say1Server) error
	Say2(Greeter_Say2Server) error
	Say3(context.Context, *HelloRequest) (*HelloReply, error)
	Say4(Greeter_Say4Server) error
	mustEmbedUnimplementedGreeterServer()
}

// UnimplementedGreeterServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (UnimplementedGreeterServer) Say1(*SayWhat, Greeter_Say1Server) error {
	return status.Errorf(codes.Unimplemented, "method Say1 not implemented")
}
func (UnimplementedGreeterServer) Say2(Greeter_Say2Server) error {
	return status.Errorf(codes.Unimplemented, "method Say2 not implemented")
}
func (UnimplementedGreeterServer) Say3(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Say3 not implemented")
}
func (UnimplementedGreeterServer) Say4(Greeter_Say4Server) error {
	return status.Errorf(codes.Unimplemented, "method Say4 not implemented")
}
func (UnimplementedGreeterServer) mustEmbedUnimplementedGreeterServer() {}

// UnsafeGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServer will
// result in compilation errors.
type UnsafeGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
}

func RegisterGreeterServer(s grpc.ServiceRegistrar, srv GreeterServer) {
	s.RegisterService(&Greeter_ServiceDesc, srv)
}

func _Greeter_Say1_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SayWhat)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreeterServer).Say1(m, &greeterSay1Server{stream})
}

type Greeter_Say1Server interface {
	Send(*ReplyWhat) error
	grpc.ServerStream
}

type greeterSay1Server struct {
	grpc.ServerStream
}

func (x *greeterSay1Server) Send(m *ReplyWhat) error {
	return x.ServerStream.SendMsg(m)
}

func _Greeter_Say2_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).Say2(&greeterSay2Server{stream})
}

type Greeter_Say2Server interface {
	SendAndClose(*ReplyWhat) error
	Recv() (*SayWhat, error)
	grpc.ServerStream
}

type greeterSay2Server struct {
	grpc.ServerStream
}

func (x *greeterSay2Server) SendAndClose(m *ReplyWhat) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterSay2Server) Recv() (*SayWhat, error) {
	m := new(SayWhat)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Greeter_Say3_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Say3(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Greeter/Say3",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Say3(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_Say4_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).Say4(&greeterSay4Server{stream})
}

type Greeter_Say4Server interface {
	Send(*ReplyWhat) error
	Recv() (*SayWhat, error)
	grpc.ServerStream
}

type greeterSay4Server struct {
	grpc.ServerStream
}

func (x *greeterSay4Server) Send(m *ReplyWhat) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterSay4Server) Recv() (*SayWhat, error) {
	m := new(SayWhat)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Greeter_ServiceDesc is the grpc.ServiceDesc for Greeter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Say3",
			Handler:    _Greeter_Say3_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Say1",
			Handler:       _Greeter_Say1_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Say2",
			Handler:       _Greeter_Say2_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Say4",
			Handler:       _Greeter_Say4_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "rpc/proto/hello.proto",
}
