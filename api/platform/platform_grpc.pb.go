// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.1
// source: platform/platform.proto

package platform

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

const (
	Platform_CollectData_FullMethodName = "/api.platform.Platform/CollectData"
)

// PlatformClient is the client API for Platform service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlatformClient interface {
	CollectData(ctx context.Context, in *CollectDataRequest, opts ...grpc.CallOption) (*CollectDataReply, error)
}

type platformClient struct {
	cc grpc.ClientConnInterface
}

func NewPlatformClient(cc grpc.ClientConnInterface) PlatformClient {
	return &platformClient{cc}
}

func (c *platformClient) CollectData(ctx context.Context, in *CollectDataRequest, opts ...grpc.CallOption) (*CollectDataReply, error) {
	out := new(CollectDataReply)
	err := c.cc.Invoke(ctx, Platform_CollectData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlatformServer is the server API for Platform service.
// All implementations must embed UnimplementedPlatformServer
// for forward compatibility
type PlatformServer interface {
	CollectData(context.Context, *CollectDataRequest) (*CollectDataReply, error)
	mustEmbedUnimplementedPlatformServer()
}

// UnimplementedPlatformServer must be embedded to have forward compatible implementations.
type UnimplementedPlatformServer struct {
}

func (UnimplementedPlatformServer) CollectData(context.Context, *CollectDataRequest) (*CollectDataReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CollectData not implemented")
}
func (UnimplementedPlatformServer) mustEmbedUnimplementedPlatformServer() {}

// UnsafePlatformServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlatformServer will
// result in compilation errors.
type UnsafePlatformServer interface {
	mustEmbedUnimplementedPlatformServer()
}

func RegisterPlatformServer(s grpc.ServiceRegistrar, srv PlatformServer) {
	s.RegisterService(&Platform_ServiceDesc, srv)
}

func _Platform_CollectData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformServer).CollectData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Platform_CollectData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformServer).CollectData(ctx, req.(*CollectDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Platform_ServiceDesc is the grpc.ServiceDesc for Platform service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Platform_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.platform.Platform",
	HandlerType: (*PlatformServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CollectData",
			Handler:    _Platform_CollectData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "platform/platform.proto",
}
