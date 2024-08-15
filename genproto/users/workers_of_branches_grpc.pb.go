// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: workers_of_branches.proto

package users

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

// WorkersOfBranchesServiceClient is the client API for WorkersOfBranchesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkersOfBranchesServiceClient interface {
	Create(ctx context.Context, in *CreateWorker, opts ...grpc.CallOption) (*WorkerId, error)
	GetByUUID(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*Worker, error)
	GetByWorkerId(ctx context.Context, in *WorkerId, opts ...grpc.CallOption) (*Worker, error)
	GetAll(ctx context.Context, in *WorkerFilter, opts ...grpc.CallOption) (*Workers, error)
	Update(ctx context.Context, in *Worker, opts ...grpc.CallOption) (*Worker, error)
	Delete(ctx context.Context, in *WorkerId, opts ...grpc.CallOption) (*Void, error)
	CheckWorkerExists(ctx context.Context, in *WorkerId, opts ...grpc.CallOption) (*Void, error)
}

type workersOfBranchesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkersOfBranchesServiceClient(cc grpc.ClientConnInterface) WorkersOfBranchesServiceClient {
	return &workersOfBranchesServiceClient{cc}
}

func (c *workersOfBranchesServiceClient) Create(ctx context.Context, in *CreateWorker, opts ...grpc.CallOption) (*WorkerId, error) {
	out := new(WorkerId)
	err := c.cc.Invoke(ctx, "/users.WorkersOfBranchesService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workersOfBranchesServiceClient) GetByUUID(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*Worker, error) {
	out := new(Worker)
	err := c.cc.Invoke(ctx, "/users.WorkersOfBranchesService/GetByUUID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workersOfBranchesServiceClient) GetByWorkerId(ctx context.Context, in *WorkerId, opts ...grpc.CallOption) (*Worker, error) {
	out := new(Worker)
	err := c.cc.Invoke(ctx, "/users.WorkersOfBranchesService/GetByWorkerId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workersOfBranchesServiceClient) GetAll(ctx context.Context, in *WorkerFilter, opts ...grpc.CallOption) (*Workers, error) {
	out := new(Workers)
	err := c.cc.Invoke(ctx, "/users.WorkersOfBranchesService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workersOfBranchesServiceClient) Update(ctx context.Context, in *Worker, opts ...grpc.CallOption) (*Worker, error) {
	out := new(Worker)
	err := c.cc.Invoke(ctx, "/users.WorkersOfBranchesService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workersOfBranchesServiceClient) Delete(ctx context.Context, in *WorkerId, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/users.WorkersOfBranchesService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workersOfBranchesServiceClient) CheckWorkerExists(ctx context.Context, in *WorkerId, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/users.WorkersOfBranchesService/CheckWorkerExists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkersOfBranchesServiceServer is the server API for WorkersOfBranchesService service.
// All implementations must embed UnimplementedWorkersOfBranchesServiceServer
// for forward compatibility
type WorkersOfBranchesServiceServer interface {
	Create(context.Context, *CreateWorker) (*WorkerId, error)
	GetByUUID(context.Context, *PrimaryKey) (*Worker, error)
	GetByWorkerId(context.Context, *WorkerId) (*Worker, error)
	GetAll(context.Context, *WorkerFilter) (*Workers, error)
	Update(context.Context, *Worker) (*Worker, error)
	Delete(context.Context, *WorkerId) (*Void, error)
	CheckWorkerExists(context.Context, *WorkerId) (*Void, error)
	mustEmbedUnimplementedWorkersOfBranchesServiceServer()
}

// UnimplementedWorkersOfBranchesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWorkersOfBranchesServiceServer struct {
}

func (UnimplementedWorkersOfBranchesServiceServer) Create(context.Context, *CreateWorker) (*WorkerId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedWorkersOfBranchesServiceServer) GetByUUID(context.Context, *PrimaryKey) (*Worker, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByUUID not implemented")
}
func (UnimplementedWorkersOfBranchesServiceServer) GetByWorkerId(context.Context, *WorkerId) (*Worker, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByWorkerId not implemented")
}
func (UnimplementedWorkersOfBranchesServiceServer) GetAll(context.Context, *WorkerFilter) (*Workers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedWorkersOfBranchesServiceServer) Update(context.Context, *Worker) (*Worker, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedWorkersOfBranchesServiceServer) Delete(context.Context, *WorkerId) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedWorkersOfBranchesServiceServer) CheckWorkerExists(context.Context, *WorkerId) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckWorkerExists not implemented")
}
func (UnimplementedWorkersOfBranchesServiceServer) mustEmbedUnimplementedWorkersOfBranchesServiceServer() {
}

// UnsafeWorkersOfBranchesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkersOfBranchesServiceServer will
// result in compilation errors.
type UnsafeWorkersOfBranchesServiceServer interface {
	mustEmbedUnimplementedWorkersOfBranchesServiceServer()
}

func RegisterWorkersOfBranchesServiceServer(s grpc.ServiceRegistrar, srv WorkersOfBranchesServiceServer) {
	s.RegisterService(&WorkersOfBranchesService_ServiceDesc, srv)
}

func _WorkersOfBranchesService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWorker)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkersOfBranchesServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.WorkersOfBranchesService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkersOfBranchesServiceServer).Create(ctx, req.(*CreateWorker))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkersOfBranchesService_GetByUUID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkersOfBranchesServiceServer).GetByUUID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.WorkersOfBranchesService/GetByUUID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkersOfBranchesServiceServer).GetByUUID(ctx, req.(*PrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkersOfBranchesService_GetByWorkerId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkerId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkersOfBranchesServiceServer).GetByWorkerId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.WorkersOfBranchesService/GetByWorkerId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkersOfBranchesServiceServer).GetByWorkerId(ctx, req.(*WorkerId))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkersOfBranchesService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkerFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkersOfBranchesServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.WorkersOfBranchesService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkersOfBranchesServiceServer).GetAll(ctx, req.(*WorkerFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkersOfBranchesService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Worker)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkersOfBranchesServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.WorkersOfBranchesService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkersOfBranchesServiceServer).Update(ctx, req.(*Worker))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkersOfBranchesService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkerId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkersOfBranchesServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.WorkersOfBranchesService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkersOfBranchesServiceServer).Delete(ctx, req.(*WorkerId))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkersOfBranchesService_CheckWorkerExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkerId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkersOfBranchesServiceServer).CheckWorkerExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.WorkersOfBranchesService/CheckWorkerExists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkersOfBranchesServiceServer).CheckWorkerExists(ctx, req.(*WorkerId))
	}
	return interceptor(ctx, in, info, handler)
}

// WorkersOfBranchesService_ServiceDesc is the grpc.ServiceDesc for WorkersOfBranchesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WorkersOfBranchesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "users.WorkersOfBranchesService",
	HandlerType: (*WorkersOfBranchesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _WorkersOfBranchesService_Create_Handler,
		},
		{
			MethodName: "GetByUUID",
			Handler:    _WorkersOfBranchesService_GetByUUID_Handler,
		},
		{
			MethodName: "GetByWorkerId",
			Handler:    _WorkersOfBranchesService_GetByWorkerId_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _WorkersOfBranchesService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _WorkersOfBranchesService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _WorkersOfBranchesService_Delete_Handler,
		},
		{
			MethodName: "CheckWorkerExists",
			Handler:    _WorkersOfBranchesService_CheckWorkerExists_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "workers_of_branches.proto",
}
