// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: test.proto

package protos

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
	TestRpcService_AddUserData_FullMethodName = "/TestRpcService/addUserData"
	TestRpcService_GetUserData_FullMethodName = "/TestRpcService/getUserData"
)

// TestRpcServiceClient is the client API for TestRpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestRpcServiceClient interface {
	// Write data to database
	AddUserData(ctx context.Context, in *AddUserDataRequest, opts ...grpc.CallOption) (*AddUserDataResponse, error)
	// Return information based on uid and the specified field set (required)
	GetUserData(ctx context.Context, in *GetUserDataRequest, opts ...grpc.CallOption) (*GetUserDataResponse, error)
}

type testRpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTestRpcServiceClient(cc grpc.ClientConnInterface) TestRpcServiceClient {
	return &testRpcServiceClient{cc}
}

func (c *testRpcServiceClient) AddUserData(ctx context.Context, in *AddUserDataRequest, opts ...grpc.CallOption) (*AddUserDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddUserDataResponse)
	err := c.cc.Invoke(ctx, TestRpcService_AddUserData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testRpcServiceClient) GetUserData(ctx context.Context, in *GetUserDataRequest, opts ...grpc.CallOption) (*GetUserDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserDataResponse)
	err := c.cc.Invoke(ctx, TestRpcService_GetUserData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestRpcServiceServer is the server API for TestRpcService service.
// All implementations must embed UnimplementedTestRpcServiceServer
// for forward compatibility.
type TestRpcServiceServer interface {
	// Write data to database
	AddUserData(context.Context, *AddUserDataRequest) (*AddUserDataResponse, error)
	// Return information based on uid and the specified field set (required)
	GetUserData(context.Context, *GetUserDataRequest) (*GetUserDataResponse, error)
	mustEmbedUnimplementedTestRpcServiceServer()
}

// UnimplementedTestRpcServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTestRpcServiceServer struct{}

func (UnimplementedTestRpcServiceServer) AddUserData(context.Context, *AddUserDataRequest) (*AddUserDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserData not implemented")
}
func (UnimplementedTestRpcServiceServer) GetUserData(context.Context, *GetUserDataRequest) (*GetUserDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserData not implemented")
}
func (UnimplementedTestRpcServiceServer) mustEmbedUnimplementedTestRpcServiceServer() {}
func (UnimplementedTestRpcServiceServer) testEmbeddedByValue()                        {}

// UnsafeTestRpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestRpcServiceServer will
// result in compilation errors.
type UnsafeTestRpcServiceServer interface {
	mustEmbedUnimplementedTestRpcServiceServer()
}

func RegisterTestRpcServiceServer(s grpc.ServiceRegistrar, srv TestRpcServiceServer) {
	// If the following call pancis, it indicates UnimplementedTestRpcServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TestRpcService_ServiceDesc, srv)
}

func _TestRpcService_AddUserData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestRpcServiceServer).AddUserData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestRpcService_AddUserData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestRpcServiceServer).AddUserData(ctx, req.(*AddUserDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestRpcService_GetUserData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestRpcServiceServer).GetUserData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestRpcService_GetUserData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestRpcServiceServer).GetUserData(ctx, req.(*GetUserDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TestRpcService_ServiceDesc is the grpc.ServiceDesc for TestRpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestRpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TestRpcService",
	HandlerType: (*TestRpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "addUserData",
			Handler:    _TestRpcService_AddUserData_Handler,
		},
		{
			MethodName: "getUserData",
			Handler:    _TestRpcService_GetUserData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}
