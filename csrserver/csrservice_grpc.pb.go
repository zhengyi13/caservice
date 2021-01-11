// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package csrserver

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// CSRServiceClient is the client API for CSRService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CSRServiceClient interface {
	CreateCSR(ctx context.Context, in *CSRRequest, opts ...grpc.CallOption) (*CSRResponse, error)
	SignCSR(ctx context.Context, in *SignCSRRequest, opts ...grpc.CallOption) (*SignCSRResponse, error)
}

type cSRServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCSRServiceClient(cc grpc.ClientConnInterface) CSRServiceClient {
	return &cSRServiceClient{cc}
}

func (c *cSRServiceClient) CreateCSR(ctx context.Context, in *CSRRequest, opts ...grpc.CallOption) (*CSRResponse, error) {
	out := new(CSRResponse)
	err := c.cc.Invoke(ctx, "/csrserver.CSRService/CreateCSR", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cSRServiceClient) SignCSR(ctx context.Context, in *SignCSRRequest, opts ...grpc.CallOption) (*SignCSRResponse, error) {
	out := new(SignCSRResponse)
	err := c.cc.Invoke(ctx, "/csrserver.CSRService/SignCSR", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CSRServiceServer is the server API for CSRService service.
// All implementations must embed UnimplementedCSRServiceServer
// for forward compatibility
type CSRServiceServer interface {
	CreateCSR(context.Context, *CSRRequest) (*CSRResponse, error)
	SignCSR(context.Context, *SignCSRRequest) (*SignCSRResponse, error)
	mustEmbedUnimplementedCSRServiceServer()
}

// UnimplementedCSRServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCSRServiceServer struct {
}

func (UnimplementedCSRServiceServer) CreateCSR(context.Context, *CSRRequest) (*CSRResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCSR not implemented")
}
func (UnimplementedCSRServiceServer) SignCSR(context.Context, *SignCSRRequest) (*SignCSRResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignCSR not implemented")
}
func (UnimplementedCSRServiceServer) mustEmbedUnimplementedCSRServiceServer() {}

// UnsafeCSRServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CSRServiceServer will
// result in compilation errors.
type UnsafeCSRServiceServer interface {
	mustEmbedUnimplementedCSRServiceServer()
}

func RegisterCSRServiceServer(s grpc.ServiceRegistrar, srv CSRServiceServer) {
	s.RegisterService(&_CSRService_serviceDesc, srv)
}

func _CSRService_CreateCSR_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CSRRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CSRServiceServer).CreateCSR(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/csrserver.CSRService/CreateCSR",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CSRServiceServer).CreateCSR(ctx, req.(*CSRRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CSRService_SignCSR_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignCSRRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CSRServiceServer).SignCSR(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/csrserver.CSRService/SignCSR",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CSRServiceServer).SignCSR(ctx, req.(*SignCSRRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CSRService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "csrserver.CSRService",
	HandlerType: (*CSRServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCSR",
			Handler:    _CSRService_CreateCSR_Handler,
		},
		{
			MethodName: "SignCSR",
			Handler:    _CSRService_SignCSR_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "csrservice.proto",
}
