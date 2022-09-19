// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: student.proto

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

// StudentServiceClient is the client API for StudentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StudentServiceClient interface {
	FindStudentByID(ctx context.Context, in *StudentByIDRequest, opts ...grpc.CallOption) (*StudentResponse, error)
	FindAllStudents(ctx context.Context, in *AllStudentsRequest, opts ...grpc.CallOption) (*StudentsResponse, error)
}

type studentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStudentServiceClient(cc grpc.ClientConnInterface) StudentServiceClient {
	return &studentServiceClient{cc}
}

func (c *studentServiceClient) FindStudentByID(ctx context.Context, in *StudentByIDRequest, opts ...grpc.CallOption) (*StudentResponse, error) {
	out := new(StudentResponse)
	err := c.cc.Invoke(ctx, "/proto.StudentService/FindStudentByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) FindAllStudents(ctx context.Context, in *AllStudentsRequest, opts ...grpc.CallOption) (*StudentsResponse, error) {
	out := new(StudentsResponse)
	err := c.cc.Invoke(ctx, "/proto.StudentService/FindAllStudents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StudentServiceServer is the server API for StudentService service.
// All implementations must embed UnimplementedStudentServiceServer
// for forward compatibility
type StudentServiceServer interface {
	FindStudentByID(context.Context, *StudentByIDRequest) (*StudentResponse, error)
	FindAllStudents(context.Context, *AllStudentsRequest) (*StudentsResponse, error)
	mustEmbedUnimplementedStudentServiceServer()
}

// UnimplementedStudentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStudentServiceServer struct {
}

func (UnimplementedStudentServiceServer) FindStudentByID(context.Context, *StudentByIDRequest) (*StudentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindStudentByID not implemented")
}
func (UnimplementedStudentServiceServer) FindAllStudents(context.Context, *AllStudentsRequest) (*StudentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllStudents not implemented")
}
func (UnimplementedStudentServiceServer) mustEmbedUnimplementedStudentServiceServer() {}

// UnsafeStudentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StudentServiceServer will
// result in compilation errors.
type UnsafeStudentServiceServer interface {
	mustEmbedUnimplementedStudentServiceServer()
}

func RegisterStudentServiceServer(s grpc.ServiceRegistrar, srv StudentServiceServer) {
	s.RegisterService(&StudentService_ServiceDesc, srv)
}

func _StudentService_FindStudentByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StudentByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).FindStudentByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.StudentService/FindStudentByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).FindStudentByID(ctx, req.(*StudentByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_FindAllStudents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllStudentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).FindAllStudents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.StudentService/FindAllStudents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).FindAllStudents(ctx, req.(*AllStudentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StudentService_ServiceDesc is the grpc.ServiceDesc for StudentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StudentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.StudentService",
	HandlerType: (*StudentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindStudentByID",
			Handler:    _StudentService_FindStudentByID_Handler,
		},
		{
			MethodName: "FindAllStudents",
			Handler:    _StudentService_FindAllStudents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "student.proto",
}