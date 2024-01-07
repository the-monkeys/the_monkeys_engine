package main

import (
	"net"

	"github.com/sirupsen/logrus"
	"github.com/the-monkeys/the_monkeys/microservices/auth_service/pkg/config"
	"github.com/the-monkeys/the_monkeys/microservices/auth_service/pkg/db"
	"github.com/the-monkeys/the_monkeys/microservices/auth_service/pkg/pb"
	"github.com/the-monkeys/the_monkeys/microservices/auth_service/pkg/services"
	"github.com/the-monkeys/the_monkeys/microservices/auth_service/pkg/utils"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		logrus.Fatalf("cannot load auth service config, error: %v", err)
	}

	dbHandler, err := db.NewAuthDBHandler(cfg.DBUrl)
	if err != nil {
		logrus.Fatalf("cannot connect the the db: %+v", err)
	}

	jwt := utils.JwtWrapper{
		SecretKey:       cfg.JWTSecretKey,
		Issuer:          "go-grpc-auth-svc",
		ExpirationHours: 24 * 365,
	}

	lis, err := net.Listen("tcp", cfg.AuthAddr)
	if err != nil {
		logrus.Fatalf("auth service cannot listen at address %s, error: %v", cfg.AuthAddr, err)
	}

	authServer := services.NewAuthServer(dbHandler, jwt, cfg)

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, authServer)

	logrus.Info("starting the authentication server at address: ", cfg.AuthAddr)
	if err := grpcServer.Serve(lis); err != nil {
		logrus.Fatalf("gRPC auth server cannot start, error: %v", err)
	}
}
