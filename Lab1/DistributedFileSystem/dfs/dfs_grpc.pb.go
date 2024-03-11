// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: dfs/dfs.proto

package dfs

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

// DFSClient is the client API for DFS service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DFSClient interface {
	// 1-2 client - master tracker
	RequestToUpload(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RequestToUploadResponse, error)
	// 3-4 client - data keeper
	UploadFile(ctx context.Context, in *UploadFileRequest, opts ...grpc.CallOption) (*Empty, error)
	// 5 data keeper - master tracker
	UploadSuccess(ctx context.Context, in *UploadSuccessRequest, opts ...grpc.CallOption) (*Empty, error)
	// 6 master tracker - client
	NotifyClient(ctx context.Context, in *NotifyClientRequest, opts ...grpc.CallOption) (*Empty, error)
	// client - master tracker
	RequestToDownload(ctx context.Context, in *RequestToDownloadRequest, opts ...grpc.CallOption) (*RequestToDownloadResponse, error)
	// client - data keeper
	DownloadFile(ctx context.Context, in *DownloadFileRequest, opts ...grpc.CallOption) (*DownloadFileResponse, error)
	// Heartbeats: master tracker - data keeper
	PingMasterTracker(ctx context.Context, in *PingMasterTrackerRequest, opts ...grpc.CallOption) (*Empty, error)
}

type dFSClient struct {
	cc grpc.ClientConnInterface
}

func NewDFSClient(cc grpc.ClientConnInterface) DFSClient {
	return &dFSClient{cc}
}

func (c *dFSClient) RequestToUpload(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RequestToUploadResponse, error) {
	out := new(RequestToUploadResponse)
	err := c.cc.Invoke(ctx, "/dfs.DFS/RequestToUpload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dFSClient) UploadFile(ctx context.Context, in *UploadFileRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/dfs.DFS/UploadFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dFSClient) UploadSuccess(ctx context.Context, in *UploadSuccessRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/dfs.DFS/UploadSuccess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dFSClient) NotifyClient(ctx context.Context, in *NotifyClientRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/dfs.DFS/NotifyClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dFSClient) RequestToDownload(ctx context.Context, in *RequestToDownloadRequest, opts ...grpc.CallOption) (*RequestToDownloadResponse, error) {
	out := new(RequestToDownloadResponse)
	err := c.cc.Invoke(ctx, "/dfs.DFS/RequestToDownload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dFSClient) DownloadFile(ctx context.Context, in *DownloadFileRequest, opts ...grpc.CallOption) (*DownloadFileResponse, error) {
	out := new(DownloadFileResponse)
	err := c.cc.Invoke(ctx, "/dfs.DFS/DownloadFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dFSClient) PingMasterTracker(ctx context.Context, in *PingMasterTrackerRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/dfs.DFS/PingMasterTracker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DFSServer is the server API for DFS service.
// All implementations must embed UnimplementedDFSServer
// for forward compatibility
type DFSServer interface {
	// 1-2 client - master tracker
	RequestToUpload(context.Context, *Empty) (*RequestToUploadResponse, error)
	// 3-4 client - data keeper
	UploadFile(context.Context, *UploadFileRequest) (*Empty, error)
	// 5 data keeper - master tracker
	UploadSuccess(context.Context, *UploadSuccessRequest) (*Empty, error)
	// 6 master tracker - client
	NotifyClient(context.Context, *NotifyClientRequest) (*Empty, error)
	// client - master tracker
	RequestToDownload(context.Context, *RequestToDownloadRequest) (*RequestToDownloadResponse, error)
	// client - data keeper
	DownloadFile(context.Context, *DownloadFileRequest) (*DownloadFileResponse, error)
	// Heartbeats: master tracker - data keeper
	PingMasterTracker(context.Context, *PingMasterTrackerRequest) (*Empty, error)
	mustEmbedUnimplementedDFSServer()
}

// UnimplementedDFSServer must be embedded to have forward compatible implementations.
type UnimplementedDFSServer struct {
}

func (UnimplementedDFSServer) RequestToUpload(context.Context, *Empty) (*RequestToUploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestToUpload not implemented")
}
func (UnimplementedDFSServer) UploadFile(context.Context, *UploadFileRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedDFSServer) UploadSuccess(context.Context, *UploadSuccessRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadSuccess not implemented")
}
func (UnimplementedDFSServer) NotifyClient(context.Context, *NotifyClientRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyClient not implemented")
}
func (UnimplementedDFSServer) RequestToDownload(context.Context, *RequestToDownloadRequest) (*RequestToDownloadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestToDownload not implemented")
}
func (UnimplementedDFSServer) DownloadFile(context.Context, *DownloadFileRequest) (*DownloadFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
}
func (UnimplementedDFSServer) PingMasterTracker(context.Context, *PingMasterTrackerRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingMasterTracker not implemented")
}
func (UnimplementedDFSServer) mustEmbedUnimplementedDFSServer() {}

// UnsafeDFSServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DFSServer will
// result in compilation errors.
type UnsafeDFSServer interface {
	mustEmbedUnimplementedDFSServer()
}

func RegisterDFSServer(s grpc.ServiceRegistrar, srv DFSServer) {
	s.RegisterService(&DFS_ServiceDesc, srv)
}

func _DFS_RequestToUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DFSServer).RequestToUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dfs.DFS/RequestToUpload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DFSServer).RequestToUpload(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DFS_UploadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DFSServer).UploadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dfs.DFS/UploadFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DFSServer).UploadFile(ctx, req.(*UploadFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DFS_UploadSuccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadSuccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DFSServer).UploadSuccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dfs.DFS/UploadSuccess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DFSServer).UploadSuccess(ctx, req.(*UploadSuccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DFS_NotifyClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DFSServer).NotifyClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dfs.DFS/NotifyClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DFSServer).NotifyClient(ctx, req.(*NotifyClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DFS_RequestToDownload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestToDownloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DFSServer).RequestToDownload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dfs.DFS/RequestToDownload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DFSServer).RequestToDownload(ctx, req.(*RequestToDownloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DFS_DownloadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DFSServer).DownloadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dfs.DFS/DownloadFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DFSServer).DownloadFile(ctx, req.(*DownloadFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DFS_PingMasterTracker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingMasterTrackerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DFSServer).PingMasterTracker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dfs.DFS/PingMasterTracker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DFSServer).PingMasterTracker(ctx, req.(*PingMasterTrackerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DFS_ServiceDesc is the grpc.ServiceDesc for DFS service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DFS_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dfs.DFS",
	HandlerType: (*DFSServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestToUpload",
			Handler:    _DFS_RequestToUpload_Handler,
		},
		{
			MethodName: "UploadFile",
			Handler:    _DFS_UploadFile_Handler,
		},
		{
			MethodName: "UploadSuccess",
			Handler:    _DFS_UploadSuccess_Handler,
		},
		{
			MethodName: "NotifyClient",
			Handler:    _DFS_NotifyClient_Handler,
		},
		{
			MethodName: "RequestToDownload",
			Handler:    _DFS_RequestToDownload_Handler,
		},
		{
			MethodName: "DownloadFile",
			Handler:    _DFS_DownloadFile_Handler,
		},
		{
			MethodName: "PingMasterTracker",
			Handler:    _DFS_PingMasterTracker_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dfs/dfs.proto",
}
