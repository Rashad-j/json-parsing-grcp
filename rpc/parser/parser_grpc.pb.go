// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: rpc/parser/parser.proto

package parser

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

const (
	JsonParsingService_ParseJsonFiles_FullMethodName = "/JsonParsingService/ParseJsonFiles"
)

// JsonParsingServiceClient is the client API for JsonParsingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JsonParsingServiceClient interface {
	ParseJsonFiles(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*JsonResponse, error)
}

type jsonParsingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJsonParsingServiceClient(cc grpc.ClientConnInterface) JsonParsingServiceClient {
	return &jsonParsingServiceClient{cc}
}

func (c *jsonParsingServiceClient) ParseJsonFiles(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*JsonResponse, error) {
	out := new(JsonResponse)
	err := c.cc.Invoke(ctx, JsonParsingService_ParseJsonFiles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JsonParsingServiceServer is the server API for JsonParsingService service.
// All implementations must embed UnimplementedJsonParsingServiceServer
// for forward compatibility
type JsonParsingServiceServer interface {
	ParseJsonFiles(context.Context, *EmptyRequest) (*JsonResponse, error)
	mustEmbedUnimplementedJsonParsingServiceServer()
}

// UnimplementedJsonParsingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedJsonParsingServiceServer struct {
}

func (UnimplementedJsonParsingServiceServer) ParseJsonFiles(context.Context, *EmptyRequest) (*JsonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParseJsonFiles not implemented")
}
func (UnimplementedJsonParsingServiceServer) mustEmbedUnimplementedJsonParsingServiceServer() {}

// UnsafeJsonParsingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JsonParsingServiceServer will
// result in compilation errors.
type UnsafeJsonParsingServiceServer interface {
	mustEmbedUnimplementedJsonParsingServiceServer()
}

func RegisterJsonParsingServiceServer(s grpc.ServiceRegistrar, srv JsonParsingServiceServer) {
	s.RegisterService(&JsonParsingService_ServiceDesc, srv)
}

func _JsonParsingService_ParseJsonFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JsonParsingServiceServer).ParseJsonFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JsonParsingService_ParseJsonFiles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JsonParsingServiceServer).ParseJsonFiles(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// JsonParsingService_ServiceDesc is the grpc.ServiceDesc for JsonParsingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JsonParsingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "JsonParsingService",
	HandlerType: (*JsonParsingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ParseJsonFiles",
			Handler:    _JsonParsingService_ParseJsonFiles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/parser/parser.proto",
}
