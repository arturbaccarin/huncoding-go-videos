// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: stock.proto

package pb

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
	StockService_StreamStockPrices_FullMethodName = "/pb.StockService/StreamStockPrices"
)

// StockServiceClient is the client API for StockService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StockServiceClient interface {
	StreamStockPrices(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[StockRequest, StockResponse], error)
}

type stockServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStockServiceClient(cc grpc.ClientConnInterface) StockServiceClient {
	return &stockServiceClient{cc}
}

func (c *stockServiceClient) StreamStockPrices(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[StockRequest, StockResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &StockService_ServiceDesc.Streams[0], StockService_StreamStockPrices_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StockRequest, StockResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type StockService_StreamStockPricesClient = grpc.BidiStreamingClient[StockRequest, StockResponse]

// StockServiceServer is the server API for StockService service.
// All implementations must embed UnimplementedStockServiceServer
// for forward compatibility.
type StockServiceServer interface {
	StreamStockPrices(grpc.BidiStreamingServer[StockRequest, StockResponse]) error
	mustEmbedUnimplementedStockServiceServer()
}

// UnimplementedStockServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedStockServiceServer struct{}

func (UnimplementedStockServiceServer) StreamStockPrices(grpc.BidiStreamingServer[StockRequest, StockResponse]) error {
	return status.Errorf(codes.Unimplemented, "method StreamStockPrices not implemented")
}
func (UnimplementedStockServiceServer) mustEmbedUnimplementedStockServiceServer() {}
func (UnimplementedStockServiceServer) testEmbeddedByValue()                      {}

// UnsafeStockServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StockServiceServer will
// result in compilation errors.
type UnsafeStockServiceServer interface {
	mustEmbedUnimplementedStockServiceServer()
}

func RegisterStockServiceServer(s grpc.ServiceRegistrar, srv StockServiceServer) {
	// If the following call pancis, it indicates UnimplementedStockServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&StockService_ServiceDesc, srv)
}

func _StockService_StreamStockPrices_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StockServiceServer).StreamStockPrices(&grpc.GenericServerStream[StockRequest, StockResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type StockService_StreamStockPricesServer = grpc.BidiStreamingServer[StockRequest, StockResponse]

// StockService_ServiceDesc is the grpc.ServiceDesc for StockService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StockService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.StockService",
	HandlerType: (*StockServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamStockPrices",
			Handler:       _StockService_StreamStockPrices_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "stock.proto",
}
