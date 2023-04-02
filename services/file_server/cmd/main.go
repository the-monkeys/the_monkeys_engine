package main

import (
	"net"

	"github.com/89minutes/the_new_project/services/file_server/config"
	"github.com/89minutes/the_new_project/services/file_server/constant"
	"github.com/89minutes/the_new_project/services/file_server/service/pb"
	"github.com/89minutes/the_new_project/services/file_server/service/server"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadUserConfig()
	if err != nil {
		logrus.Errorf("failed to load user config, error: %+v", err)
	}
	log := logrus.New()

	lis, err := net.Listen("tcp", cfg.FileService)
	if err != nil {
		log.Errorf("failed to listen at port %v, error: %+v", cfg.FileService, err)
	}

	fileService := server.NewFileService(constant.BLOG_FILES)

	grpcServer := grpc.NewServer()

	pb.RegisterUploadBlogFileServer(grpcServer, fileService)

	log.Infof("the user service started at: %v", cfg.FileService)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
