// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: hostname.proto

package hostname

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	HostnameService_SetHostName_FullMethodName = "/hostname.HostnameService/SetHostName"
)

// HostnameServiceClient is the client API for HostnameService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HostnameServiceClient interface {
	SetHostName(ctx context.Context, in *SetHostnameRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type hostnameServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHostnameServiceClient(cc grpc.ClientConnInterface) HostnameServiceClient {
	return &hostnameServiceClient{cc}
}

func (c *hostnameServiceClient) SetHostName(ctx context.Context, in *SetHostnameRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, HostnameService_SetHostName_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HostnameServiceServer is the server API for HostnameService service.
// All implementations must embed UnimplementedHostnameServiceServer
// for forward compatibility
type HostnameServiceServer interface {
	SetHostName(context.Context, *SetHostnameRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedHostnameServiceServer()
}

// UnimplementedHostnameServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHostnameServiceServer struct {
}

func (UnimplementedHostnameServiceServer) SetHostName(context.Context, *SetHostnameRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetHostName not implemented")
}
func (UnimplementedHostnameServiceServer) mustEmbedUnimplementedHostnameServiceServer() {}

// UnsafeHostnameServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HostnameServiceServer will
// result in compilation errors.
type UnsafeHostnameServiceServer interface {
	mustEmbedUnimplementedHostnameServiceServer()
}

func RegisterHostnameServiceServer(s grpc.ServiceRegistrar, srv HostnameServiceServer) {
	s.RegisterService(&HostnameService_ServiceDesc, srv)
}

func _HostnameService_SetHostName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetHostnameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostnameServiceServer).SetHostName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HostnameService_SetHostName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostnameServiceServer).SetHostName(ctx, req.(*SetHostnameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HostnameService_ServiceDesc is the grpc.ServiceDesc for HostnameService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HostnameService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hostname.HostnameService",
	HandlerType: (*HostnameServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetHostName",
			Handler:    _HostnameService_SetHostName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hostname.proto",
}