// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: foods.proto

package orders_service

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

// FoodServiceClient is the client API for FoodService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FoodServiceClient interface {
	Create(ctx context.Context, in *CreateFood, opts ...grpc.CallOption) (*Food, error)
	GetById(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*Food, error)
	GetTitleById(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*Title, error)
	GetAll(ctx context.Context, in *FoodFilter, opts ...grpc.CallOption) (*Foods, error)
	Update(ctx context.Context, in *UpdateFood, opts ...grpc.CallOption) (*Food, error)
	Delete(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*Void, error)
}

type foodServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFoodServiceClient(cc grpc.ClientConnInterface) FoodServiceClient {
	return &foodServiceClient{cc}
}

func (c *foodServiceClient) Create(ctx context.Context, in *CreateFood, opts ...grpc.CallOption) (*Food, error) {
	out := new(Food)
	err := c.cc.Invoke(ctx, "/orders_service.FoodService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foodServiceClient) GetById(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*Food, error) {
	out := new(Food)
	err := c.cc.Invoke(ctx, "/orders_service.FoodService/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foodServiceClient) GetTitleById(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*Title, error) {
	out := new(Title)
	err := c.cc.Invoke(ctx, "/orders_service.FoodService/GetTitleById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foodServiceClient) GetAll(ctx context.Context, in *FoodFilter, opts ...grpc.CallOption) (*Foods, error) {
	out := new(Foods)
	err := c.cc.Invoke(ctx, "/orders_service.FoodService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foodServiceClient) Update(ctx context.Context, in *UpdateFood, opts ...grpc.CallOption) (*Food, error) {
	out := new(Food)
	err := c.cc.Invoke(ctx, "/orders_service.FoodService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foodServiceClient) Delete(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/orders_service.FoodService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FoodServiceServer is the server API for FoodService service.
// All implementations must embed UnimplementedFoodServiceServer
// for forward compatibility
type FoodServiceServer interface {
	Create(context.Context, *CreateFood) (*Food, error)
	GetById(context.Context, *PrimaryKey) (*Food, error)
	GetTitleById(context.Context, *PrimaryKey) (*Title, error)
	GetAll(context.Context, *FoodFilter) (*Foods, error)
	Update(context.Context, *UpdateFood) (*Food, error)
	Delete(context.Context, *PrimaryKey) (*Void, error)
	mustEmbedUnimplementedFoodServiceServer()
}

// UnimplementedFoodServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFoodServiceServer struct {
}

func (UnimplementedFoodServiceServer) Create(context.Context, *CreateFood) (*Food, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedFoodServiceServer) GetById(context.Context, *PrimaryKey) (*Food, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedFoodServiceServer) GetTitleById(context.Context, *PrimaryKey) (*Title, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTitleById not implemented")
}
func (UnimplementedFoodServiceServer) GetAll(context.Context, *FoodFilter) (*Foods, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedFoodServiceServer) Update(context.Context, *UpdateFood) (*Food, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedFoodServiceServer) Delete(context.Context, *PrimaryKey) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedFoodServiceServer) mustEmbedUnimplementedFoodServiceServer() {}

// UnsafeFoodServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FoodServiceServer will
// result in compilation errors.
type UnsafeFoodServiceServer interface {
	mustEmbedUnimplementedFoodServiceServer()
}

func RegisterFoodServiceServer(s grpc.ServiceRegistrar, srv FoodServiceServer) {
	s.RegisterService(&FoodService_ServiceDesc, srv)
}

func _FoodService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFood)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoodServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orders_service.FoodService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoodServiceServer).Create(ctx, req.(*CreateFood))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoodService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoodServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orders_service.FoodService/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoodServiceServer).GetById(ctx, req.(*PrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoodService_GetTitleById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoodServiceServer).GetTitleById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orders_service.FoodService/GetTitleById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoodServiceServer).GetTitleById(ctx, req.(*PrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoodService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FoodFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoodServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orders_service.FoodService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoodServiceServer).GetAll(ctx, req.(*FoodFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoodService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFood)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoodServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orders_service.FoodService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoodServiceServer).Update(ctx, req.(*UpdateFood))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoodService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoodServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orders_service.FoodService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoodServiceServer).Delete(ctx, req.(*PrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

// FoodService_ServiceDesc is the grpc.ServiceDesc for FoodService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FoodService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "orders_service.FoodService",
	HandlerType: (*FoodServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _FoodService_Create_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _FoodService_GetById_Handler,
		},
		{
			MethodName: "GetTitleById",
			Handler:    _FoodService_GetTitleById_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _FoodService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _FoodService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _FoodService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "foods.proto",
}
