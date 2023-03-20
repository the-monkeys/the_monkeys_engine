package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"database/sql"

	isv "github.com/89minutes/the_new_project/apis/interservice/blogs/pb"
	"github.com/89minutes/the_new_project/services/user_service/service/database"
	"github.com/89minutes/the_new_project/services/user_service/service/pb"
	"github.com/89minutes/the_new_project/services/user_service/service/utils"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	db          *database.UserDbHandler
	log         *logrus.Logger
	blogService isv.BlogServiceClient
	pb.UnimplementedUserServiceServer
}

func NewUserService(usd *database.UserDbHandler, log *logrus.Logger, blogService isv.BlogServiceClient) *UserService {
	return &UserService{db: usd, log: log, blogService: blogService}
}

func (us *UserService) GetMyProfile(ctx context.Context, req *pb.GetMyProfileReq) (*pb.GetMyProfileRes, error) {
	us.log.Infof("user %v has requested for their profile", req.GetId())

	resp, err := us.db.GetMyProfile(req.GetId())
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			us.log.Infof("cannot fine the user with id %v ", req.GetId())
			return nil, status.Errorf(codes.NotFound, "failed to find the record, error: %v", err)
		case sql.ErrTxDone:
			log.Println("The transaction has already been committed or rolled back.")
			return nil, status.Errorf(codes.Internal, "failed to find the record, error: %v", err)
		case sql.ErrConnDone:
			log.Println("The database connection has been closed.")
			return nil, status.Errorf(codes.Unavailable, "failed to find the record, error: %v", err)
		default:
			log.Printf("An internal server error occurred: %v\n", err)
			return nil, status.Errorf(codes.Internal, "failed to find the record, error: %v", err)
		}
	}

	return resp, nil
}

// TODO: Send an email after profile update
func (us *UserService) SetMyProfile(ctx context.Context, req *pb.SetMyProfileReq) (*pb.SetMyProfileRes, error) {
	us.log.Infof("the user %s has requested to update profile", req.GetEmail())
	if err := us.db.UpdateMyProfile(req); err != nil {
		switch err {
		case sql.ErrNoRows:
			us.log.Infof("cannot fine the user with id %v ", req.GetEmail())
			return nil, status.Errorf(codes.NotFound, "failed to find the record, error: %v", err)
		case sql.ErrTxDone:
			log.Println("The transaction has already been committed or rolled back.")
			return nil, status.Errorf(codes.Internal, "failed to find the record, error: %v", err)
		case sql.ErrConnDone:
			log.Println("The database connection has been closed.")
			return nil, status.Errorf(codes.Unavailable, "failed to find the record, error: %v", err)
		default:
			log.Printf("An internal server error occurred: %v\n", err)
			return nil, status.Errorf(codes.Internal, "failed to find the record, error: %v", err)
		}
	}
	return &pb.SetMyProfileRes{
		Status: http.StatusOK,
	}, nil
}

func (us *UserService) UploadProfile(stream pb.UserService_UploadProfileServer) error {
	var byteSlice []byte
	var chunkId int64
	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		byteSlice = append(byteSlice, chunk.Data...)
		chunkId = chunk.Id
	}
	us.log.Infof("updating profile pic for user: %v", chunkId)
	err := us.db.UploadProfilePic(byteSlice, chunkId)
	if err != nil {
		return err
	}

	return stream.SendAndClose(&pb.UploadProfilePicRes{
		Status: http.StatusOK,
		Id:     chunkId,
	})
}

func (us *UserService) Download(req *pb.GetProfilePicReq, stream pb.UserService_DownloadServer) error {
	xb := []byte{}
	err := us.db.Psql.QueryRow("SELECT profile_pic from the_monkeys_user WHERE id=$1", req.Id).Scan(&xb)
	if err != nil {
		us.log.Errorf("cannot get the profile pic, error: %v", err)
		return utils.Errors(err)
	}

	if err := stream.Send(&pb.GetProfilePicRes{
		Data: xb,
	}); err != nil {
		us.log.Errorf("error while sending stream, error %+v", err)
	}

	return nil
}

func (us *UserService) DeleteMyProfile(ctx context.Context, req *pb.DeleteMyAccountReq) (*pb.DeleteMyAccountRes, error) {
	us.log.Infof("The used %v is requested to delete their account", req.GetId())

	// TODO: Run transaction
	err := us.db.DeactivateMyAccount(req.GetId())
	if err != nil {
		us.log.Errorf("cannot get the profile pic, error: %v", err)
		return nil, utils.Errors(err)
	}

	resp, err := us.db.GetMyProfile(req.GetId())
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			us.log.Infof("cannot fine the user with id %v ", req.GetId())
			return nil, status.Errorf(codes.NotFound, "failed to find the record, error: %v", err)
		case sql.ErrTxDone:
			log.Println("The transaction has already been committed or rolled back.")
			return nil, status.Errorf(codes.Internal, "failed to find the record, error: %v", err)
		case sql.ErrConnDone:
			log.Println("The database connection has been closed.")
			return nil, status.Errorf(codes.Unavailable, "failed to find the record, error: %v", err)
		default:
			log.Printf("An internal server error occurred: %v\n", err)
			return nil, status.Errorf(codes.Internal, "failed to find the record, error: %v", err)
		}
	}
	// TODO: Set all the users status key in the blog as disabled and not show users blog to the portal.
	isvRes, err := us.blogService.SetUserDeactivated(context.Background(), &isv.SetUserDeactivatedReq{Email: resp.Email})
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, status.Errorf(codes.Internal, "failed to find the record, error: %v", err)
	}

	fmt.Printf("isvRes: %v\n", isvRes)

	return &pb.DeleteMyAccountRes{
		Status: http.StatusOK,
		Id:     req.Id,
	}, nil
}
