// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package core

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ConnectorClient is the client API for Connector service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConnectorClient interface {
	Call(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type connectorClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectorClient(cc grpc.ClientConnInterface) ConnectorClient {
	return &connectorClient{cc}
}

func (c *connectorClient) Call(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/core.Connector/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectorServer is the server API for Connector service.
// All implementations must embed UnimplementedConnectorServer
// for forward compatibility
type ConnectorServer interface {
	Call(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedConnectorServer()
}

// UnimplementedConnectorServer must be embedded to have forward compatible implementations.
type UnimplementedConnectorServer struct {
}

func (UnimplementedConnectorServer) Call(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}
func (UnimplementedConnectorServer) mustEmbedUnimplementedConnectorServer() {}

// UnsafeConnectorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConnectorServer will
// result in compilation errors.
type UnsafeConnectorServer interface {
	mustEmbedUnimplementedConnectorServer()
}

func RegisterConnectorServer(s grpc.ServiceRegistrar, srv ConnectorServer) {
	s.RegisterService(&_Connector_serviceDesc, srv)
}

func _Connector_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.Connector/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorServer).Call(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Connector_serviceDesc = grpc.ServiceDesc{
	ServiceName: "core.Connector",
	HandlerType: (*ConnectorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _Connector_Call_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "net.proto",
}
