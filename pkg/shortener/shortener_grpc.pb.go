// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: shortener.proto

package shortener

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

// ShortenerClient is the client API for Shortener service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShortenerClient interface {
	GetShortUrl(ctx context.Context, in *GetShortUrlRequest, opts ...grpc.CallOption) (*GetShortUrlResponse, error)
	GetOriginalUrl(ctx context.Context, in *GetOriginalUrlRequest, opts ...grpc.CallOption) (*GetOriginalUrlResponse, error)
}

type shortenerClient struct {
	cc grpc.ClientConnInterface
}

func NewShortenerClient(cc grpc.ClientConnInterface) ShortenerClient {
	return &shortenerClient{cc}
}

func (c *shortenerClient) GetShortUrl(ctx context.Context, in *GetShortUrlRequest, opts ...grpc.CallOption) (*GetShortUrlResponse, error) {
	out := new(GetShortUrlResponse)
	err := c.cc.Invoke(ctx, "/protos.shortener.Shortener/GetShortUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortenerClient) GetOriginalUrl(ctx context.Context, in *GetOriginalUrlRequest, opts ...grpc.CallOption) (*GetOriginalUrlResponse, error) {
	out := new(GetOriginalUrlResponse)
	err := c.cc.Invoke(ctx, "/protos.shortener.Shortener/GetOriginalUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShortenerServer is the server API for Shortener service.
// All implementations must embed UnimplementedShortenerServer
// for forward compatibility
type ShortenerServer interface {
	GetShortUrl(context.Context, *GetShortUrlRequest) (*GetShortUrlResponse, error)
	GetOriginalUrl(context.Context, *GetOriginalUrlRequest) (*GetOriginalUrlResponse, error)
	mustEmbedUnimplementedShortenerServer()
}

// UnimplementedShortenerServer must be embedded to have forward compatible implementations.
type UnimplementedShortenerServer struct {
}

func (UnimplementedShortenerServer) GetShortUrl(context.Context, *GetShortUrlRequest) (*GetShortUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShortUrl not implemented")
}
func (UnimplementedShortenerServer) GetOriginalUrl(context.Context, *GetOriginalUrlRequest) (*GetOriginalUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOriginalUrl not implemented")
}
func (UnimplementedShortenerServer) mustEmbedUnimplementedShortenerServer() {}

// UnsafeShortenerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShortenerServer will
// result in compilation errors.
type UnsafeShortenerServer interface {
	mustEmbedUnimplementedShortenerServer()
}

func RegisterShortenerServer(s grpc.ServiceRegistrar, srv ShortenerServer) {
	s.RegisterService(&Shortener_ServiceDesc, srv)
}

func _Shortener_GetShortUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShortUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerServer).GetShortUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.shortener.Shortener/GetShortUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerServer).GetShortUrl(ctx, req.(*GetShortUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shortener_GetOriginalUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOriginalUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerServer).GetOriginalUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.shortener.Shortener/GetOriginalUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerServer).GetOriginalUrl(ctx, req.(*GetOriginalUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Shortener_ServiceDesc is the grpc.ServiceDesc for Shortener service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Shortener_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.shortener.Shortener",
	HandlerType: (*ShortenerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetShortUrl",
			Handler:    _Shortener_GetShortUrl_Handler,
		},
		{
			MethodName: "GetOriginalUrl",
			Handler:    _Shortener_GetOriginalUrl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shortener.proto",
}
