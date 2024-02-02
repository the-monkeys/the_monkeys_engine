package main

import (
	"net"

	"github.com/sirupsen/logrus"
	"github.com/the-monkeys/the_monkeys/apis/serviceconn/gateway_user/pb"
	"github.com/the-monkeys/the_monkeys/config"
	"github.com/the-monkeys/the_monkeys/microservices/the_monkeys_users/internal/server"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		logrus.Errorf("failed to load user config, error: %+v", err)
	}
	log := logrus.New()

	// db := database.NewUserDbHandler(cfg, log)

	lis, err := net.Listen("tcp", cfg.Microservices.TheMonkeysUser)
	if err != nil {
		log.Errorf("failed to listen at port %v, error: %+v", cfg.Microservices.TheMonkeysUser, err)
	}

	// conn, err := grpc.Dial(cfg.Microservices.TheMonkeysBlog, grpc.WithInsecure())
	// if err != nil {
	// 	log.Errorf("failed to dial to blog service at %v, error: %+v", cfg.Microservices.TheMonkeysBlog, err)
	// 	return
	// }

	// userService := server.NewUserService(db, log, isv.NewBlogServiceClient(conn))
	userService := server.NewUserSvc()

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, userService)

	log.Infof("✅ the user service started at: %v", cfg.Microservices.TheMonkeysUser)
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
