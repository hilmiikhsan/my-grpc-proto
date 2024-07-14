// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: proto/hello/hello.proto

package hello

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	HelloService_SayHello_FullMethodName           = "/hello.HelloService/SayHello"
	HelloService_SayManyHellos_FullMethodName      = "/hello.HelloService/SayManyHellos"
	HelloService_SayHelloToEveryone_FullMethodName = "/hello.HelloService/SayHelloToEveryone"
	HelloService_SayHelloContnuous_FullMethodName  = "/hello.HelloService/SayHelloContnuous"
)

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloServiceClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	SayManyHellos(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (HelloService_SayManyHellosClient, error)
	SayHelloToEveryone(ctx context.Context, opts ...grpc.CallOption) (HelloService_SayHelloToEveryoneClient, error)
	SayHelloContnuous(ctx context.Context, opts ...grpc.CallOption) (HelloService_SayHelloContnuousClient, error)
}

type helloServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloServiceClient(cc grpc.ClientConnInterface) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, HelloService_SayHello_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServiceClient) SayManyHellos(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (HelloService_SayManyHellosClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &HelloService_ServiceDesc.Streams[0], HelloService_SayManyHellos_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceSayManyHellosClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HelloService_SayManyHellosClient interface {
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type helloServiceSayManyHellosClient struct {
	grpc.ClientStream
}

func (x *helloServiceSayManyHellosClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloServiceClient) SayHelloToEveryone(ctx context.Context, opts ...grpc.CallOption) (HelloService_SayHelloToEveryoneClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &HelloService_ServiceDesc.Streams[1], HelloService_SayHelloToEveryone_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceSayHelloToEveryoneClient{ClientStream: stream}
	return x, nil
}

type HelloService_SayHelloToEveryoneClient interface {
	Send(*HelloRequest) error
	CloseAndRecv() (*HelloResponse, error)
	grpc.ClientStream
}

type helloServiceSayHelloToEveryoneClient struct {
	grpc.ClientStream
}

func (x *helloServiceSayHelloToEveryoneClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloServiceSayHelloToEveryoneClient) CloseAndRecv() (*HelloResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloServiceClient) SayHelloContnuous(ctx context.Context, opts ...grpc.CallOption) (HelloService_SayHelloContnuousClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &HelloService_ServiceDesc.Streams[2], HelloService_SayHelloContnuous_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceSayHelloContnuousClient{ClientStream: stream}
	return x, nil
}

type HelloService_SayHelloContnuousClient interface {
	Send(*HelloRequest) error
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type helloServiceSayHelloContnuousClient struct {
	grpc.ClientStream
}

func (x *helloServiceSayHelloContnuousClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloServiceSayHelloContnuousClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloServiceServer is the server API for HelloService service.
// All implementations must embed UnimplementedHelloServiceServer
// for forward compatibility
type HelloServiceServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloResponse, error)
	SayManyHellos(*HelloRequest, HelloService_SayManyHellosServer) error
	SayHelloToEveryone(HelloService_SayHelloToEveryoneServer) error
	SayHelloContnuous(HelloService_SayHelloContnuousServer) error
	mustEmbedUnimplementedHelloServiceServer()
}

// UnimplementedHelloServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHelloServiceServer struct {
}

func (UnimplementedHelloServiceServer) SayHello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedHelloServiceServer) SayManyHellos(*HelloRequest, HelloService_SayManyHellosServer) error {
	return status.Errorf(codes.Unimplemented, "method SayManyHellos not implemented")
}
func (UnimplementedHelloServiceServer) SayHelloToEveryone(HelloService_SayHelloToEveryoneServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloToEveryone not implemented")
}
func (UnimplementedHelloServiceServer) SayHelloContnuous(HelloService_SayHelloContnuousServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloContnuous not implemented")
}
func (UnimplementedHelloServiceServer) mustEmbedUnimplementedHelloServiceServer() {}

// UnsafeHelloServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloServiceServer will
// result in compilation errors.
type UnsafeHelloServiceServer interface {
	mustEmbedUnimplementedHelloServiceServer()
}

func RegisterHelloServiceServer(s grpc.ServiceRegistrar, srv HelloServiceServer) {
	s.RegisterService(&HelloService_ServiceDesc, srv)
}

func _HelloService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HelloService_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloService_SayManyHellos_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HelloServiceServer).SayManyHellos(m, &helloServiceSayManyHellosServer{ServerStream: stream})
}

type HelloService_SayManyHellosServer interface {
	Send(*HelloResponse) error
	grpc.ServerStream
}

type helloServiceSayManyHellosServer struct {
	grpc.ServerStream
}

func (x *helloServiceSayManyHellosServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _HelloService_SayHelloToEveryone_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServiceServer).SayHelloToEveryone(&helloServiceSayHelloToEveryoneServer{ServerStream: stream})
}

type HelloService_SayHelloToEveryoneServer interface {
	SendAndClose(*HelloResponse) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type helloServiceSayHelloToEveryoneServer struct {
	grpc.ServerStream
}

func (x *helloServiceSayHelloToEveryoneServer) SendAndClose(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloServiceSayHelloToEveryoneServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _HelloService_SayHelloContnuous_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServiceServer).SayHelloContnuous(&helloServiceSayHelloContnuousServer{ServerStream: stream})
}

type HelloService_SayHelloContnuousServer interface {
	Send(*HelloResponse) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type helloServiceSayHelloContnuousServer struct {
	grpc.ServerStream
}

func (x *helloServiceSayHelloContnuousServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloServiceSayHelloContnuousServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloService_ServiceDesc is the grpc.ServiceDesc for HelloService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HelloService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hello.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _HelloService_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayManyHellos",
			Handler:       _HelloService_SayManyHellos_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SayHelloToEveryone",
			Handler:       _HelloService_SayHelloToEveryone_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SayHelloContnuous",
			Handler:       _HelloService_SayHelloContnuous_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/hello/hello.proto",
}
