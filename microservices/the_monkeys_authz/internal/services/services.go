package services

import (
	"context"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/the-monkeys/the_monkeys/apis/serviceconn/gateway_authz/pb"
	"github.com/the-monkeys/the_monkeys/common"
	"github.com/the-monkeys/the_monkeys/config"
	"github.com/the-monkeys/the_monkeys/microservices/the_monkeys_authz/internal/db"
	"github.com/the-monkeys/the_monkeys/microservices/the_monkeys_authz/internal/models"

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

func (as *AuthzSvc) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	logrus.Infof("got the request data: %+v", req)

	claims, err := as.jwt.ValidateToken(req.Token)
	if err != nil {
		logrus.Errorf("cannot validate the json token, error: %v", err)
		return &pb.ValidateResponse{
			StatusCode: http.StatusBadRequest,
			Error: &pb.Error{
				Status:  http.StatusUnauthorized,
				Error:   common.ErrTokenIsNotValid.Error(),
				Message: common.ErrTokenIsNotValid.Error(),
			},
		}, nil
	}

	var user models.TheMonkeysUser
	// Check if the email exists
	if err := as.dbCli.PsqlClient.QueryRow("SELECT email, password FROM the_monkeys_user WHERE email=$1;", claims.Email).
		Scan(&user.Email, &user.Password); err != nil {
		logrus.Errorf("cannot validate token as the email %s doesn't exist, error: %+v", claims.Email, err)
		return &pb.ValidateResponse{
			StatusCode: http.StatusNotFound,
			Error: &pb.Error{
				Status:  http.StatusUnauthorized,
				Error:   common.ErrTokenIsNotValid.Error(),
				Message: common.ErrTokenIsNotValid.Error(),
			},
		}, nil
	}

	return &pb.ValidateResponse{
		StatusCode: http.StatusOK,
		UserId:     user.Id,
		Email:      claims.Email,
	}, nil
}
