// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: catalog.proto

package catalog

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

// CatalogClient is the client API for Catalog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CatalogClient interface {
	CreateRestaurant(ctx context.Context, in *CreateRestaurantRequest, opts ...grpc.CallOption) (*RestaurantResponse, error)
	CreateMenuItem(ctx context.Context, in *CreateMenuItemRequest, opts ...grpc.CallOption) (*MenuItemResponse, error)
	FetchAllRestaurants(ctx context.Context, in *FetchAllRestaurantsRequest, opts ...grpc.CallOption) (*RestaurantsResponse, error)
	FetchRestaurantMenuItems(ctx context.Context, in *FetchMenuItemsRequest, opts ...grpc.CallOption) (*MenuItemsResponse, error)
}

type catalogClient struct {
	cc grpc.ClientConnInterface
}

func NewCatalogClient(cc grpc.ClientConnInterface) CatalogClient {
	return &catalogClient{cc}
}

func (c *catalogClient) CreateRestaurant(ctx context.Context, in *CreateRestaurantRequest, opts ...grpc.CallOption) (*RestaurantResponse, error) {
	out := new(RestaurantResponse)
	err := c.cc.Invoke(ctx, "/Catalog/CreateRestaurant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogClient) CreateMenuItem(ctx context.Context, in *CreateMenuItemRequest, opts ...grpc.CallOption) (*MenuItemResponse, error) {
	out := new(MenuItemResponse)
	err := c.cc.Invoke(ctx, "/Catalog/CreateMenuItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogClient) FetchAllRestaurants(ctx context.Context, in *FetchAllRestaurantsRequest, opts ...grpc.CallOption) (*RestaurantsResponse, error) {
	out := new(RestaurantsResponse)
	err := c.cc.Invoke(ctx, "/Catalog/FetchAllRestaurants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogClient) FetchRestaurantMenuItems(ctx context.Context, in *FetchMenuItemsRequest, opts ...grpc.CallOption) (*MenuItemsResponse, error) {
	out := new(MenuItemsResponse)
	err := c.cc.Invoke(ctx, "/Catalog/FetchRestaurantMenuItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatalogServer is the server API for Catalog service.
// All implementations must embed UnimplementedCatalogServer
// for forward compatibility
type CatalogServer interface {
	CreateRestaurant(context.Context, *CreateRestaurantRequest) (*RestaurantResponse, error)
	CreateMenuItem(context.Context, *CreateMenuItemRequest) (*MenuItemResponse, error)
	FetchAllRestaurants(context.Context, *FetchAllRestaurantsRequest) (*RestaurantsResponse, error)
	FetchRestaurantMenuItems(context.Context, *FetchMenuItemsRequest) (*MenuItemsResponse, error)
	mustEmbedUnimplementedCatalogServer()
}

// UnimplementedCatalogServer must be embedded to have forward compatible implementations.
type UnimplementedCatalogServer struct {
}

func (UnimplementedCatalogServer) CreateRestaurant(context.Context, *CreateRestaurantRequest) (*RestaurantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRestaurant not implemented")
}
func (UnimplementedCatalogServer) CreateMenuItem(context.Context, *CreateMenuItemRequest) (*MenuItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMenuItem not implemented")
}
func (UnimplementedCatalogServer) FetchAllRestaurants(context.Context, *FetchAllRestaurantsRequest) (*RestaurantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchAllRestaurants not implemented")
}
func (UnimplementedCatalogServer) FetchRestaurantMenuItems(context.Context, *FetchMenuItemsRequest) (*MenuItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchRestaurantMenuItems not implemented")
}
func (UnimplementedCatalogServer) mustEmbedUnimplementedCatalogServer() {}

// UnsafeCatalogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CatalogServer will
// result in compilation errors.
type UnsafeCatalogServer interface {
	mustEmbedUnimplementedCatalogServer()
}

func RegisterCatalogServer(s grpc.ServiceRegistrar, srv CatalogServer) {
	s.RegisterService(&Catalog_ServiceDesc, srv)
}

func _Catalog_CreateRestaurant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRestaurantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).CreateRestaurant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Catalog/CreateRestaurant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).CreateRestaurant(ctx, req.(*CreateRestaurantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catalog_CreateMenuItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMenuItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).CreateMenuItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Catalog/CreateMenuItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).CreateMenuItem(ctx, req.(*CreateMenuItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catalog_FetchAllRestaurants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchAllRestaurantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).FetchAllRestaurants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Catalog/FetchAllRestaurants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).FetchAllRestaurants(ctx, req.(*FetchAllRestaurantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Catalog_FetchRestaurantMenuItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchMenuItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServer).FetchRestaurantMenuItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Catalog/FetchRestaurantMenuItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServer).FetchRestaurantMenuItems(ctx, req.(*FetchMenuItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Catalog_ServiceDesc is the grpc.ServiceDesc for Catalog service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Catalog_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Catalog",
	HandlerType: (*CatalogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRestaurant",
			Handler:    _Catalog_CreateRestaurant_Handler,
		},
		{
			MethodName: "CreateMenuItem",
			Handler:    _Catalog_CreateMenuItem_Handler,
		},
		{
			MethodName: "FetchAllRestaurants",
			Handler:    _Catalog_FetchAllRestaurants_Handler,
		},
		{
			MethodName: "FetchRestaurantMenuItems",
			Handler:    _Catalog_FetchRestaurantMenuItems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "catalog.proto",
}
