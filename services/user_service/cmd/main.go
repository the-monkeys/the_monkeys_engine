package main

import (
	"net"

	"github.com/sirupsen/logrus"
	isv "github.com/the-monkeys/the_monkeys/apis/interservice/blogs/pb"
	"github.com/the-monkeys/the_monkeys/services/user_service/service/config"
	"github.com/the-monkeys/the_monkeys/services/user_service/service/database"
	"github.com/the-monkeys/the_monkeys/services/user_service/service/pb"
	"github.com/the-monkeys/the_monkeys/services/user_service/service/server"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadUserConfig()
	if err != nil {
		logrus.Errorf("failed to load user config, error: %+v", err)
	}
	log := logrus.New()

	db := database.NewUserDbHandler(cfg.DBUrl, log)

	lis, err := net.Listen("tcp", cfg.UserSrvPort)
	if err != nil {
		log.Errorf("failed to listen at port %v, error: %+v", cfg.UserSrvPort, err)
	}

	conn, err := grpc.Dial(cfg.BlogAndPostSvcURL, grpc.WithInsecure())
	if err != nil {
		log.Errorf("failed to dial to blog service at %v, error: %+v", cfg.BlogAndPostSvcURL, err)
		return
	}

	userService := server.NewUserService(db, log, isv.NewBlogServiceClient(conn))

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, userService)

	log.Infof("the user service started at: %v", cfg.UserSrvPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}

func BlogServiceConn(addr string) (*grpc.ClientConn, error) {
	logrus.Infof("gRPC dialing to the blog server: %v", addr)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	return conn, err
}
