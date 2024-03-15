// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: apis/serviceconn/gateway_file_service/pb/gw_file_svc.proto

package pb

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

// UploadBlogFileClient is the client API for UploadBlogFile service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UploadBlogFileClient interface {
	// Lets an user upload profile pic into the file server
	UploadProfilePic(ctx context.Context, opts ...grpc.CallOption) (UploadBlogFile_UploadProfilePicClient, error)
	// Lets an user get profile pic into the file server
	GetProfilePic(ctx context.Context, in *GetProfilePicReq, opts ...grpc.CallOption) (UploadBlogFile_GetProfilePicClient, error)
	// Lets a user delete the profile picture
	DeleteProfilePic(ctx context.Context, in *DeleteProfilePicReq, opts ...grpc.CallOption) (*DeleteProfilePicRes, error)
	UploadBlogFile(ctx context.Context, opts ...grpc.CallOption) (UploadBlogFile_UploadBlogFileClient, error)
	GetBlogFile(ctx context.Context, in *GetBlogFileReq, opts ...grpc.CallOption) (UploadBlogFile_GetBlogFileClient, error)
	DeleteBlogFile(ctx context.Context, in *DeleteBlogFileReq, opts ...grpc.CallOption) (*DeleteBlogFileRes, error)
}

type uploadBlogFileClient struct {
	cc grpc.ClientConnInterface
}

func NewUploadBlogFileClient(cc grpc.ClientConnInterface) UploadBlogFileClient {
	return &uploadBlogFileClient{cc}
}

func (c *uploadBlogFileClient) UploadProfilePic(ctx context.Context, opts ...grpc.CallOption) (UploadBlogFile_UploadProfilePicClient, error) {
	stream, err := c.cc.NewStream(ctx, &UploadBlogFile_ServiceDesc.Streams[0], "/auth_svc.UploadBlogFile/UploadProfilePic", opts...)
	if err != nil {
		return nil, err
	}
	x := &uploadBlogFileUploadProfilePicClient{stream}
	return x, nil
}

type UploadBlogFile_UploadProfilePicClient interface {
	Send(*UploadProfilePicReq) error
	CloseAndRecv() (*UploadProfilePicRes, error)
	grpc.ClientStream
}

type uploadBlogFileUploadProfilePicClient struct {
	grpc.ClientStream
}

func (x *uploadBlogFileUploadProfilePicClient) Send(m *UploadProfilePicReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *uploadBlogFileUploadProfilePicClient) CloseAndRecv() (*UploadProfilePicRes, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadProfilePicRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *uploadBlogFileClient) GetProfilePic(ctx context.Context, in *GetProfilePicReq, opts ...grpc.CallOption) (UploadBlogFile_GetProfilePicClient, error) {
	stream, err := c.cc.NewStream(ctx, &UploadBlogFile_ServiceDesc.Streams[1], "/auth_svc.UploadBlogFile/GetProfilePic", opts...)
	if err != nil {
		return nil, err
	}
	x := &uploadBlogFileGetProfilePicClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UploadBlogFile_GetProfilePicClient interface {
	Recv() (*GetProfilePicRes, error)
	grpc.ClientStream
}

type uploadBlogFileGetProfilePicClient struct {
	grpc.ClientStream
}

func (x *uploadBlogFileGetProfilePicClient) Recv() (*GetProfilePicRes, error) {
	m := new(GetProfilePicRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *uploadBlogFileClient) DeleteProfilePic(ctx context.Context, in *DeleteProfilePicReq, opts ...grpc.CallOption) (*DeleteProfilePicRes, error) {
	out := new(DeleteProfilePicRes)
	err := c.cc.Invoke(ctx, "/auth_svc.UploadBlogFile/DeleteProfilePic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uploadBlogFileClient) UploadBlogFile(ctx context.Context, opts ...grpc.CallOption) (UploadBlogFile_UploadBlogFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &UploadBlogFile_ServiceDesc.Streams[2], "/auth_svc.UploadBlogFile/UploadBlogFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &uploadBlogFileUploadBlogFileClient{stream}
	return x, nil
}

type UploadBlogFile_UploadBlogFileClient interface {
	Send(*UploadBlogFileReq) error
	CloseAndRecv() (*UploadBlogFileRes, error)
	grpc.ClientStream
}

type uploadBlogFileUploadBlogFileClient struct {
	grpc.ClientStream
}

func (x *uploadBlogFileUploadBlogFileClient) Send(m *UploadBlogFileReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *uploadBlogFileUploadBlogFileClient) CloseAndRecv() (*UploadBlogFileRes, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadBlogFileRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *uploadBlogFileClient) GetBlogFile(ctx context.Context, in *GetBlogFileReq, opts ...grpc.CallOption) (UploadBlogFile_GetBlogFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &UploadBlogFile_ServiceDesc.Streams[3], "/auth_svc.UploadBlogFile/GetBlogFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &uploadBlogFileGetBlogFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UploadBlogFile_GetBlogFileClient interface {
	Recv() (*GetBlogFileRes, error)
	grpc.ClientStream
}

type uploadBlogFileGetBlogFileClient struct {
	grpc.ClientStream
}

func (x *uploadBlogFileGetBlogFileClient) Recv() (*GetBlogFileRes, error) {
	m := new(GetBlogFileRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *uploadBlogFileClient) DeleteBlogFile(ctx context.Context, in *DeleteBlogFileReq, opts ...grpc.CallOption) (*DeleteBlogFileRes, error) {
	out := new(DeleteBlogFileRes)
	err := c.cc.Invoke(ctx, "/auth_svc.UploadBlogFile/DeleteBlogFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UploadBlogFileServer is the server API for UploadBlogFile service.
// All implementations must embed UnimplementedUploadBlogFileServer
// for forward compatibility
type UploadBlogFileServer interface {
	// Lets an user upload profile pic into the file server
	UploadProfilePic(UploadBlogFile_UploadProfilePicServer) error
	// Lets an user get profile pic into the file server
	GetProfilePic(*GetProfilePicReq, UploadBlogFile_GetProfilePicServer) error
	// Lets a user delete the profile picture
	DeleteProfilePic(context.Context, *DeleteProfilePicReq) (*DeleteProfilePicRes, error)
	UploadBlogFile(UploadBlogFile_UploadBlogFileServer) error
	GetBlogFile(*GetBlogFileReq, UploadBlogFile_GetBlogFileServer) error
	DeleteBlogFile(context.Context, *DeleteBlogFileReq) (*DeleteBlogFileRes, error)
	mustEmbedUnimplementedUploadBlogFileServer()
}

// UnimplementedUploadBlogFileServer must be embedded to have forward compatible implementations.
type UnimplementedUploadBlogFileServer struct {
}

func (UnimplementedUploadBlogFileServer) UploadProfilePic(UploadBlogFile_UploadProfilePicServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadProfilePic not implemented")
}
func (UnimplementedUploadBlogFileServer) GetProfilePic(*GetProfilePicReq, UploadBlogFile_GetProfilePicServer) error {
	return status.Errorf(codes.Unimplemented, "method GetProfilePic not implemented")
}
func (UnimplementedUploadBlogFileServer) DeleteProfilePic(context.Context, *DeleteProfilePicReq) (*DeleteProfilePicRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProfilePic not implemented")
}
func (UnimplementedUploadBlogFileServer) UploadBlogFile(UploadBlogFile_UploadBlogFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadBlogFile not implemented")
}
func (UnimplementedUploadBlogFileServer) GetBlogFile(*GetBlogFileReq, UploadBlogFile_GetBlogFileServer) error {
	return status.Errorf(codes.Unimplemented, "method GetBlogFile not implemented")
}
func (UnimplementedUploadBlogFileServer) DeleteBlogFile(context.Context, *DeleteBlogFileReq) (*DeleteBlogFileRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBlogFile not implemented")
}
func (UnimplementedUploadBlogFileServer) mustEmbedUnimplementedUploadBlogFileServer() {}

// UnsafeUploadBlogFileServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UploadBlogFileServer will
// result in compilation errors.
type UnsafeUploadBlogFileServer interface {
	mustEmbedUnimplementedUploadBlogFileServer()
}

func RegisterUploadBlogFileServer(s grpc.ServiceRegistrar, srv UploadBlogFileServer) {
	s.RegisterService(&UploadBlogFile_ServiceDesc, srv)
}

func _UploadBlogFile_UploadProfilePic_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UploadBlogFileServer).UploadProfilePic(&uploadBlogFileUploadProfilePicServer{stream})
}

type UploadBlogFile_UploadProfilePicServer interface {
	SendAndClose(*UploadProfilePicRes) error
	Recv() (*UploadProfilePicReq, error)
	grpc.ServerStream
}

type uploadBlogFileUploadProfilePicServer struct {
	grpc.ServerStream
}

func (x *uploadBlogFileUploadProfilePicServer) SendAndClose(m *UploadProfilePicRes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *uploadBlogFileUploadProfilePicServer) Recv() (*UploadProfilePicReq, error) {
	m := new(UploadProfilePicReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _UploadBlogFile_GetProfilePic_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetProfilePicReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UploadBlogFileServer).GetProfilePic(m, &uploadBlogFileGetProfilePicServer{stream})
}

type UploadBlogFile_GetProfilePicServer interface {
	Send(*GetProfilePicRes) error
	grpc.ServerStream
}

type uploadBlogFileGetProfilePicServer struct {
	grpc.ServerStream
}

func (x *uploadBlogFileGetProfilePicServer) Send(m *GetProfilePicRes) error {
	return x.ServerStream.SendMsg(m)
}

func _UploadBlogFile_DeleteProfilePic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProfilePicReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UploadBlogFileServer).DeleteProfilePic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_svc.UploadBlogFile/DeleteProfilePic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UploadBlogFileServer).DeleteProfilePic(ctx, req.(*DeleteProfilePicReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UploadBlogFile_UploadBlogFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UploadBlogFileServer).UploadBlogFile(&uploadBlogFileUploadBlogFileServer{stream})
}

type UploadBlogFile_UploadBlogFileServer interface {
	SendAndClose(*UploadBlogFileRes) error
	Recv() (*UploadBlogFileReq, error)
	grpc.ServerStream
}

type uploadBlogFileUploadBlogFileServer struct {
	grpc.ServerStream
}

func (x *uploadBlogFileUploadBlogFileServer) SendAndClose(m *UploadBlogFileRes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *uploadBlogFileUploadBlogFileServer) Recv() (*UploadBlogFileReq, error) {
	m := new(UploadBlogFileReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _UploadBlogFile_GetBlogFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetBlogFileReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UploadBlogFileServer).GetBlogFile(m, &uploadBlogFileGetBlogFileServer{stream})
}

type UploadBlogFile_GetBlogFileServer interface {
	Send(*GetBlogFileRes) error
	grpc.ServerStream
}

type uploadBlogFileGetBlogFileServer struct {
	grpc.ServerStream
}

func (x *uploadBlogFileGetBlogFileServer) Send(m *GetBlogFileRes) error {
	return x.ServerStream.SendMsg(m)
}

func _UploadBlogFile_DeleteBlogFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBlogFileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UploadBlogFileServer).DeleteBlogFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_svc.UploadBlogFile/DeleteBlogFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UploadBlogFileServer).DeleteBlogFile(ctx, req.(*DeleteBlogFileReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UploadBlogFile_ServiceDesc is the grpc.ServiceDesc for UploadBlogFile service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UploadBlogFile_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth_svc.UploadBlogFile",
	HandlerType: (*UploadBlogFileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteProfilePic",
			Handler:    _UploadBlogFile_DeleteProfilePic_Handler,
		},
		{
			MethodName: "DeleteBlogFile",
			Handler:    _UploadBlogFile_DeleteBlogFile_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadProfilePic",
			Handler:       _UploadBlogFile_UploadProfilePic_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetProfilePic",
			Handler:       _UploadBlogFile_GetProfilePic_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "UploadBlogFile",
			Handler:       _UploadBlogFile_UploadBlogFile_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetBlogFile",
			Handler:       _UploadBlogFile_GetBlogFile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "apis/serviceconn/gateway_file_service/pb/gw_file_svc.proto",
}