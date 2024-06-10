// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: apis/serviceconn/gateway_user/pb/gw_user.proto

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
	GetUserActivities(ctx context.Context, in *UserActivityReq, opts ...grpc.CallOption) (*UserActivityRes, error)
	GetUserProfile(ctx context.Context, in *UserProfileReq, opts ...grpc.CallOption) (*UserProfileRes, error)
	UpdateUserProfile(ctx context.Context, in *UpdateUserProfileReq, opts ...grpc.CallOption) (*UpdateUserProfileRes, error)
	DeleteUserProfile(ctx context.Context, in *DeleteUserProfileReq, opts ...grpc.CallOption) (*DeleteUserProfileRes, error)
	GetAllTopics(ctx context.Context, in *GetTopicsRequests, opts ...grpc.CallOption) (*GetTopicsResponse, error)
	GetAllCategories(ctx context.Context, in *GetAllCategoriesReq, opts ...grpc.CallOption) (*GetAllCategoriesRes, error)
	UpdateUsername(ctx context.Context, in *UpdateUsernameReq, opts ...grpc.CallOption) (*UpdateUsernameRes, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUserActivities(ctx context.Context, in *UserActivityReq, opts ...grpc.CallOption) (*UserActivityRes, error) {
	out := new(UserActivityRes)
	err := c.cc.Invoke(ctx, "/auth_svc.UserService/GetUserActivities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserProfile(ctx context.Context, in *UserProfileReq, opts ...grpc.CallOption) (*UserProfileRes, error) {
	out := new(UserProfileRes)
	err := c.cc.Invoke(ctx, "/auth_svc.UserService/GetUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserProfile(ctx context.Context, in *UpdateUserProfileReq, opts ...grpc.CallOption) (*UpdateUserProfileRes, error) {
	out := new(UpdateUserProfileRes)
	err := c.cc.Invoke(ctx, "/auth_svc.UserService/UpdateUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUserProfile(ctx context.Context, in *DeleteUserProfileReq, opts ...grpc.CallOption) (*DeleteUserProfileRes, error) {
	out := new(DeleteUserProfileRes)
	err := c.cc.Invoke(ctx, "/auth_svc.UserService/DeleteUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAllTopics(ctx context.Context, in *GetTopicsRequests, opts ...grpc.CallOption) (*GetTopicsResponse, error) {
	out := new(GetTopicsResponse)
	err := c.cc.Invoke(ctx, "/auth_svc.UserService/GetAllTopics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAllCategories(ctx context.Context, in *GetAllCategoriesReq, opts ...grpc.CallOption) (*GetAllCategoriesRes, error) {
	out := new(GetAllCategoriesRes)
	err := c.cc.Invoke(ctx, "/auth_svc.UserService/GetAllCategories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUsername(ctx context.Context, in *UpdateUsernameReq, opts ...grpc.CallOption) (*UpdateUsernameRes, error) {
	out := new(UpdateUsernameRes)
	err := c.cc.Invoke(ctx, "/auth_svc.UserService/UpdateUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	GetUserActivities(context.Context, *UserActivityReq) (*UserActivityRes, error)
	GetUserProfile(context.Context, *UserProfileReq) (*UserProfileRes, error)
	UpdateUserProfile(context.Context, *UpdateUserProfileReq) (*UpdateUserProfileRes, error)
	DeleteUserProfile(context.Context, *DeleteUserProfileReq) (*DeleteUserProfileRes, error)
	GetAllTopics(context.Context, *GetTopicsRequests) (*GetTopicsResponse, error)
	GetAllCategories(context.Context, *GetAllCategoriesReq) (*GetAllCategoriesRes, error)
	UpdateUsername(context.Context, *UpdateUsernameReq) (*UpdateUsernameRes, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) GetUserActivities(context.Context, *UserActivityReq) (*UserActivityRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserActivities not implemented")
}
func (UnimplementedUserServiceServer) GetUserProfile(context.Context, *UserProfileReq) (*UserProfileRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProfile not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserProfile(context.Context, *UpdateUserProfileReq) (*UpdateUserProfileRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserProfile not implemented")
}
func (UnimplementedUserServiceServer) DeleteUserProfile(context.Context, *DeleteUserProfileReq) (*DeleteUserProfileRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserProfile not implemented")
}
func (UnimplementedUserServiceServer) GetAllTopics(context.Context, *GetTopicsRequests) (*GetTopicsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTopics not implemented")
}
func (UnimplementedUserServiceServer) GetAllCategories(context.Context, *GetAllCategoriesReq) (*GetAllCategoriesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCategories not implemented")
}
func (UnimplementedUserServiceServer) UpdateUsername(context.Context, *UpdateUsernameReq) (*UpdateUsernameRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUsername not implemented")
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

func _UserService_GetUserActivities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserActivityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserActivities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_svc.UserService/GetUserActivities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserActivities(ctx, req.(*UserActivityReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserProfileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_svc.UserService/GetUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserProfile(ctx, req.(*UserProfileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserProfileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_svc.UserService/UpdateUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserProfile(ctx, req.(*UpdateUserProfileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserProfileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_svc.UserService/DeleteUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUserProfile(ctx, req.(*DeleteUserProfileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetAllTopics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopicsRequests)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAllTopics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_svc.UserService/GetAllTopics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAllTopics(ctx, req.(*GetTopicsRequests))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetAllCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllCategoriesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAllCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_svc.UserService/GetAllCategories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAllCategories(ctx, req.(*GetAllCategoriesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUsernameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_svc.UserService/UpdateUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUsername(ctx, req.(*UpdateUsernameReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth_svc.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserActivities",
			Handler:    _UserService_GetUserActivities_Handler,
		},
		{
			MethodName: "GetUserProfile",
			Handler:    _UserService_GetUserProfile_Handler,
		},
		{
			MethodName: "UpdateUserProfile",
			Handler:    _UserService_UpdateUserProfile_Handler,
		},
		{
			MethodName: "DeleteUserProfile",
			Handler:    _UserService_DeleteUserProfile_Handler,
		},
		{
			MethodName: "GetAllTopics",
			Handler:    _UserService_GetAllTopics_Handler,
		},
		{
			MethodName: "GetAllCategories",
			Handler:    _UserService_GetAllCategories_Handler,
		},
		{
			MethodName: "UpdateUsername",
			Handler:    _UserService_UpdateUsername_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apis/serviceconn/gateway_user/pb/gw_user.proto",
}
