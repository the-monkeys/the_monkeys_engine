package server

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/the-monkeys/the_monkeys/apis/serviceconn/gateway_user/pb"
	"github.com/the-monkeys/the_monkeys/microservices/the_monkeys_users/internal/database"
)

type UserSvc struct {
	dbConn database.UserDb
	log    *logrus.Logger
	pb.UnimplementedUserServiceServer
}

func NewUserSvc(dbConn database.UserDb, log *logrus.Logger) *UserSvc {
	return &UserSvc{
		dbConn: dbConn,
		log:    log,
	}
}

func (us *UserSvc) GetUserProfile(ctx context.Context, req *pb.UserProfileReq) (*pb.UserProfileRes, error) {
	user, err := us.dbConn.CheckIfEmailExist(req.Email)
	if err != nil {
		us.log.Errorf("error while finding the user profile: %v", err)
		return nil, err
	}
	return &pb.UserProfileRes{
		ProfileId: user.ProfileId,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		// DateOfBirth:   user.Dat,
		// RoleId:        user.Role,
		// Bio:           user.ProfileId,
		// AvatarUrl:     user.ProfileId,
		// CreatedAt:     user.ProfileId,
		// UpdatedAt:     user.ProfileId,
		// Address:       user.ProfileId,
		// ContactNumber: user.ProfileId,
		UserStatus: user.UserStatus,
	}, nil
}

func (us *UserSvc) GetUserActivities(ctx context.Context, req *pb.UserActivityReq) (*pb.UserActivityRes, error) {
	logrus.Infof("Trying to fetch user activities for: %v", req.Email)

	return &pb.UserActivityRes{}, nil
}
