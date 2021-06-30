// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package stockupdate

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

// StockUpdateServiceClient is the client API for StockUpdateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StockUpdateServiceClient interface {
	GetStocks(ctx context.Context, in *GetStoreStock, opts ...grpc.CallOption) (*Stocks, error)
	AggregateStock(ctx context.Context, in *AggregateStoreStock, opts ...grpc.CallOption) (*Stocks, error)
}

type stockUpdateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStockUpdateServiceClient(cc grpc.ClientConnInterface) StockUpdateServiceClient {
	return &stockUpdateServiceClient{cc}
}

func (c *stockUpdateServiceClient) GetStocks(ctx context.Context, in *GetStoreStock, opts ...grpc.CallOption) (*Stocks, error) {
	out := new(Stocks)
	err := c.cc.Invoke(ctx, "/stockupdate.StockUpdateService/GetStocks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockUpdateServiceClient) AggregateStock(ctx context.Context, in *AggregateStoreStock, opts ...grpc.CallOption) (*Stocks, error) {
	out := new(Stocks)
	err := c.cc.Invoke(ctx, "/stockupdate.StockUpdateService/AggregateStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StockUpdateServiceServer is the server API for StockUpdateService service.
// All implementations must embed UnimplementedStockUpdateServiceServer
// for forward compatibility
type StockUpdateServiceServer interface {
	GetStocks(context.Context, *GetStoreStock) (*Stocks, error)
	AggregateStock(context.Context, *AggregateStoreStock) (*Stocks, error)
	mustEmbedUnimplementedStockUpdateServiceServer()
}

// UnimplementedStockUpdateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStockUpdateServiceServer struct {
}

func (UnimplementedStockUpdateServiceServer) GetStocks(context.Context, *GetStoreStock) (*Stocks, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStocks not implemented")
}
func (UnimplementedStockUpdateServiceServer) AggregateStock(context.Context, *AggregateStoreStock) (*Stocks, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AggregateStock not implemented")
}
func (UnimplementedStockUpdateServiceServer) mustEmbedUnimplementedStockUpdateServiceServer() {}

// UnsafeStockUpdateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StockUpdateServiceServer will
// result in compilation errors.
type UnsafeStockUpdateServiceServer interface {
	mustEmbedUnimplementedStockUpdateServiceServer()
}

func RegisterStockUpdateServiceServer(s grpc.ServiceRegistrar, srv StockUpdateServiceServer) {
	s.RegisterService(&StockUpdateService_ServiceDesc, srv)
}

func _StockUpdateService_GetStocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStoreStock)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockUpdateServiceServer).GetStocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stockupdate.StockUpdateService/GetStocks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockUpdateServiceServer).GetStocks(ctx, req.(*GetStoreStock))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockUpdateService_AggregateStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AggregateStoreStock)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockUpdateServiceServer).AggregateStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stockupdate.StockUpdateService/AggregateStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockUpdateServiceServer).AggregateStock(ctx, req.(*AggregateStoreStock))
	}
	return interceptor(ctx, in, info, handler)
}

// StockUpdateService_ServiceDesc is the grpc.ServiceDesc for StockUpdateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StockUpdateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stockupdate.StockUpdateService",
	HandlerType: (*StockUpdateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStocks",
			Handler:    _StockUpdateService_GetStocks_Handler,
		},
		{
			MethodName: "AggregateStock",
			Handler:    _StockUpdateService_AggregateStock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stockupdate.proto",
}