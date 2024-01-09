package main

import (
	"log"
	"net"

	isv "github.com/the-monkeys/the_monkeys/apis/interservice/blogs/pb"
	"github.com/the-monkeys/the_monkeys/config"
	"github.com/the-monkeys/the_monkeys/microservices/the_monkeys_blog/blog_service/pb"
	"github.com/the-monkeys/the_monkeys/microservices/the_monkeys_blog/blog_service/service"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.GetConfig()

	if err != nil {
		log.Fatalln("failed to load the config file, error: ", err)
	}

	lis, err := net.Listen("tcp", cfg.Microservices.TheMonkeysBlog)
	if err != nil {
		log.Fatalf("article and service server failed to listen at port %v, error: %v",
			cfg.Microservices.TheMonkeysBlog, err)
	}

	logger := logrus.New()

	osClient, err := service.NewOpenSearchClient(cfg.Opensearch.Address, cfg.Opensearch.Username, cfg.Opensearch.Password, logger)
	if err != nil {
		logger.Fatalf("cannot get the opensearch client, error: %v", err)
	}

	blogService := service.NewBlogService(*osClient, logger)
	interservice := service.NewInterservice(*osClient, logger)

	grpcServer := grpc.NewServer()

	pb.RegisterBlogsAndPostServiceServer(grpcServer, blogService)
	isv.RegisterBlogServiceServer(grpcServer, interservice)

	logrus.Info("starting the blog server at address: ", cfg.Microservices.TheMonkeysBlog)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
