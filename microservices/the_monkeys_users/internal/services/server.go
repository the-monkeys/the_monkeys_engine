package services

import (
	"context"
	"fmt"

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
	us.log.Infof("user %v has requested profile info.", req.Email)
	_, err := us.dbConn.CheckIfEmailExist(req.Email)
	if err != nil {
		us.log.Errorf("the user doesn't exists: %v", err)
		return nil, err
	}

	userDetails, err := us.dbConn.GetMyProfile(req.Email)
	if err != nil {
		us.log.Errorf("error while finding the user profile: %v", err)
		return nil, err
	}

	return &pb.UserProfileRes{
		ProfileId:   userDetails.ProfileId,
		Username:    userDetails.Username,
		FirstName:   userDetails.FirstName,
		LastName:    userDetails.LastName,
		DateOfBirth: userDetails.DateOfBirth.Time.String(),
		Bio:         userDetails.Bio.String,
		AvatarUrl:   userDetails.AvatarUrl.String,
		// CreatedAt:     userDetails.CreatedAt.Time.String(),
		// UpdatedAt:     userDetails.UpdatedAt,
		Address:       userDetails.Address.String,
		ContactNumber: userDetails.ContactNumber.Int64,
		UserStatus:    userDetails.UserStatus,
	}, err
}

func (us *UserSvc) GetUserActivities(ctx context.Context, req *pb.UserActivityReq) (*pb.UserActivityRes, error) {
	logrus.Infof("Trying to fetch user activities for: %v", req.Email)

	return &pb.UserActivityRes{}, nil
}
func (us *UserSvc) UpdateUserProfile(ctx context.Context, req *pb.UpdateUserProfileReq) (*pb.UpdateUserProfileRes, error){
	fmt.Printf("req: %+v\n", req)
	return nil,nil
}
