// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: categories.proto

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

// CategoriesServiceClient is the client API for CategoriesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CategoriesServiceClient interface {
	Create(ctx context.Context, in *CreateCategory, opts ...grpc.CallOption) (*Category, error)
	GetById(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*Category, error)
	GetAll(ctx context.Context, in *CategoryFilter, opts ...grpc.CallOption) (*Categories, error)
	Update(ctx context.Context, in *Category, opts ...grpc.CallOption) (*Category, error)
	Delete(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*Void, error)
	CheckCategoryExist(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*Void, error)
}

type categoriesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCategoriesServiceClient(cc grpc.ClientConnInterface) CategoriesServiceClient {
	return &categoriesServiceClient{cc}
}

func (c *categoriesServiceClient) Create(ctx context.Context, in *CreateCategory, opts ...grpc.CallOption) (*Category, error) {
	out := new(Category)
	err := c.cc.Invoke(ctx, "/orders_service.categories_service/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoriesServiceClient) GetById(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*Category, error) {
	out := new(Category)
	err := c.cc.Invoke(ctx, "/orders_service.categories_service/getById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoriesServiceClient) GetAll(ctx context.Context, in *CategoryFilter, opts ...grpc.CallOption) (*Categories, error) {
	out := new(Categories)
	err := c.cc.Invoke(ctx, "/orders_service.categories_service/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoriesServiceClient) Update(ctx context.Context, in *Category, opts ...grpc.CallOption) (*Category, error) {
	out := new(Category)
	err := c.cc.Invoke(ctx, "/orders_service.categories_service/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoriesServiceClient) Delete(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/orders_service.categories_service/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoriesServiceClient) CheckCategoryExist(ctx context.Context, in *PrimaryKey, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/orders_service.categories_service/CheckCategoryExist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CategoriesServiceServer is the server API for CategoriesService service.
// All implementations must embed UnimplementedCategoriesServiceServer
// for forward compatibility
type CategoriesServiceServer interface {
	Create(context.Context, *CreateCategory) (*Category, error)
	GetById(context.Context, *PrimaryKey) (*Category, error)
	GetAll(context.Context, *CategoryFilter) (*Categories, error)
	Update(context.Context, *Category) (*Category, error)
	Delete(context.Context, *PrimaryKey) (*Void, error)
	CheckCategoryExist(context.Context, *PrimaryKey) (*Void, error)
	mustEmbedUnimplementedCategoriesServiceServer()
}

// UnimplementedCategoriesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCategoriesServiceServer struct {
}

func (UnimplementedCategoriesServiceServer) Create(context.Context, *CreateCategory) (*Category, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCategoriesServiceServer) GetById(context.Context, *PrimaryKey) (*Category, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedCategoriesServiceServer) GetAll(context.Context, *CategoryFilter) (*Categories, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedCategoriesServiceServer) Update(context.Context, *Category) (*Category, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCategoriesServiceServer) Delete(context.Context, *PrimaryKey) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCategoriesServiceServer) CheckCategoryExist(context.Context, *PrimaryKey) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckCategoryExist not implemented")
}
func (UnimplementedCategoriesServiceServer) mustEmbedUnimplementedCategoriesServiceServer() {}

// UnsafeCategoriesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CategoriesServiceServer will
// result in compilation errors.
type UnsafeCategoriesServiceServer interface {
	mustEmbedUnimplementedCategoriesServiceServer()
}

func RegisterCategoriesServiceServer(s grpc.ServiceRegistrar, srv CategoriesServiceServer) {
	s.RegisterService(&CategoriesService_ServiceDesc, srv)
}

func _CategoriesService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCategory)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoriesServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orders_service.categories_service/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoriesServiceServer).Create(ctx, req.(*CreateCategory))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoriesService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoriesServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orders_service.categories_service/getById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoriesServiceServer).GetById(ctx, req.(*PrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoriesService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoriesServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orders_service.categories_service/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoriesServiceServer).GetAll(ctx, req.(*CategoryFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoriesService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Category)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoriesServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orders_service.categories_service/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoriesServiceServer).Update(ctx, req.(*Category))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoriesService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoriesServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orders_service.categories_service/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoriesServiceServer).Delete(ctx, req.(*PrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoriesService_CheckCategoryExist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoriesServiceServer).CheckCategoryExist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orders_service.categories_service/CheckCategoryExist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoriesServiceServer).CheckCategoryExist(ctx, req.(*PrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

// CategoriesService_ServiceDesc is the grpc.ServiceDesc for CategoriesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CategoriesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "orders_service.categories_service",
	HandlerType: (*CategoriesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _CategoriesService_Create_Handler,
		},
		{
			MethodName: "getById",
			Handler:    _CategoriesService_GetById_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _CategoriesService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CategoriesService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CategoriesService_Delete_Handler,
		},
		{
			MethodName: "CheckCategoryExist",
			Handler:    _CategoriesService_CheckCategoryExist_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "categories.proto",
}
