package blogsandposts

import (
	"context"
	"net/http"

	"github.com/89minutes/the_new_project/services/api_gateway/config"
	"github.com/89minutes/the_new_project/services/api_gateway/pkg/auth"
	"github.com/89minutes/the_new_project/services/api_gateway/pkg/blogsandposts/pb"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type BlogServiceClient struct {
	Client pb.BlogsAndPostServiceClient
}

func NewUserServiceClient(cfg *config.Config) pb.BlogsAndPostServiceClient {
	cc, err := grpc.Dial(cfg.BlogAndPostSvcURL, grpc.WithInsecure())
	if err != nil {
		logrus.Errorf("cannot dial to grpc user server: %v", err)
	}

	return pb.NewBlogsAndPostServiceClient(cc)
}

func RegisterBlogRouter(router *gin.Engine, cfg *config.Config, authClient *auth.ServiceClient) *BlogServiceClient {
	mware := auth.InitAuthMiddleware(authClient)

	usc := &BlogServiceClient{
		Client: NewUserServiceClient(cfg),
	}
	routes := router.Group("/api/v1/post")
	routes.Use(mware.AuthRequired)
	routes.POST("/create", usc.CreateABlog)

	return usc
}

func (asc *BlogServiceClient) CreateABlog(ctx *gin.Context) {

	body := CreatePostRequestBody{}
	if err := ctx.BindJSON(&body); err != nil {
		logrus.Errorf("cannot bind json to struct, error: %v", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := asc.Client.CreateABlog(context.Background(), &pb.CreateBlogReq{
		Id:         uuid.NewString(),
		Title:      body.Title,
		Content:    body.Content,
		AuthorName: body.Author,
		AuthorId:   body.AuthorId,
		Published:  body.Published,
		Tags:       body.Tags,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	logrus.Errorf("Response: %+v", res)
	ctx.JSON(http.StatusAccepted, &res)

}