// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: services/user_service/service/pb/user.proto

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	GetMyProfile(ctx context.Context, in *GetMyProfileReq, opts ...grpc.CallOption) (*GetMyProfileRes, error)
	SetMyProfile(ctx context.Context, in *SetMyProfileReq, opts ...grpc.CallOption) (*SetMyProfileRes, error)
	UploadProfile(ctx context.Context, opts ...grpc.CallOption) (UserService_UploadProfileClient, error)
	Download(ctx context.Context, in *GetProfilePicReq, opts ...grpc.CallOption) (UserService_DownloadClient, error)
	DeleteMyProfile(ctx context.Context, in *DeleteMyAccountReq, opts ...grpc.CallOption) (*DeleteMyAccountRes, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetMyProfile(ctx context.Context, in *GetMyProfileReq, opts ...grpc.CallOption) (*GetMyProfileRes, error) {
	out := new(GetMyProfileRes)
	err := c.cc.Invoke(ctx, "/auth.UserService/GetMyProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SetMyProfile(ctx context.Context, in *SetMyProfileReq, opts ...grpc.CallOption) (*SetMyProfileRes, error) {
	out := new(SetMyProfileRes)
	err := c.cc.Invoke(ctx, "/auth.UserService/SetMyProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UploadProfile(ctx context.Context, opts ...grpc.CallOption) (UserService_UploadProfileClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserService_ServiceDesc.Streams[0], "/auth.UserService/UploadProfile", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceUploadProfileClient{stream}
	return x, nil
}

type UserService_UploadProfileClient interface {
	Send(*UploadProfilePicReq) error
	CloseAndRecv() (*UploadProfilePicRes, error)
	grpc.ClientStream
}

type userServiceUploadProfileClient struct {
	grpc.ClientStream
}

func (x *userServiceUploadProfileClient) Send(m *UploadProfilePicReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *userServiceUploadProfileClient) CloseAndRecv() (*UploadProfilePicRes, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadProfilePicRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userServiceClient) Download(ctx context.Context, in *GetProfilePicReq, opts ...grpc.CallOption) (UserService_DownloadClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserService_ServiceDesc.Streams[1], "/auth.UserService/Download", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceDownloadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UserService_DownloadClient interface {
	Recv() (*GetProfilePicRes, error)
	grpc.ClientStream
}

type userServiceDownloadClient struct {
	grpc.ClientStream
}

func (x *userServiceDownloadClient) Recv() (*GetProfilePicRes, error) {
	m := new(GetProfilePicRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userServiceClient) DeleteMyProfile(ctx context.Context, in *DeleteMyAccountReq, opts ...grpc.CallOption) (*DeleteMyAccountRes, error) {
	out := new(DeleteMyAccountRes)
	err := c.cc.Invoke(ctx, "/auth.UserService/DeleteMyProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	GetMyProfile(context.Context, *GetMyProfileReq) (*GetMyProfileRes, error)
	SetMyProfile(context.Context, *SetMyProfileReq) (*SetMyProfileRes, error)
	UploadProfile(UserService_UploadProfileServer) error
	Download(*GetProfilePicReq, UserService_DownloadServer) error
	DeleteMyProfile(context.Context, *DeleteMyAccountReq) (*DeleteMyAccountRes, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) GetMyProfile(context.Context, *GetMyProfileReq) (*GetMyProfileRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyProfile not implemented")
}
func (UnimplementedUserServiceServer) SetMyProfile(context.Context, *SetMyProfileReq) (*SetMyProfileRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMyProfile not implemented")
}
func (UnimplementedUserServiceServer) UploadProfile(UserService_UploadProfileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadProfile not implemented")
}
func (UnimplementedUserServiceServer) Download(*GetProfilePicReq, UserService_DownloadServer) error {
	return status.Errorf(codes.Unimplemented, "method Download not implemented")
}
func (UnimplementedUserServiceServer) DeleteMyProfile(context.Context, *DeleteMyAccountReq) (*DeleteMyAccountRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMyProfile not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_GetMyProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMyProfileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetMyProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.UserService/GetMyProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetMyProfile(ctx, req.(*GetMyProfileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SetMyProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetMyProfileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SetMyProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.UserService/SetMyProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SetMyProfile(ctx, req.(*SetMyProfileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UploadProfile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UserServiceServer).UploadProfile(&userServiceUploadProfileServer{stream})
}

type UserService_UploadProfileServer interface {
	SendAndClose(*UploadProfilePicRes) error
	Recv() (*UploadProfilePicReq, error)
	grpc.ServerStream
}

type userServiceUploadProfileServer struct {
	grpc.ServerStream
}

func (x *userServiceUploadProfileServer) SendAndClose(m *UploadProfilePicRes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *userServiceUploadProfileServer) Recv() (*UploadProfilePicReq, error) {
	m := new(UploadProfilePicReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _UserService_Download_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetProfilePicReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserServiceServer).Download(m, &userServiceDownloadServer{stream})
}

type UserService_DownloadServer interface {
	Send(*GetProfilePicRes) error
	grpc.ServerStream
}

type userServiceDownloadServer struct {
	grpc.ServerStream
}

func (x *userServiceDownloadServer) Send(m *GetProfilePicRes) error {
	return x.ServerStream.SendMsg(m)
}

func _UserService_DeleteMyProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMyAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteMyProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.UserService/DeleteMyProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteMyProfile(ctx, req.(*DeleteMyAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMyProfile",
			Handler:    _UserService_GetMyProfile_Handler,
		},
		{
			MethodName: "SetMyProfile",
			Handler:    _UserService_SetMyProfile_Handler,
		},
		{
			MethodName: "DeleteMyProfile",
			Handler:    _UserService_DeleteMyProfile_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadProfile",
			Handler:       _UserService_UploadProfile_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Download",
			Handler:       _UserService_Download_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "services/user_service/service/pb/user.proto",
}
