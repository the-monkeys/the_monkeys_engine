package file_server

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/89minutes/the_new_project/services/api_gateway/config"
	"github.com/89minutes/the_new_project/services/api_gateway/errors"
	"github.com/89minutes/the_new_project/services/api_gateway/pkg/auth"
	"github.com/89minutes/the_new_project/services/api_gateway/pkg/file_server/pb"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type FileServiceClient struct {
	Client pb.UploadBlogFileClient
}

func NewFileServiceClient(cfg *config.Address) pb.UploadBlogFileClient {
	cc, err := grpc.Dial(cfg.FileService, grpc.WithInsecure())
	if err != nil {
		logrus.Errorf("cannot dial to grpc file server: %v", err)
	}
	logrus.Infof("The Gateway is dialing to file gRPC server at: %v", cfg.FileService)
	return pb.NewUploadBlogFileClient(cc)
}

func RegisterUserRouter(router *gin.Engine, cfg *config.Address, authClient *auth.ServiceClient) *FileServiceClient {
	mware := auth.InitAuthMiddleware(authClient)

	usc := &FileServiceClient{
		Client: NewFileServiceClient(cfg),
	}
	routes := router.Group("/api/v1/files")

	routes.GET("/post/:id/:fileName", usc.GetBlogFile)

	// route defined to get profile pic

	routes.Use(mware.AuthRequired)
	routes.POST("/post/:id", usc.UploadBlogFile)
	routes.DELETE("/post/:id/:fileName", usc.DeleteBlogFile)

	// route defined to access profile
	routes.POST("/profile/:user_id/profile", usc.UploadProfilePic)

	return usc
}

func (asc *FileServiceClient) UploadBlogFile(ctx *gin.Context) {
	// get Id of the blog from the URL
	blogId := ctx.Param("id")

	// Get file from the form file section
	file, fileHeader, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	defer file.Close()

	// Read the file and make it slice of bytes
	imageData, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading image data:", err)
	}

	stream, err := asc.Client.UploadBlogFile(context.Background())
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	chunk := &pb.UploadBlogFileReq{
		BlogId:   blogId,
		Data:     imageData,
		FileName: fileHeader.Filename,
	}
	err = stream.Send(chunk)
	if err != nil {
		log.Fatal("cannot send file info to server: ", err, stream.RecvMsg(nil))
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// log.Printf("%+v\n", response)
	ctx.JSON(http.StatusAccepted, resp)
}

func (asc *FileServiceClient) GetBlogFile(ctx *gin.Context) {
	blogId := ctx.Param("id")
	fileName := ctx.Param("fileName")

	stream, err := asc.Client.GetBlogFile(context.Background(), &pb.GetBlogFileReq{
		BlogId:   blogId,
		FileName: fileName,
	})
	if err != nil {
		logrus.Errorf("cannot connect to user rpc server, error: %v", err)
		_ = ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	resp, err := stream.Recv()
	if err == io.EOF {

	}
	if err != nil {
		errors.RestError(ctx, err, "file_service")
		logrus.Errorf("cannot get the stream data, error: %+v", err)
		return
	}

	// TODO: Remove the comment lines
	// ctx.Header("Content-Disposition", "attachment; filename=file-name.txt")
	// ctx.Data(http.StatusOK, "application/octet-stream", resp.Data)

	// ctx.JSON(http.StatusAccepted, "uploaded")
	ctx.Writer.Write(resp.Data)
}

func (asc *FileServiceClient) DeleteBlogFile(ctx *gin.Context) {
	blogId := ctx.Param("id")
	fileName := ctx.Param("fileName")

	res, err := asc.Client.DeleteBlogFile(context.Background(), &pb.DeleteBlogFileReq{
		BlogId:   blogId,
		FileName: fileName,
	})

	if err != nil {
		logrus.Errorf("cannot connect to user rpc server, error: %v", err)
		_ = ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusAccepted, res)
}

func (asc *FileServiceClient) UploadProfilePic(ctx *gin.Context) {
	// get Id of the blog from the URL
	userId := ctx.Param("user_id")

	// Get file from the form file section
	file, fileHeader, err := ctx.Request.FormFile("profile_pic")
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	defer file.Close()

	// Read the file and make it slice of bytes
	imageData, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading image data:", err)
	}

	stream, err := asc.Client.UploadProfilePic(context.Background())
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	chunk := &pb.UploadProfilePicReq{
		UserId:   userId,
		Data:     imageData,
		FileType: fileHeader.Filename,
	}
	err = stream.Send(chunk)
	if err != nil {
		log.Fatal("cannot send file info to server: ", err, stream.RecvMsg(nil))
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// log.Printf("%+v\n", response)
	ctx.JSON(http.StatusAccepted, resp)
}
