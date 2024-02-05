package server

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/the-monkeys/the_monkeys/apis/serviceconn/gateway_user/pb"
)

type UserSvc struct {
	pb.UnimplementedUserServiceServer
}

func NewUserSvc() *UserSvc {
	return &UserSvc{}
}

func (us *UserSvc) GetUserActivities(ctx context.Context, req *pb.UserActivityReq) (*pb.UserActivityRes, error) {
	logrus.Infof("Trying to fetch user activities for: %v", req.Email)

	return &pb.UserActivityRes{}, nil
}
