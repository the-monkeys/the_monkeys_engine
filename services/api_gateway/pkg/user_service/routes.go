package user_service

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/the-monkeys/the_monkeys/services/api_gateway/config"
	"github.com/the-monkeys/the_monkeys/services/api_gateway/errors"
	"github.com/the-monkeys/the_monkeys/services/api_gateway/pkg/auth"
	"github.com/the-monkeys/the_monkeys/services/api_gateway/pkg/user_service/pb"
	"google.golang.org/grpc"
)

type UserServiceClient struct {
	Client pb.UserServiceClient
}

func NewUserServiceClient(cfg *config.Address) pb.UserServiceClient {
	cc, err := grpc.Dial(cfg.UserService, grpc.WithInsecure())
	if err != nil {
		logrus.Errorf("cannot dial to grpc user server: %v", err)
	}
	logrus.Infof("The Gateway is dialing to user gRPC server at: %v", cfg.UserService)
	return pb.NewUserServiceClient(cc)
}

func RegisterUserRouter(router *gin.Engine, cfg *config.Address, authClient *auth.ServiceClient) *UserServiceClient {
	mware := auth.InitAuthMiddleware(authClient)

	usc := &UserServiceClient{
		Client: NewUserServiceClient(cfg),
	}
	routes := router.Group("/api/v1/profile")
	routes.Use(mware.AuthRequired)
	routes.GET("/user/:id", usc.GetProfile)
	routes.POST("/user/:id", usc.UpdateProfile)
	routes.POST("/user/deactivate/:id", usc.DeleteMyAccount)

	return usc
}

func (asc *UserServiceClient) GetProfile(ctx *gin.Context) {
	// get id
	id := ctx.Param("id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := asc.Client.GetMyProfile(context.Background(), &pb.GetMyProfileReq{
		Id: userId,
	})

	if err != nil {
		errors.RestError(ctx, err, "user")
		return
	}

	ctx.JSON(http.StatusAccepted, &res)
}

func (asc *UserServiceClient) UpdateProfile(ctx *gin.Context) {
	// get id
	id := ctx.Param("id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	body := UpdateProfile{}
	if err := ctx.BindJSON(&body); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := asc.Client.SetMyProfile(context.Background(), &pb.SetMyProfileReq{
		FirstName:   body.FirstName,
		LastName:    body.LastName,
		CountryCode: body.CountryCode,
		MobileNo:    body.MobileNo,
		About:       body.About,
		Instagram:   body.Instagram,
		Twitter:     body.Twitter,
		Email:       body.Email,
		Id:          userId,
	})

	if err != nil {
		errors.RestError(ctx, err, "user")
		return
	}

	ctx.JSON(http.StatusAccepted, &res)
}

func (asc *UserServiceClient) DeleteMyAccount(ctx *gin.Context) {
	// get id
	id := ctx.Param("id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := asc.Client.DeleteMyProfile(context.Background(), &pb.DeleteMyAccountReq{Id: userId})
	if err != nil {
		logrus.Errorf("cannot connect to user rpc server, error: %v", err)
		errors.RestError(ctx, err, "user_service")
		return
	}

	ctx.JSON(http.StatusOK, res)
}
