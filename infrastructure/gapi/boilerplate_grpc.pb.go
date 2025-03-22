// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: boilerplate.proto

package gapi

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	BoilerplateService_CreateBoilerplate_FullMethodName = "/boilerplate.BoilerplateService/CreateBoilerplate"
	BoilerplateService_GetBoilerplate_FullMethodName    = "/boilerplate.BoilerplateService/GetBoilerplate"
)

// BoilerplateServiceClient is the client API for BoilerplateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BoilerplateServiceClient interface {
	CreateBoilerplate(ctx context.Context, in *CreateBoilerplateRequest, opts ...grpc.CallOption) (*CreateBoilerplateResponse, error)
	GetBoilerplate(ctx context.Context, in *GetBoilerplateRequest, opts ...grpc.CallOption) (*GetBoilerplateResponse, error)
}

type boilerplateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBoilerplateServiceClient(cc grpc.ClientConnInterface) BoilerplateServiceClient {
	return &boilerplateServiceClient{cc}
}

func (c *boilerplateServiceClient) CreateBoilerplate(ctx context.Context, in *CreateBoilerplateRequest, opts ...grpc.CallOption) (*CreateBoilerplateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateBoilerplateResponse)
	err := c.cc.Invoke(ctx, BoilerplateService_CreateBoilerplate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boilerplateServiceClient) GetBoilerplate(ctx context.Context, in *GetBoilerplateRequest, opts ...grpc.CallOption) (*GetBoilerplateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBoilerplateResponse)
	err := c.cc.Invoke(ctx, BoilerplateService_GetBoilerplate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BoilerplateServiceServer is the server API for BoilerplateService service.
// All implementations must embed UnimplementedBoilerplateServiceServer
// for forward compatibility.
type BoilerplateServiceServer interface {
	CreateBoilerplate(context.Context, *CreateBoilerplateRequest) (*CreateBoilerplateResponse, error)
	GetBoilerplate(context.Context, *GetBoilerplateRequest) (*GetBoilerplateResponse, error)
	mustEmbedUnimplementedBoilerplateServiceServer()
}

// UnimplementedBoilerplateServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBoilerplateServiceServer struct{}

func (UnimplementedBoilerplateServiceServer) CreateBoilerplate(context.Context, *CreateBoilerplateRequest) (*CreateBoilerplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBoilerplate not implemented")
}
func (UnimplementedBoilerplateServiceServer) GetBoilerplate(context.Context, *GetBoilerplateRequest) (*GetBoilerplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBoilerplate not implemented")
}
func (UnimplementedBoilerplateServiceServer) mustEmbedUnimplementedBoilerplateServiceServer() {}
func (UnimplementedBoilerplateServiceServer) testEmbeddedByValue()                            {}

// UnsafeBoilerplateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BoilerplateServiceServer will
// result in compilation errors.
type UnsafeBoilerplateServiceServer interface {
	mustEmbedUnimplementedBoilerplateServiceServer()
}

func RegisterBoilerplateServiceServer(s grpc.ServiceRegistrar, srv BoilerplateServiceServer) {
	// If the following call pancis, it indicates UnimplementedBoilerplateServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BoilerplateService_ServiceDesc, srv)
}

func _BoilerplateService_CreateBoilerplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBoilerplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoilerplateServiceServer).CreateBoilerplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BoilerplateService_CreateBoilerplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoilerplateServiceServer).CreateBoilerplate(ctx, req.(*CreateBoilerplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BoilerplateService_GetBoilerplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBoilerplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoilerplateServiceServer).GetBoilerplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BoilerplateService_GetBoilerplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoilerplateServiceServer).GetBoilerplate(ctx, req.(*GetBoilerplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BoilerplateService_ServiceDesc is the grpc.ServiceDesc for BoilerplateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BoilerplateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "boilerplate.BoilerplateService",
	HandlerType: (*BoilerplateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBoilerplate",
			Handler:    _BoilerplateService_CreateBoilerplate_Handler,
		},
		{
			MethodName: "GetBoilerplate",
			Handler:    _BoilerplateService_GetBoilerplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "boilerplate.proto",
}
