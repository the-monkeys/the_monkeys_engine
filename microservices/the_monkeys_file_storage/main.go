package main

import (
	"net"
	"os"

	"github.com/the-monkeys/the_monkeys/common"
	"github.com/the-monkeys/the_monkeys/config"
	"github.com/the-monkeys/the_monkeys/microservices/the_monkeys_file_storage/constant"
	"github.com/the-monkeys/the_monkeys/microservices/the_monkeys_file_storage/internal/pb"
	"github.com/the-monkeys/the_monkeys/microservices/the_monkeys_file_storage/internal/server"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func init() {
	// Define the complete path including `/` and the folder name
	folderPath := "/" + constant.ProfileDir

	// Check if the directory already exists
	_, err := os.Stat(folderPath)

	// If the directory doesn't exist, create it with permissions 0755
	if os.IsNotExist(err) {
		err = os.MkdirAll(folderPath, 0755)
	}

	if err != nil {
		os.Exit(0)
	}
}

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

	fileService := server.NewFileService(constant.ProfileDir, common.PROFILE_PIC_DIR)
	// newFileServer := server.NewFileServer(common.PROFILE_PIC_DIR, common.BLOG_FILES, log)

	grpcServer := grpc.NewServer()

	pb.RegisterUploadBlogFileServer(grpcServer, fileService)
	// fs.RegisterFileServiceServer(grpcServer, newFileServer)

	log.Infof("✅ the file server started at: %v", cfg.Microservices.TheMonkeysFileStore)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
