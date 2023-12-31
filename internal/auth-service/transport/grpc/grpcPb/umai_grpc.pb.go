// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: transport/grpc/proto/umai.proto

package grpcPb

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

// IdentityServiceClient is the client API for IdentityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IdentityServiceClient interface {
	GetResIdentity(ctx context.Context, in *IdentityResReq, opts ...grpc.CallOption) (*IdentityResResp, error)
	GetUserIdentity(ctx context.Context, in *IdentityReq, opts ...grpc.CallOption) (*IdentityResp, error)
}

type identityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIdentityServiceClient(cc grpc.ClientConnInterface) IdentityServiceClient {
	return &identityServiceClient{cc}
}

func (c *identityServiceClient) GetResIdentity(ctx context.Context, in *IdentityResReq, opts ...grpc.CallOption) (*IdentityResResp, error) {
	out := new(IdentityResResp)
	err := c.cc.Invoke(ctx, "/grpc.IdentityService/GetResIdentity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityServiceClient) GetUserIdentity(ctx context.Context, in *IdentityReq, opts ...grpc.CallOption) (*IdentityResp, error) {
	out := new(IdentityResp)
	err := c.cc.Invoke(ctx, "/grpc.IdentityService/GetUserIdentity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentityServiceServer is the server API for IdentityService service.
// All implementations must embed UnimplementedIdentityServiceServer
// for forward compatibility
type IdentityServiceServer interface {
	GetResIdentity(context.Context, *IdentityResReq) (*IdentityResResp, error)
	GetUserIdentity(context.Context, *IdentityReq) (*IdentityResp, error)
	mustEmbedUnimplementedIdentityServiceServer()
}

// UnimplementedIdentityServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIdentityServiceServer struct {
}

func (UnimplementedIdentityServiceServer) GetResIdentity(context.Context, *IdentityResReq) (*IdentityResResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResIdentity not implemented")
}
func (UnimplementedIdentityServiceServer) GetUserIdentity(context.Context, *IdentityReq) (*IdentityResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserIdentity not implemented")
}
func (UnimplementedIdentityServiceServer) mustEmbedUnimplementedIdentityServiceServer() {}

// UnsafeIdentityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IdentityServiceServer will
// result in compilation errors.
type UnsafeIdentityServiceServer interface {
	mustEmbedUnimplementedIdentityServiceServer()
}

func RegisterIdentityServiceServer(s grpc.ServiceRegistrar, srv IdentityServiceServer) {
	s.RegisterService(&IdentityService_ServiceDesc, srv)
}

func _IdentityService_GetResIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdentityResReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServiceServer).GetResIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.IdentityService/GetResIdentity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServiceServer).GetResIdentity(ctx, req.(*IdentityResReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityService_GetUserIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdentityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServiceServer).GetUserIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.IdentityService/GetUserIdentity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServiceServer).GetUserIdentity(ctx, req.(*IdentityReq))
	}
	return interceptor(ctx, in, info, handler)
}

// IdentityService_ServiceDesc is the grpc.ServiceDesc for IdentityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IdentityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.IdentityService",
	HandlerType: (*IdentityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetResIdentity",
			Handler:    _IdentityService_GetResIdentity_Handler,
		},
		{
			MethodName: "GetUserIdentity",
			Handler:    _IdentityService_GetUserIdentity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transport/grpc/proto/umai.proto",
}
