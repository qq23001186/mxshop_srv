// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: order.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OrderClient is the client API for Order service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderClient interface {
	//购物车
	CartItemList(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*CartItemListResponse, error)
	CreateCartItem(ctx context.Context, in *CartItemRequest, opts ...grpc.CallOption) (*ShopCartInfoResponse, error)
	UpdateCartItem(ctx context.Context, in *CartItemRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteCartItem(ctx context.Context, in *CartItemRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	//订单
	CreateOrder(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderInfoResponse, error)
	OrderList(ctx context.Context, in *OrderFilterRequest, opts ...grpc.CallOption) (*OrderListResponse, error)
	OrderDetail(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderInfoDetailResponse, error)
	UpdateOrderStatus(ctx context.Context, in *OrderStatus, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type orderClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderClient(cc grpc.ClientConnInterface) OrderClient {
	return &orderClient{cc}
}

func (c *orderClient) CartItemList(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*CartItemListResponse, error) {
	out := new(CartItemListResponse)
	err := c.cc.Invoke(ctx, "/Order/CartItemList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) CreateCartItem(ctx context.Context, in *CartItemRequest, opts ...grpc.CallOption) (*ShopCartInfoResponse, error) {
	out := new(ShopCartInfoResponse)
	err := c.cc.Invoke(ctx, "/Order/CreateCartItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) UpdateCartItem(ctx context.Context, in *CartItemRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Order/UpdateCartItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) DeleteCartItem(ctx context.Context, in *CartItemRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Order/DeleteCartItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) CreateOrder(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderInfoResponse, error) {
	out := new(OrderInfoResponse)
	err := c.cc.Invoke(ctx, "/Order/CreateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) OrderList(ctx context.Context, in *OrderFilterRequest, opts ...grpc.CallOption) (*OrderListResponse, error) {
	out := new(OrderListResponse)
	err := c.cc.Invoke(ctx, "/Order/OrderList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) OrderDetail(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderInfoDetailResponse, error) {
	out := new(OrderInfoDetailResponse)
	err := c.cc.Invoke(ctx, "/Order/OrderDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) UpdateOrderStatus(ctx context.Context, in *OrderStatus, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Order/UpdateOrderStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServer is the server API for Order service.
// All implementations must embed UnimplementedOrderServer
// for forward compatibility
type OrderServer interface {
	//购物车
	CartItemList(context.Context, *UserInfo) (*CartItemListResponse, error)
	CreateCartItem(context.Context, *CartItemRequest) (*ShopCartInfoResponse, error)
	UpdateCartItem(context.Context, *CartItemRequest) (*emptypb.Empty, error)
	DeleteCartItem(context.Context, *CartItemRequest) (*emptypb.Empty, error)
	//订单
	CreateOrder(context.Context, *OrderRequest) (*OrderInfoResponse, error)
	OrderList(context.Context, *OrderFilterRequest) (*OrderListResponse, error)
	OrderDetail(context.Context, *OrderRequest) (*OrderInfoDetailResponse, error)
	UpdateOrderStatus(context.Context, *OrderStatus) (*emptypb.Empty, error)
	mustEmbedUnimplementedOrderServer()
}

// UnimplementedOrderServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServer struct {
}

func (UnimplementedOrderServer) CartItemList(context.Context, *UserInfo) (*CartItemListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CartItemList not implemented")
}
func (UnimplementedOrderServer) CreateCartItem(context.Context, *CartItemRequest) (*ShopCartInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCartItem not implemented")
}
func (UnimplementedOrderServer) UpdateCartItem(context.Context, *CartItemRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCartItem not implemented")
}
func (UnimplementedOrderServer) DeleteCartItem(context.Context, *CartItemRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCartItem not implemented")
}
func (UnimplementedOrderServer) CreateOrder(context.Context, *OrderRequest) (*OrderInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedOrderServer) OrderList(context.Context, *OrderFilterRequest) (*OrderListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderList not implemented")
}
func (UnimplementedOrderServer) OrderDetail(context.Context, *OrderRequest) (*OrderInfoDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderDetail not implemented")
}
func (UnimplementedOrderServer) UpdateOrderStatus(context.Context, *OrderStatus) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrderStatus not implemented")
}
func (UnimplementedOrderServer) mustEmbedUnimplementedOrderServer() {}

// UnsafeOrderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServer will
// result in compilation errors.
type UnsafeOrderServer interface {
	mustEmbedUnimplementedOrderServer()
}

func RegisterOrderServer(s grpc.ServiceRegistrar, srv OrderServer) {
	s.RegisterService(&Order_ServiceDesc, srv)
}

func _Order_CartItemList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).CartItemList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Order/CartItemList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).CartItemList(ctx, req.(*UserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_CreateCartItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).CreateCartItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Order/CreateCartItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).CreateCartItem(ctx, req.(*CartItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_UpdateCartItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).UpdateCartItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Order/UpdateCartItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).UpdateCartItem(ctx, req.(*CartItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_DeleteCartItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).DeleteCartItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Order/DeleteCartItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).DeleteCartItem(ctx, req.(*CartItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Order/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).CreateOrder(ctx, req.(*OrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_OrderList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).OrderList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Order/OrderList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).OrderList(ctx, req.(*OrderFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_OrderDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).OrderDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Order/OrderDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).OrderDetail(ctx, req.(*OrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_UpdateOrderStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).UpdateOrderStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Order/UpdateOrderStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).UpdateOrderStatus(ctx, req.(*OrderStatus))
	}
	return interceptor(ctx, in, info, handler)
}

// Order_ServiceDesc is the grpc.ServiceDesc for Order service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Order_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Order",
	HandlerType: (*OrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CartItemList",
			Handler:    _Order_CartItemList_Handler,
		},
		{
			MethodName: "CreateCartItem",
			Handler:    _Order_CreateCartItem_Handler,
		},
		{
			MethodName: "UpdateCartItem",
			Handler:    _Order_UpdateCartItem_Handler,
		},
		{
			MethodName: "DeleteCartItem",
			Handler:    _Order_DeleteCartItem_Handler,
		},
		{
			MethodName: "CreateOrder",
			Handler:    _Order_CreateOrder_Handler,
		},
		{
			MethodName: "OrderList",
			Handler:    _Order_OrderList_Handler,
		},
		{
			MethodName: "OrderDetail",
			Handler:    _Order_OrderDetail_Handler,
		},
		{
			MethodName: "UpdateOrderStatus",
			Handler:    _Order_UpdateOrderStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order.proto",
}
