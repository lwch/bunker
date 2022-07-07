// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: bunker.proto

package network

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BunkerClient is the client API for Bunker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BunkerClient interface {
	// shell
	RunShell(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*RunShellArguments, error)
	ShellResize(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ShellResizeArguments, error)
	ShellForward(ctx context.Context, opts ...grpc.CallOption) (Bunker_ShellForwardClient, error)
}

type bunkerClient struct {
	cc grpc.ClientConnInterface
}

func NewBunkerClient(cc grpc.ClientConnInterface) BunkerClient {
	return &bunkerClient{cc}
}

func (c *bunkerClient) RunShell(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*RunShellArguments, error) {
	out := new(RunShellArguments)
	err := c.cc.Invoke(ctx, "/Bunker/RunShell", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bunkerClient) ShellResize(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ShellResizeArguments, error) {
	out := new(ShellResizeArguments)
	err := c.cc.Invoke(ctx, "/Bunker/ShellResize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bunkerClient) ShellForward(ctx context.Context, opts ...grpc.CallOption) (Bunker_ShellForwardClient, error) {
	stream, err := c.cc.NewStream(ctx, &Bunker_ServiceDesc.Streams[0], "/Bunker/ShellForward", opts...)
	if err != nil {
		return nil, err
	}
	x := &bunkerShellForwardClient{stream}
	return x, nil
}

type Bunker_ShellForwardClient interface {
	Send(*wrapperspb.BytesValue) error
	Recv() (*wrapperspb.BytesValue, error)
	grpc.ClientStream
}

type bunkerShellForwardClient struct {
	grpc.ClientStream
}

func (x *bunkerShellForwardClient) Send(m *wrapperspb.BytesValue) error {
	return x.ClientStream.SendMsg(m)
}

func (x *bunkerShellForwardClient) Recv() (*wrapperspb.BytesValue, error) {
	m := new(wrapperspb.BytesValue)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BunkerServer is the server API for Bunker service.
// All implementations must embed UnimplementedBunkerServer
// for forward compatibility
type BunkerServer interface {
	// shell
	RunShell(context.Context, *emptypb.Empty) (*RunShellArguments, error)
	ShellResize(context.Context, *emptypb.Empty) (*ShellResizeArguments, error)
	ShellForward(Bunker_ShellForwardServer) error
	mustEmbedUnimplementedBunkerServer()
}

// UnimplementedBunkerServer must be embedded to have forward compatible implementations.
type UnimplementedBunkerServer struct {
}

func (UnimplementedBunkerServer) RunShell(context.Context, *emptypb.Empty) (*RunShellArguments, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunShell not implemented")
}
func (UnimplementedBunkerServer) ShellResize(context.Context, *emptypb.Empty) (*ShellResizeArguments, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShellResize not implemented")
}
func (UnimplementedBunkerServer) ShellForward(Bunker_ShellForwardServer) error {
	return status.Errorf(codes.Unimplemented, "method ShellForward not implemented")
}
func (UnimplementedBunkerServer) mustEmbedUnimplementedBunkerServer() {}

// UnsafeBunkerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BunkerServer will
// result in compilation errors.
type UnsafeBunkerServer interface {
	mustEmbedUnimplementedBunkerServer()
}

func RegisterBunkerServer(s grpc.ServiceRegistrar, srv BunkerServer) {
	s.RegisterService(&Bunker_ServiceDesc, srv)
}

func _Bunker_RunShell_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BunkerServer).RunShell(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Bunker/RunShell",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BunkerServer).RunShell(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bunker_ShellResize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BunkerServer).ShellResize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Bunker/ShellResize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BunkerServer).ShellResize(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bunker_ShellForward_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BunkerServer).ShellForward(&bunkerShellForwardServer{stream})
}

type Bunker_ShellForwardServer interface {
	Send(*wrapperspb.BytesValue) error
	Recv() (*wrapperspb.BytesValue, error)
	grpc.ServerStream
}

type bunkerShellForwardServer struct {
	grpc.ServerStream
}

func (x *bunkerShellForwardServer) Send(m *wrapperspb.BytesValue) error {
	return x.ServerStream.SendMsg(m)
}

func (x *bunkerShellForwardServer) Recv() (*wrapperspb.BytesValue, error) {
	m := new(wrapperspb.BytesValue)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Bunker_ServiceDesc is the grpc.ServiceDesc for Bunker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Bunker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Bunker",
	HandlerType: (*BunkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RunShell",
			Handler:    _Bunker_RunShell_Handler,
		},
		{
			MethodName: "ShellResize",
			Handler:    _Bunker_ShellResize_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ShellForward",
			Handler:       _Bunker_ShellForward_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "bunker.proto",
}
