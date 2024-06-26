// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: proto/student.proto

package student

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
	CRUD_CreateStudent_FullMethodName = "/CRUD/CreateStudent"
	CRUD_ReadStudent_FullMethodName   = "/CRUD/ReadStudent"
	CRUD_UpdateStudent_FullMethodName = "/CRUD/UpdateStudent"
	CRUD_DeleteStudent_FullMethodName = "/CRUD/DeleteStudent"
)

// CRUDClient is the client API for CRUD service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CRUDClient interface {
	CreateStudent(ctx context.Context, in *Student, opts ...grpc.CallOption) (*ID, error)
	ReadStudent(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Student, error)
	UpdateStudent(ctx context.Context, in *Student, opts ...grpc.CallOption) (*ID, error)
	DeleteStudent(ctx context.Context, in *ID, opts ...grpc.CallOption) (*ID, error)
}

type cRUDClient struct {
	cc grpc.ClientConnInterface
}

func NewCRUDClient(cc grpc.ClientConnInterface) CRUDClient {
	return &cRUDClient{cc}
}

func (c *cRUDClient) CreateStudent(ctx context.Context, in *Student, opts ...grpc.CallOption) (*ID, error) {
	out := new(ID)
	err := c.cc.Invoke(ctx, CRUD_CreateStudent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDClient) ReadStudent(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Student, error) {
	out := new(Student)
	err := c.cc.Invoke(ctx, CRUD_ReadStudent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDClient) UpdateStudent(ctx context.Context, in *Student, opts ...grpc.CallOption) (*ID, error) {
	out := new(ID)
	err := c.cc.Invoke(ctx, CRUD_UpdateStudent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDClient) DeleteStudent(ctx context.Context, in *ID, opts ...grpc.CallOption) (*ID, error) {
	out := new(ID)
	err := c.cc.Invoke(ctx, CRUD_DeleteStudent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CRUDServer is the server API for CRUD service.
// All implementations must embed UnimplementedCRUDServer
// for forward compatibility
type CRUDServer interface {
	CreateStudent(context.Context, *Student) (*ID, error)
	ReadStudent(context.Context, *ID) (*Student, error)
	UpdateStudent(context.Context, *Student) (*ID, error)
	DeleteStudent(context.Context, *ID) (*ID, error)
	mustEmbedUnimplementedCRUDServer()
}

// UnimplementedCRUDServer must be embedded to have forward compatible implementations.
type UnimplementedCRUDServer struct {
}

func (UnimplementedCRUDServer) CreateStudent(context.Context, *Student) (*ID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStudent not implemented")
}
func (UnimplementedCRUDServer) ReadStudent(context.Context, *ID) (*Student, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadStudent not implemented")
}
func (UnimplementedCRUDServer) UpdateStudent(context.Context, *Student) (*ID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStudent not implemented")
}
func (UnimplementedCRUDServer) DeleteStudent(context.Context, *ID) (*ID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStudent not implemented")
}
func (UnimplementedCRUDServer) mustEmbedUnimplementedCRUDServer() {}

// UnsafeCRUDServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CRUDServer will
// result in compilation errors.
type UnsafeCRUDServer interface {
	mustEmbedUnimplementedCRUDServer()
}

func RegisterCRUDServer(s grpc.ServiceRegistrar, srv CRUDServer) {
	s.RegisterService(&CRUD_ServiceDesc, srv)
}

func _CRUD_CreateStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Student)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).CreateStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CRUD_CreateStudent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).CreateStudent(ctx, req.(*Student))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUD_ReadStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).ReadStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CRUD_ReadStudent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).ReadStudent(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUD_UpdateStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Student)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).UpdateStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CRUD_UpdateStudent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).UpdateStudent(ctx, req.(*Student))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUD_DeleteStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).DeleteStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CRUD_DeleteStudent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).DeleteStudent(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

// CRUD_ServiceDesc is the grpc.ServiceDesc for CRUD service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CRUD_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CRUD",
	HandlerType: (*CRUDServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStudent",
			Handler:    _CRUD_CreateStudent_Handler,
		},
		{
			MethodName: "ReadStudent",
			Handler:    _CRUD_ReadStudent_Handler,
		},
		{
			MethodName: "UpdateStudent",
			Handler:    _CRUD_UpdateStudent_Handler,
		},
		{
			MethodName: "DeleteStudent",
			Handler:    _CRUD_DeleteStudent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/student.proto",
}
