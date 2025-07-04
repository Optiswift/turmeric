// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: exchange_rate_alert.proto

package proto

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
	ExchangeRateAlertService_CreateExchangeRateAlert_FullMethodName  = "/optiswift.proto.mint.ExchangeRateAlertService/CreateExchangeRateAlert"
	ExchangeRateAlertService_GetExchangeRateAlerts_FullMethodName    = "/optiswift.proto.mint.ExchangeRateAlertService/GetExchangeRateAlerts"
	ExchangeRateAlertService_DeleteExchangeRateAlerts_FullMethodName = "/optiswift.proto.mint.ExchangeRateAlertService/DeleteExchangeRateAlerts"
)

// ExchangeRateAlertServiceClient is the client API for ExchangeRateAlertService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExchangeRateAlertServiceClient interface {
	CreateExchangeRateAlert(ctx context.Context, in *CreateExchangeRateAlertRequest, opts ...grpc.CallOption) (*ExchangeRateAlert, error)
	GetExchangeRateAlerts(ctx context.Context, in *GetExchangeRateAlertsRequest, opts ...grpc.CallOption) (*GetExchangeRateAlertsResponse, error)
	DeleteExchangeRateAlerts(ctx context.Context, in *DeleteExchangeRateAlertsRequest, opts ...grpc.CallOption) (*DeleteExchangeRateAlertsResponse, error)
}

type exchangeRateAlertServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExchangeRateAlertServiceClient(cc grpc.ClientConnInterface) ExchangeRateAlertServiceClient {
	return &exchangeRateAlertServiceClient{cc}
}

func (c *exchangeRateAlertServiceClient) CreateExchangeRateAlert(ctx context.Context, in *CreateExchangeRateAlertRequest, opts ...grpc.CallOption) (*ExchangeRateAlert, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExchangeRateAlert)
	err := c.cc.Invoke(ctx, ExchangeRateAlertService_CreateExchangeRateAlert_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeRateAlertServiceClient) GetExchangeRateAlerts(ctx context.Context, in *GetExchangeRateAlertsRequest, opts ...grpc.CallOption) (*GetExchangeRateAlertsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetExchangeRateAlertsResponse)
	err := c.cc.Invoke(ctx, ExchangeRateAlertService_GetExchangeRateAlerts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeRateAlertServiceClient) DeleteExchangeRateAlerts(ctx context.Context, in *DeleteExchangeRateAlertsRequest, opts ...grpc.CallOption) (*DeleteExchangeRateAlertsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteExchangeRateAlertsResponse)
	err := c.cc.Invoke(ctx, ExchangeRateAlertService_DeleteExchangeRateAlerts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExchangeRateAlertServiceServer is the server API for ExchangeRateAlertService service.
// All implementations must embed UnimplementedExchangeRateAlertServiceServer
// for forward compatibility.
type ExchangeRateAlertServiceServer interface {
	CreateExchangeRateAlert(context.Context, *CreateExchangeRateAlertRequest) (*ExchangeRateAlert, error)
	GetExchangeRateAlerts(context.Context, *GetExchangeRateAlertsRequest) (*GetExchangeRateAlertsResponse, error)
	DeleteExchangeRateAlerts(context.Context, *DeleteExchangeRateAlertsRequest) (*DeleteExchangeRateAlertsResponse, error)
	mustEmbedUnimplementedExchangeRateAlertServiceServer()
}

// UnimplementedExchangeRateAlertServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedExchangeRateAlertServiceServer struct{}

func (UnimplementedExchangeRateAlertServiceServer) CreateExchangeRateAlert(context.Context, *CreateExchangeRateAlertRequest) (*ExchangeRateAlert, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateExchangeRateAlert not implemented")
}
func (UnimplementedExchangeRateAlertServiceServer) GetExchangeRateAlerts(context.Context, *GetExchangeRateAlertsRequest) (*GetExchangeRateAlertsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExchangeRateAlerts not implemented")
}
func (UnimplementedExchangeRateAlertServiceServer) DeleteExchangeRateAlerts(context.Context, *DeleteExchangeRateAlertsRequest) (*DeleteExchangeRateAlertsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteExchangeRateAlerts not implemented")
}
func (UnimplementedExchangeRateAlertServiceServer) mustEmbedUnimplementedExchangeRateAlertServiceServer() {
}
func (UnimplementedExchangeRateAlertServiceServer) testEmbeddedByValue() {}

// UnsafeExchangeRateAlertServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExchangeRateAlertServiceServer will
// result in compilation errors.
type UnsafeExchangeRateAlertServiceServer interface {
	mustEmbedUnimplementedExchangeRateAlertServiceServer()
}

func RegisterExchangeRateAlertServiceServer(s grpc.ServiceRegistrar, srv ExchangeRateAlertServiceServer) {
	// If the following call pancis, it indicates UnimplementedExchangeRateAlertServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ExchangeRateAlertService_ServiceDesc, srv)
}

func _ExchangeRateAlertService_CreateExchangeRateAlert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateExchangeRateAlertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeRateAlertServiceServer).CreateExchangeRateAlert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExchangeRateAlertService_CreateExchangeRateAlert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeRateAlertServiceServer).CreateExchangeRateAlert(ctx, req.(*CreateExchangeRateAlertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExchangeRateAlertService_GetExchangeRateAlerts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExchangeRateAlertsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeRateAlertServiceServer).GetExchangeRateAlerts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExchangeRateAlertService_GetExchangeRateAlerts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeRateAlertServiceServer).GetExchangeRateAlerts(ctx, req.(*GetExchangeRateAlertsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExchangeRateAlertService_DeleteExchangeRateAlerts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteExchangeRateAlertsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeRateAlertServiceServer).DeleteExchangeRateAlerts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExchangeRateAlertService_DeleteExchangeRateAlerts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeRateAlertServiceServer).DeleteExchangeRateAlerts(ctx, req.(*DeleteExchangeRateAlertsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExchangeRateAlertService_ServiceDesc is the grpc.ServiceDesc for ExchangeRateAlertService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExchangeRateAlertService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "optiswift.proto.mint.ExchangeRateAlertService",
	HandlerType: (*ExchangeRateAlertServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateExchangeRateAlert",
			Handler:    _ExchangeRateAlertService_CreateExchangeRateAlert_Handler,
		},
		{
			MethodName: "GetExchangeRateAlerts",
			Handler:    _ExchangeRateAlertService_GetExchangeRateAlerts_Handler,
		},
		{
			MethodName: "DeleteExchangeRateAlerts",
			Handler:    _ExchangeRateAlertService_DeleteExchangeRateAlerts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "exchange_rate_alert.proto",
}
