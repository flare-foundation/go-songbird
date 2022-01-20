// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: gresponsewriter.proto

package gresponsewriterproto

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

// WriterClient is the client API for Writer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WriterClient interface {
	Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error)
	WriteHeader(ctx context.Context, in *WriteHeaderRequest, opts ...grpc.CallOption) (*WriteHeaderResponse, error)
	Flush(ctx context.Context, in *FlushRequest, opts ...grpc.CallOption) (*FlushResponse, error)
	Hijack(ctx context.Context, in *HijackRequest, opts ...grpc.CallOption) (*HijackResponse, error)
}

type writerClient struct {
	cc grpc.ClientConnInterface
}

func NewWriterClient(cc grpc.ClientConnInterface) WriterClient {
	return &writerClient{cc}
}

func (c *writerClient) Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := c.cc.Invoke(ctx, "/gresponsewriterproto.Writer/Write", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *writerClient) WriteHeader(ctx context.Context, in *WriteHeaderRequest, opts ...grpc.CallOption) (*WriteHeaderResponse, error) {
	out := new(WriteHeaderResponse)
	err := c.cc.Invoke(ctx, "/gresponsewriterproto.Writer/WriteHeader", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *writerClient) Flush(ctx context.Context, in *FlushRequest, opts ...grpc.CallOption) (*FlushResponse, error) {
	out := new(FlushResponse)
	err := c.cc.Invoke(ctx, "/gresponsewriterproto.Writer/Flush", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *writerClient) Hijack(ctx context.Context, in *HijackRequest, opts ...grpc.CallOption) (*HijackResponse, error) {
	out := new(HijackResponse)
	err := c.cc.Invoke(ctx, "/gresponsewriterproto.Writer/Hijack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WriterServer is the server API for Writer service.
// All implementations must embed UnimplementedWriterServer
// for forward compatibility
type WriterServer interface {
	Write(context.Context, *WriteRequest) (*WriteResponse, error)
	WriteHeader(context.Context, *WriteHeaderRequest) (*WriteHeaderResponse, error)
	Flush(context.Context, *FlushRequest) (*FlushResponse, error)
	Hijack(context.Context, *HijackRequest) (*HijackResponse, error)
	mustEmbedUnimplementedWriterServer()
}

// UnimplementedWriterServer must be embedded to have forward compatible implementations.
type UnimplementedWriterServer struct {
}

func (UnimplementedWriterServer) Write(context.Context, *WriteRequest) (*WriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}
func (UnimplementedWriterServer) WriteHeader(context.Context, *WriteHeaderRequest) (*WriteHeaderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteHeader not implemented")
}
func (UnimplementedWriterServer) Flush(context.Context, *FlushRequest) (*FlushResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Flush not implemented")
}
func (UnimplementedWriterServer) Hijack(context.Context, *HijackRequest) (*HijackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hijack not implemented")
}
func (UnimplementedWriterServer) mustEmbedUnimplementedWriterServer() {}

// UnsafeWriterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WriterServer will
// result in compilation errors.
type UnsafeWriterServer interface {
	mustEmbedUnimplementedWriterServer()
}

func RegisterWriterServer(s grpc.ServiceRegistrar, srv WriterServer) {
	s.RegisterService(&Writer_ServiceDesc, srv)
}

func _Writer_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WriterServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gresponsewriterproto.Writer/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WriterServer).Write(ctx, req.(*WriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Writer_WriteHeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteHeaderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WriterServer).WriteHeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gresponsewriterproto.Writer/WriteHeader",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WriterServer).WriteHeader(ctx, req.(*WriteHeaderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Writer_Flush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlushRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WriterServer).Flush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gresponsewriterproto.Writer/Flush",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WriterServer).Flush(ctx, req.(*FlushRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Writer_Hijack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HijackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WriterServer).Hijack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gresponsewriterproto.Writer/Hijack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WriterServer).Hijack(ctx, req.(*HijackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Writer_ServiceDesc is the grpc.ServiceDesc for Writer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Writer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gresponsewriterproto.Writer",
	HandlerType: (*WriterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Write",
			Handler:    _Writer_Write_Handler,
		},
		{
			MethodName: "WriteHeader",
			Handler:    _Writer_WriteHeader_Handler,
		},
		{
			MethodName: "Flush",
			Handler:    _Writer_Flush_Handler,
		},
		{
			MethodName: "Hijack",
			Handler:    _Writer_Hijack_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gresponsewriter.proto",
}
