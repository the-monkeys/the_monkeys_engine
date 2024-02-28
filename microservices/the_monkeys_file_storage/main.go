package main

import (
	"net"

	"github.com/the-monkeys/the_monkeys/common"
	"github.com/the-monkeys/the_monkeys/config"
	"github.com/the-monkeys/the_monkeys/microservices/the_monkeys_file_storage/constant"
	"github.com/the-monkeys/the_monkeys/microservices/the_monkeys_file_storage/internal/pb"
	"github.com/the-monkeys/the_monkeys/microservices/the_monkeys_file_storage/internal/server"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		logrus.Errorf("Failed to load file server config, error: %+v", err)
	}

	log := logrus.New()

	lis, err := net.Listen("tcp", cfg.Microservices.TheMonkeysFileStore)
	if err != nil {
		log.Errorf("File server failed to listen at port %v, error: %+v", cfg.Microservices.TheMonkeysFileStore, err)
	}

	fileService := server.NewFileService(constant.BLOG_FILES, common.PROFILE_PIC_DIR)
	// newFileServer := server.NewFileServer(common.PROFILE_PIC_DIR, common.BLOG_FILES, log)

	grpcServer := grpc.NewServer()

	pb.RegisterUploadBlogFileServer(grpcServer, fileService)
	// fs.RegisterFileServiceServer(grpcServer, newFileServer)

	log.Infof("The file server started at: %v", cfg.Microservices.TheMonkeysFileStore)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
