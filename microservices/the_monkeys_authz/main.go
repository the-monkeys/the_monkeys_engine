package main

import (
	"net"

	"github.com/sirupsen/logrus"
	"github.com/the-monkeys/the_monkeys/config"
	"github.com/the-monkeys/the_monkeys/microservices/the_monkeys_authz/pkg/db"
	"github.com/the-monkeys/the_monkeys/microservices/the_monkeys_authz/pkg/pb"
	"github.com/the-monkeys/the_monkeys/microservices/the_monkeys_authz/pkg/services"
	"github.com/the-monkeys/the_monkeys/microservices/the_monkeys_authz/pkg/utils"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		logrus.Fatalf("cannot load auth service config, error: %v", err)
	}

	dbHandler, err := db.NewAuthDBHandler(cfg)
	if err != nil {
		logrus.Fatalf("cannot connect the the db: %+v", err)
	}

	jwt := utils.JwtWrapper{
		SecretKey:       cfg.JWT.SecretKey,
		Issuer:          "go-grpc-auth-svc",
		ExpirationHours: 24 * 365,
	}

	lis, err := net.Listen("tcp", cfg.Microservices.TheMonkeysAuthz)
	if err != nil {
		logrus.Fatalf("auth service cannot listen at address %s, error: %v", cfg.Microservices.TheMonkeysAuthz, err)
	}

	authServer := services.NewAuthServer(dbHandler, jwt, cfg)

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, authServer)

	logrus.Info("starting the authentication server at address: ", cfg.Microservices.TheMonkeysAuthz)
	if err := grpcServer.Serve(lis); err != nil {
		logrus.Fatalf("gRPC auth server cannot start, error: %v", err)
	}
}
