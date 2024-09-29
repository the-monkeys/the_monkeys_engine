// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.1
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
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	UserService_GetUserActivities_FullMethodName     = "/auth_svc.UserService/GetUserActivities"
	UserService_GetUserProfile_FullMethodName        = "/auth_svc.UserService/GetUserProfile"
	UserService_UpdateUserProfile_FullMethodName     = "/auth_svc.UserService/UpdateUserProfile"
	UserService_DeleteUserProfile_FullMethodName     = "/auth_svc.UserService/DeleteUserProfile"
	UserService_GetAllTopics_FullMethodName          = "/auth_svc.UserService/GetAllTopics"
	UserService_GetAllCategories_FullMethodName      = "/auth_svc.UserService/GetAllCategories"
	UserService_GetUserDetailsByAccId_FullMethodName = "/auth_svc.UserService/GetUserDetailsByAccId"
	UserService_FollowTopics_FullMethodName          = "/auth_svc.UserService/FollowTopics"
	UserService_UnFollowTopics_FullMethodName        = "/auth_svc.UserService/UnFollowTopics"
	UserService_BookMarkBlog_FullMethodName          = "/auth_svc.UserService/BookMarkBlog"
	UserService_RemoveBookMark_FullMethodName        = "/auth_svc.UserService/RemoveBookMark"
	UserService_InviteCoAuthor_FullMethodName        = "/auth_svc.UserService/InviteCoAuthor"
	UserService_RevokeCoAuthorAccess_FullMethodName  = "/auth_svc.UserService/RevokeCoAuthorAccess"
	UserService_GetBlogsByUserIds_FullMethodName     = "/auth_svc.UserService/GetBlogsByUserIds"
	UserService_CreateNewTopics_FullMethodName       = "/auth_svc.UserService/CreateNewTopics"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	GetUserActivities(ctx context.Context, in *UserActivityReq, opts ...grpc.CallOption) (*UserActivityResp, error)
	GetUserProfile(ctx context.Context, in *UserProfileReq, opts ...grpc.CallOption) (*UserProfileRes, error)
	UpdateUserProfile(ctx context.Context, in *UpdateUserProfileReq, opts ...grpc.CallOption) (*UpdateUserProfileRes, error)
	DeleteUserProfile(ctx context.Context, in *DeleteUserProfileReq, opts ...grpc.CallOption) (*DeleteUserProfileRes, error)
	GetAllTopics(ctx context.Context, in *GetTopicsRequests, opts ...grpc.CallOption) (*GetTopicsResponse, error)
	GetAllCategories(ctx context.Context, in *GetAllCategoriesReq, opts ...grpc.CallOption) (*GetAllCategoriesRes, error)
	GetUserDetailsByAccId(ctx context.Context, in *UserDetailsByAccIdReq, opts ...grpc.CallOption) (*UserDetailsByAccIdResp, error)
	FollowTopics(ctx context.Context, in *TopicActionReq, opts ...grpc.CallOption) (*TopicActionRes, error)
	UnFollowTopics(ctx context.Context, in *TopicActionReq, opts ...grpc.CallOption) (*TopicActionRes, error)
	// Bookmark blog
	BookMarkBlog(ctx context.Context, in *BookMarkReq, opts ...grpc.CallOption) (*BookMarkRes, error)
	// Remove Bookmark
	RemoveBookMark(ctx context.Context, in *BookMarkReq, opts ...grpc.CallOption) (*BookMarkRes, error)
	// Invite a co author
	InviteCoAuthor(ctx context.Context, in *CoAuthorAccessReq, opts ...grpc.CallOption) (*CoAuthorAccessRes, error)
	// Accept co author invitation
	// Reject co author invitation
	// Revoke co author invitation access
	RevokeCoAuthorAccess(ctx context.Context, in *CoAuthorAccessReq, opts ...grpc.CallOption) (*CoAuthorAccessRes, error)
	GetBlogsByUserIds(ctx context.Context, in *BlogsByUserIdsReq, opts ...grpc.CallOption) (*BlogsByUserNameRes, error)
	CreateNewTopics(ctx context.Context, in *CreateTopicsReq, opts ...grpc.CallOption) (*CreateTopicsRes, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUserActivities(ctx context.Context, in *UserActivityReq, opts ...grpc.CallOption) (*UserActivityResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserActivityResp)
	err := c.cc.Invoke(ctx, UserService_GetUserActivities_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserProfile(ctx context.Context, in *UserProfileReq, opts ...grpc.CallOption) (*UserProfileRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserProfileRes)
	err := c.cc.Invoke(ctx, UserService_GetUserProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserProfile(ctx context.Context, in *UpdateUserProfileReq, opts ...grpc.CallOption) (*UpdateUserProfileRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateUserProfileRes)
	err := c.cc.Invoke(ctx, UserService_UpdateUserProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUserProfile(ctx context.Context, in *DeleteUserProfileReq, opts ...grpc.CallOption) (*DeleteUserProfileRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteUserProfileRes)
	err := c.cc.Invoke(ctx, UserService_DeleteUserProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAllTopics(ctx context.Context, in *GetTopicsRequests, opts ...grpc.CallOption) (*GetTopicsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTopicsResponse)
	err := c.cc.Invoke(ctx, UserService_GetAllTopics_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAllCategories(ctx context.Context, in *GetAllCategoriesReq, opts ...grpc.CallOption) (*GetAllCategoriesRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllCategoriesRes)
	err := c.cc.Invoke(ctx, UserService_GetAllCategories_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserDetailsByAccId(ctx context.Context, in *UserDetailsByAccIdReq, opts ...grpc.CallOption) (*UserDetailsByAccIdResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserDetailsByAccIdResp)
	err := c.cc.Invoke(ctx, UserService_GetUserDetailsByAccId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) FollowTopics(ctx context.Context, in *TopicActionReq, opts ...grpc.CallOption) (*TopicActionRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TopicActionRes)
	err := c.cc.Invoke(ctx, UserService_FollowTopics_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UnFollowTopics(ctx context.Context, in *TopicActionReq, opts ...grpc.CallOption) (*TopicActionRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TopicActionRes)
	err := c.cc.Invoke(ctx, UserService_UnFollowTopics_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) BookMarkBlog(ctx context.Context, in *BookMarkReq, opts ...grpc.CallOption) (*BookMarkRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BookMarkRes)
	err := c.cc.Invoke(ctx, UserService_BookMarkBlog_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) RemoveBookMark(ctx context.Context, in *BookMarkReq, opts ...grpc.CallOption) (*BookMarkRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BookMarkRes)
	err := c.cc.Invoke(ctx, UserService_RemoveBookMark_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) InviteCoAuthor(ctx context.Context, in *CoAuthorAccessReq, opts ...grpc.CallOption) (*CoAuthorAccessRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CoAuthorAccessRes)
	err := c.cc.Invoke(ctx, UserService_InviteCoAuthor_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) RevokeCoAuthorAccess(ctx context.Context, in *CoAuthorAccessReq, opts ...grpc.CallOption) (*CoAuthorAccessRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CoAuthorAccessRes)
	err := c.cc.Invoke(ctx, UserService_RevokeCoAuthorAccess_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetBlogsByUserIds(ctx context.Context, in *BlogsByUserIdsReq, opts ...grpc.CallOption) (*BlogsByUserNameRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BlogsByUserNameRes)
	err := c.cc.Invoke(ctx, UserService_GetBlogsByUserIds_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateNewTopics(ctx context.Context, in *CreateTopicsReq, opts ...grpc.CallOption) (*CreateTopicsRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTopicsRes)
	err := c.cc.Invoke(ctx, UserService_CreateNewTopics_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility.
type UserServiceServer interface {
	GetUserActivities(context.Context, *UserActivityReq) (*UserActivityResp, error)
	GetUserProfile(context.Context, *UserProfileReq) (*UserProfileRes, error)
	UpdateUserProfile(context.Context, *UpdateUserProfileReq) (*UpdateUserProfileRes, error)
	DeleteUserProfile(context.Context, *DeleteUserProfileReq) (*DeleteUserProfileRes, error)
	GetAllTopics(context.Context, *GetTopicsRequests) (*GetTopicsResponse, error)
	GetAllCategories(context.Context, *GetAllCategoriesReq) (*GetAllCategoriesRes, error)
	GetUserDetailsByAccId(context.Context, *UserDetailsByAccIdReq) (*UserDetailsByAccIdResp, error)
	FollowTopics(context.Context, *TopicActionReq) (*TopicActionRes, error)
	UnFollowTopics(context.Context, *TopicActionReq) (*TopicActionRes, error)
	// Bookmark blog
	BookMarkBlog(context.Context, *BookMarkReq) (*BookMarkRes, error)
	// Remove Bookmark
	RemoveBookMark(context.Context, *BookMarkReq) (*BookMarkRes, error)
	// Invite a co author
	InviteCoAuthor(context.Context, *CoAuthorAccessReq) (*CoAuthorAccessRes, error)
	// Accept co author invitation
	// Reject co author invitation
	// Revoke co author invitation access
	RevokeCoAuthorAccess(context.Context, *CoAuthorAccessReq) (*CoAuthorAccessRes, error)
	GetBlogsByUserIds(context.Context, *BlogsByUserIdsReq) (*BlogsByUserNameRes, error)
	CreateNewTopics(context.Context, *CreateTopicsReq) (*CreateTopicsRes, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserServiceServer struct{}

func (UnimplementedUserServiceServer) GetUserActivities(context.Context, *UserActivityReq) (*UserActivityResp, error) {
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
func (UnimplementedUserServiceServer) GetUserDetailsByAccId(context.Context, *UserDetailsByAccIdReq) (*UserDetailsByAccIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserDetailsByAccId not implemented")
}
func (UnimplementedUserServiceServer) FollowTopics(context.Context, *TopicActionReq) (*TopicActionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowTopics not implemented")
}
func (UnimplementedUserServiceServer) UnFollowTopics(context.Context, *TopicActionReq) (*TopicActionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnFollowTopics not implemented")
}
func (UnimplementedUserServiceServer) BookMarkBlog(context.Context, *BookMarkReq) (*BookMarkRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BookMarkBlog not implemented")
}
func (UnimplementedUserServiceServer) RemoveBookMark(context.Context, *BookMarkReq) (*BookMarkRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveBookMark not implemented")
}
func (UnimplementedUserServiceServer) InviteCoAuthor(context.Context, *CoAuthorAccessReq) (*CoAuthorAccessRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InviteCoAuthor not implemented")
}
func (UnimplementedUserServiceServer) RevokeCoAuthorAccess(context.Context, *CoAuthorAccessReq) (*CoAuthorAccessRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeCoAuthorAccess not implemented")
}
func (UnimplementedUserServiceServer) GetBlogsByUserIds(context.Context, *BlogsByUserIdsReq) (*BlogsByUserNameRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlogsByUserIds not implemented")
}
func (UnimplementedUserServiceServer) CreateNewTopics(context.Context, *CreateTopicsReq) (*CreateTopicsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewTopics not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}
func (UnimplementedUserServiceServer) testEmbeddedByValue()                     {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
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
		FullMethod: UserService_GetUserActivities_FullMethodName,
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
		FullMethod: UserService_GetUserProfile_FullMethodName,
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
		FullMethod: UserService_UpdateUserProfile_FullMethodName,
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
		FullMethod: UserService_DeleteUserProfile_FullMethodName,
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
		FullMethod: UserService_GetAllTopics_FullMethodName,
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
		FullMethod: UserService_GetAllCategories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAllCategories(ctx, req.(*GetAllCategoriesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserDetailsByAccId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDetailsByAccIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserDetailsByAccId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserDetailsByAccId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserDetailsByAccId(ctx, req.(*UserDetailsByAccIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_FollowTopics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopicActionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FollowTopics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_FollowTopics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FollowTopics(ctx, req.(*TopicActionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UnFollowTopics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopicActionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UnFollowTopics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UnFollowTopics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UnFollowTopics(ctx, req.(*TopicActionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_BookMarkBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookMarkReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).BookMarkBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_BookMarkBlog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).BookMarkBlog(ctx, req.(*BookMarkReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_RemoveBookMark_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookMarkReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RemoveBookMark(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_RemoveBookMark_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RemoveBookMark(ctx, req.(*BookMarkReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_InviteCoAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CoAuthorAccessReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).InviteCoAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_InviteCoAuthor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).InviteCoAuthor(ctx, req.(*CoAuthorAccessReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_RevokeCoAuthorAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CoAuthorAccessReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RevokeCoAuthorAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_RevokeCoAuthorAccess_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RevokeCoAuthorAccess(ctx, req.(*CoAuthorAccessReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetBlogsByUserIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlogsByUserIdsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetBlogsByUserIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetBlogsByUserIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetBlogsByUserIds(ctx, req.(*BlogsByUserIdsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateNewTopics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTopicsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateNewTopics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CreateNewTopics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateNewTopics(ctx, req.(*CreateTopicsReq))
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
			MethodName: "GetUserDetailsByAccId",
			Handler:    _UserService_GetUserDetailsByAccId_Handler,
		},
		{
			MethodName: "FollowTopics",
			Handler:    _UserService_FollowTopics_Handler,
		},
		{
			MethodName: "UnFollowTopics",
			Handler:    _UserService_UnFollowTopics_Handler,
		},
		{
			MethodName: "BookMarkBlog",
			Handler:    _UserService_BookMarkBlog_Handler,
		},
		{
			MethodName: "RemoveBookMark",
			Handler:    _UserService_RemoveBookMark_Handler,
		},
		{
			MethodName: "InviteCoAuthor",
			Handler:    _UserService_InviteCoAuthor_Handler,
		},
		{
			MethodName: "RevokeCoAuthorAccess",
			Handler:    _UserService_RevokeCoAuthorAccess_Handler,
		},
		{
			MethodName: "GetBlogsByUserIds",
			Handler:    _UserService_GetBlogsByUserIds_Handler,
		},
		{
			MethodName: "CreateNewTopics",
			Handler:    _UserService_CreateNewTopics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apis/serviceconn/gateway_user/pb/gw_user.proto",
}
