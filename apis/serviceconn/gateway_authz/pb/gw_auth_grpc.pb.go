// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: apis/serviceconn/gateway_authz/pb/gw_auth.proto

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

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*RegisterUserResponse, error)
	Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error)
	Login(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error)
	ForgotPassword(ctx context.Context, in *ForgotPasswordReq, opts ...grpc.CallOption) (*ForgotPasswordRes, error)
	ResetPassword(ctx context.Context, in *ResetPasswordReq, opts ...grpc.CallOption) (*ResetPasswordRes, error)
	UpdatePassword(ctx context.Context, in *UpdatePasswordReq, opts ...grpc.CallOption) (*UpdatePasswordRes, error)
	RequestForEmailVerification(ctx context.Context, in *EmailVerificationReq, opts ...grpc.CallOption) (*EmailVerificationRes, error)
	VerifyEmail(ctx context.Context, in *VerifyEmailReq, opts ...grpc.CallOption) (*VerifyEmailRes, error)
	UpdateUsername(ctx context.Context, in *UpdateUsernameReq, opts ...grpc.CallOption) (*UpdateUsernameRes, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*RegisterUserResponse, error) {
	out := new(RegisterUserResponse)
	err := c.cc.Invoke(ctx, "/auth_svc.AuthService/RegisterUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error) {
	out := new(ValidateResponse)
	err := c.cc.Invoke(ctx, "/auth_svc.AuthService/Validate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Login(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	out := new(LoginUserResponse)
	err := c.cc.Invoke(ctx, "/auth_svc.AuthService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ForgotPassword(ctx context.Context, in *ForgotPasswordReq, opts ...grpc.CallOption) (*ForgotPasswordRes, error) {
	out := new(ForgotPasswordRes)
	err := c.cc.Invoke(ctx, "/auth_svc.AuthService/ForgotPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ResetPassword(ctx context.Context, in *ResetPasswordReq, opts ...grpc.CallOption) (*ResetPasswordRes, error) {
	out := new(ResetPasswordRes)
	err := c.cc.Invoke(ctx, "/auth_svc.AuthService/ResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UpdatePassword(ctx context.Context, in *UpdatePasswordReq, opts ...grpc.CallOption) (*UpdatePasswordRes, error) {
	out := new(UpdatePasswordRes)
	err := c.cc.Invoke(ctx, "/auth_svc.AuthService/UpdatePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) RequestForEmailVerification(ctx context.Context, in *EmailVerificationReq, opts ...grpc.CallOption) (*EmailVerificationRes, error) {
	out := new(EmailVerificationRes)
	err := c.cc.Invoke(ctx, "/auth_svc.AuthService/RequestForEmailVerification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) VerifyEmail(ctx context.Context, in *VerifyEmailReq, opts ...grpc.CallOption) (*VerifyEmailRes, error) {
	out := new(VerifyEmailRes)
	err := c.cc.Invoke(ctx, "/auth_svc.AuthService/VerifyEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UpdateUsername(ctx context.Context, in *UpdateUsernameReq, opts ...grpc.CallOption) (*UpdateUsernameRes, error) {
	out := new(UpdateUsernameRes)
	err := c.cc.Invoke(ctx, "/auth_svc.AuthService/UpdateUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	RegisterUser(context.Context, *RegisterUserRequest) (*RegisterUserResponse, error)
	Validate(context.Context, *ValidateRequest) (*ValidateResponse, error)
	Login(context.Context, *LoginUserRequest) (*LoginUserResponse, error)
	ForgotPassword(context.Context, *ForgotPasswordReq) (*ForgotPasswordRes, error)
	ResetPassword(context.Context, *ResetPasswordReq) (*ResetPasswordRes, error)
	UpdatePassword(context.Context, *UpdatePasswordReq) (*UpdatePasswordRes, error)
	RequestForEmailVerification(context.Context, *EmailVerificationReq) (*EmailVerificationRes, error)
	VerifyEmail(context.Context, *VerifyEmailReq) (*VerifyEmailRes, error)
	UpdateUsername(context.Context, *UpdateUsernameReq) (*UpdateUsernameRes, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) RegisterUser(context.Context, *RegisterUserRequest) (*RegisterUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedAuthServiceServer) Validate(context.Context, *ValidateRequest) (*ValidateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}
func (UnimplementedAuthServiceServer) Login(context.Context, *LoginUserRequest) (*LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthServiceServer) ForgotPassword(context.Context, *ForgotPasswordReq) (*ForgotPasswordRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForgotPassword not implemented")
}
func (UnimplementedAuthServiceServer) ResetPassword(context.Context, *ResetPasswordReq) (*ResetPasswordRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (UnimplementedAuthServiceServer) UpdatePassword(context.Context, *UpdatePasswordReq) (*UpdatePasswordRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePassword not implemented")
}
func (UnimplementedAuthServiceServer) RequestForEmailVerification(context.Context, *EmailVerificationReq) (*EmailVerificationRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestForEmailVerification not implemented")
}
func (UnimplementedAuthServiceServer) VerifyEmail(context.Context, *VerifyEmailReq) (*VerifyEmailRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyEmail not implemented")
}
func (UnimplementedAuthServiceServer) UpdateUsername(context.Context, *UpdateUsernameReq) (*UpdateUsernameRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUsername not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_svc.AuthService/RegisterUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).RegisterUser(ctx, req.(*RegisterUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_svc.AuthService/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Validate(ctx, req.(*ValidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_svc.AuthService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ForgotPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForgotPasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ForgotPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_svc.AuthService/ForgotPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ForgotPassword(ctx, req.(*ForgotPasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_svc.AuthService/ResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ResetPassword(ctx, req.(*ResetPasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_svc.AuthService/UpdatePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UpdatePassword(ctx, req.(*UpdatePasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_RequestForEmailVerification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailVerificationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).RequestForEmailVerification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_svc.AuthService/RequestForEmailVerification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).RequestForEmailVerification(ctx, req.(*EmailVerificationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_VerifyEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyEmailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).VerifyEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_svc.AuthService/VerifyEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).VerifyEmail(ctx, req.(*VerifyEmailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UpdateUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUsernameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UpdateUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_svc.AuthService/UpdateUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UpdateUsername(ctx, req.(*UpdateUsernameReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth_svc.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterUser",
			Handler:    _AuthService_RegisterUser_Handler,
		},
		{
			MethodName: "Validate",
			Handler:    _AuthService_Validate_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
		{
			MethodName: "ForgotPassword",
			Handler:    _AuthService_ForgotPassword_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _AuthService_ResetPassword_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _AuthService_UpdatePassword_Handler,
		},
		{
			MethodName: "RequestForEmailVerification",
			Handler:    _AuthService_RequestForEmailVerification_Handler,
		},
		{
			MethodName: "VerifyEmail",
			Handler:    _AuthService_VerifyEmail_Handler,
		},
		{
			MethodName: "UpdateUsername",
			Handler:    _AuthService_UpdateUsername_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apis/serviceconn/gateway_authz/pb/gw_auth.proto",
}
