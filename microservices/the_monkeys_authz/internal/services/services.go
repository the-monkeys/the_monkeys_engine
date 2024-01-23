package services

import (
	"context"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/the-monkeys/the_monkeys/apis/serviceconn/gateway_authz/pb"
	"github.com/the-monkeys/the_monkeys/common"
	"github.com/the-monkeys/the_monkeys/config"
	"github.com/the-monkeys/the_monkeys/microservices/the_monkeys_authz/internal/db"
	"github.com/the-monkeys/the_monkeys/microservices/the_monkeys_authz/internal/models"

	"github.com/the-monkeys/the_monkeys/microservices/the_monkeys_authz/internal/utils"
)

type AuthzSvc struct {
	dbConn db.AuthDBHandler
	jwt    utils.JwtWrapper
	config *config.Config
	pb.UnimplementedAuthServiceServer
}

func NewAuthzSvc(dbCli db.AuthDBHandler, jwt utils.JwtWrapper, config *config.Config) *AuthzSvc {
	return &AuthzSvc{
		dbConn: dbCli,
		jwt:    jwt,
		config: config,
	}
}

func (as *AuthzSvc) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	logrus.Infof("got the request data for : %+v", req.Email)
	user := &models.TheMonkeysUser{}
	// Check if the user exists with the same email id return conflict
	_, err := as.dbConn.CheckIfEmailExist(req.Email)
	if err == nil {
		logrus.Errorf("cannot register the user, as the email %s is existing already", req.Email)
		return &pb.RegisterUserResponse{
			StatusCode: http.StatusConflict,
			Error: &pb.Error{
				Status:  http.StatusConflict,
				Message: "The email is already registered",
				Error:   "An account is already registered with this email",
			},
		}, nil
	}

	hash := string(utils.GenHash())
	encHash := utils.HashPassword(hash)

	// Create a userId and username
	user.ProfileId = utils.RandomString(12)
	user.Username = utils.RandomString(7)
	user.FirstName = req.FirstName
	user.LastName = req.GetLastName()
	user.Email = req.GetEmail()
	user.Password = utils.HashPassword(req.Password)
	// user.CreateTime = time.Now().Format(common.DATE_TIME_FORMAT)
	// user.UpdateTime = time.Now().Format(common.DATE_TIME_FORMAT)
	user.IsActive = true
	user.EmailVerificationToken = encHash
	user.EmailVerificationTimeout = time.Now().Add(time.Hour * 24)
	user.Deactivated = false
	user.LoginMethod = req.LoginMethod.String()

	logrus.Infof("registering the user with email %v", req.Email)
	userId, err := as.dbConn.RegisterUser(user)
	if err != nil {
		return nil, err
	}

	// Send email verification mail as a routine else the register api gets slower
	emailBody := utils.EmailVerificationHTML(user.Email, hash)
	go as.SendMail(user.Email, emailBody)

	logrus.Infof("user %s is successfully registered.", user.Email)

	// Generate and return token
	token, err := as.jwt.GenerateToken(user)
	if err != nil {
		logrus.Errorf("cannot create a token for %s, error: %+v", req.Email, err)
		return &pb.RegisterUserResponse{
			StatusCode: http.StatusBadRequest,
			Error: &pb.Error{
				Status:  http.StatusInternalServerError,
				Message: "You are successfully registered",
				Error:   "Try to login",
			},
		}, nil
	}

	return &pb.RegisterUserResponse{
		StatusCode:    http.StatusCreated,
		Token:         token,
		EmailVerified: false,
		UserName:      user.Username,
		Email:         user.Email,
		UserId:        userId,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
	}, nil
}

// Validate user runs to check to validate the user. It checks
// If the token is correct
// If the token is expired
// Is the token belongs to the user
// Is the user existing in the db or an active user
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

	// Check if the email exists
	as.dbConn.CheckIfEmailExist(claims.Email)
	user, err := as.dbConn.CheckIfEmailExist(claims.Email)
	if err != nil {
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
		UserName:   user.Username,
	}, nil
}

func (as *AuthzSvc) Login(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	logrus.Infof("user has requested to login with email: %s", req.Email)
	return &pb.LoginUserResponse{}, nil
}
