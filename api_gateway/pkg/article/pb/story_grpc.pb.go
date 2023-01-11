// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: api_gateway/pkg/article/pb/story.proto

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

// ArticleServiceClient is the client API for ArticleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArticleServiceClient interface {
	CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*CreateArticleResponse, error)
	GetArticles(ctx context.Context, in *GetArticlesRequest, opts ...grpc.CallOption) (ArticleService_GetArticlesClient, error)
	GetArticleById(ctx context.Context, in *GetArticleByIdReq, opts ...grpc.CallOption) (*GetArticleByIdResp, error)
}

type articleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewArticleServiceClient(cc grpc.ClientConnInterface) ArticleServiceClient {
	return &articleServiceClient{cc}
}

func (c *articleServiceClient) CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*CreateArticleResponse, error) {
	out := new(CreateArticleResponse)
	err := c.cc.Invoke(ctx, "/auth.ArticleService/CreateArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) GetArticles(ctx context.Context, in *GetArticlesRequest, opts ...grpc.CallOption) (ArticleService_GetArticlesClient, error) {
	stream, err := c.cc.NewStream(ctx, &ArticleService_ServiceDesc.Streams[0], "/auth.ArticleService/GetArticles", opts...)
	if err != nil {
		return nil, err
	}
	x := &articleServiceGetArticlesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ArticleService_GetArticlesClient interface {
	Recv() (*GetArticlesResponse, error)
	grpc.ClientStream
}

type articleServiceGetArticlesClient struct {
	grpc.ClientStream
}

func (x *articleServiceGetArticlesClient) Recv() (*GetArticlesResponse, error) {
	m := new(GetArticlesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *articleServiceClient) GetArticleById(ctx context.Context, in *GetArticleByIdReq, opts ...grpc.CallOption) (*GetArticleByIdResp, error) {
	out := new(GetArticleByIdResp)
	err := c.cc.Invoke(ctx, "/auth.ArticleService/GetArticleById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArticleServiceServer is the server API for ArticleService service.
// All implementations must embed UnimplementedArticleServiceServer
// for forward compatibility
type ArticleServiceServer interface {
	CreateArticle(context.Context, *CreateArticleRequest) (*CreateArticleResponse, error)
	GetArticles(*GetArticlesRequest, ArticleService_GetArticlesServer) error
	GetArticleById(context.Context, *GetArticleByIdReq) (*GetArticleByIdResp, error)
	mustEmbedUnimplementedArticleServiceServer()
}

// UnimplementedArticleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedArticleServiceServer struct {
}

func (UnimplementedArticleServiceServer) CreateArticle(context.Context, *CreateArticleRequest) (*CreateArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArticle not implemented")
}
func (UnimplementedArticleServiceServer) GetArticles(*GetArticlesRequest, ArticleService_GetArticlesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetArticles not implemented")
}
func (UnimplementedArticleServiceServer) GetArticleById(context.Context, *GetArticleByIdReq) (*GetArticleByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticleById not implemented")
}
func (UnimplementedArticleServiceServer) mustEmbedUnimplementedArticleServiceServer() {}

// UnsafeArticleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArticleServiceServer will
// result in compilation errors.
type UnsafeArticleServiceServer interface {
	mustEmbedUnimplementedArticleServiceServer()
}

func RegisterArticleServiceServer(s grpc.ServiceRegistrar, srv ArticleServiceServer) {
	s.RegisterService(&ArticleService_ServiceDesc, srv)
}

func _ArticleService_CreateArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).CreateArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.ArticleService/CreateArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).CreateArticle(ctx, req.(*CreateArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_GetArticles_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetArticlesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ArticleServiceServer).GetArticles(m, &articleServiceGetArticlesServer{stream})
}

type ArticleService_GetArticlesServer interface {
	Send(*GetArticlesResponse) error
	grpc.ServerStream
}

type articleServiceGetArticlesServer struct {
	grpc.ServerStream
}

func (x *articleServiceGetArticlesServer) Send(m *GetArticlesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ArticleService_GetArticleById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).GetArticleById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.ArticleService/GetArticleById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).GetArticleById(ctx, req.(*GetArticleByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ArticleService_ServiceDesc is the grpc.ServiceDesc for ArticleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArticleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.ArticleService",
	HandlerType: (*ArticleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateArticle",
			Handler:    _ArticleService_CreateArticle_Handler,
		},
		{
			MethodName: "GetArticleById",
			Handler:    _ArticleService_GetArticleById_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetArticles",
			Handler:       _ArticleService_GetArticles_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api_gateway/pkg/article/pb/story.proto",
}
