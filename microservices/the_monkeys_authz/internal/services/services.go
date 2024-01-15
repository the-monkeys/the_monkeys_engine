package services

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/the-monkeys/the_monkeys/apis/serviceconn/gateway_authz/pb"
	"github.com/the-monkeys/the_monkeys/config"
	"github.com/the-monkeys/the_monkeys/microservices/the_monkeys_authz/internal/db"

	"github.com/the-monkeys/the_monkeys/microservices/the_monkeys_authz/internal/utils"
)

type AuthzSvc struct {
	dbCli  *db.AuthDBHandler
	jwt    utils.JwtWrapper
	config *config.Config
	pb.UnimplementedAuthServiceServer
}

func NewAuthzSvc(dbCli *db.AuthDBHandler, jwt utils.JwtWrapper, config *config.Config) *AuthzSvc {
	return &AuthzSvc{
		dbCli:  dbCli,
		jwt:    jwt,
		config: config,
	}
}

func (as *AuthzSvc) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	logrus.Infof("got the request data: %+v", req)
	return &pb.RegisterUserResponse{}, nil
}
