package server

// type UserService struct {
// 	db          *database.UserDbHandler
// 	log         *logrus.Logger
// 	blogService isv.BlogServiceClient
// 	pb.UnimplementedUserServiceServer
// }

// func NewUserService(usd *database.UserDbHandler, log *logrus.Logger, blogService isv.BlogServiceClient) *UserService {
// 	return &UserService{db: usd, log: log, blogService: blogService}
// }

// func (us *UserService) GetMyProfile(ctx context.Context, req *pb.GetMyProfileReq) (*pb.GetMyProfileRes, error) {
// 	us.log.Infof("user %v has requested for their profile", req.GetId())

// 	resp, err := us.db.GetMyProfile(req.GetId())
// 	if err != nil {
// 		switch err {
// 		case sql.ErrNoRows:
// 			us.log.Infof("cannot fine the user with id %v ", req.GetId())
// 			return nil, status.Errorf(codes.NotFound, "failed to find the record, error: %v", err)
// 		case sql.ErrTxDone:
// 			log.Println("The transaction has already been committed or rolled back.")
// 			return nil, status.Errorf(codes.Internal, "failed to find the record, error: %v", err)
// 		case sql.ErrConnDone:
// 			log.Println("The database connection has been closed.")
// 			return nil, status.Errorf(codes.Unavailable, "failed to find the record, error: %v", err)
// 		default:
// 			log.Printf("An internal server error occurred: %v\n", err)
// 			return nil, status.Errorf(codes.Internal, "failed to find the record, error: %v", err)
// 		}
// 	}

// 	return resp, nil
// }

// // TODO: Send an email after profile update
// func (us *UserService) SetMyProfile(ctx context.Context, req *pb.SetMyProfileReq) (*pb.SetMyProfileRes, error) {
// 	us.log.Infof("the user %s has requested to update profile", req.GetEmail())
// 	if err := us.db.UpdateMyProfile(req); err != nil {
// 		switch err {
// 		case sql.ErrNoRows:
// 			us.log.Infof("cannot fine the user with id %v ", req.GetEmail())
// 			return nil, status.Errorf(codes.NotFound, "failed to find the record, error: %v", err)
// 		case sql.ErrTxDone:
// 			log.Println("The transaction has already been committed or rolled back.")
// 			return nil, status.Errorf(codes.Internal, "failed to find the record, error: %v", err)
// 		case sql.ErrConnDone:
// 			log.Println("The database connection has been closed.")
// 			return nil, status.Errorf(codes.Unavailable, "failed to find the record, error: %v", err)
// 		default:
// 			log.Printf("An internal server error occurred: %v\n", err)
// 			return nil, status.Errorf(codes.Internal, "failed to find the record, error: %v", err)
// 		}
// 	}
// 	return &pb.SetMyProfileRes{
// 		Status: http.StatusOK,
// 	}, nil
// }

// func (us *UserService) DeleteMyProfile(ctx context.Context, req *pb.DeleteMyAccountReq) (*pb.DeleteMyAccountRes, error) {
// 	us.log.Infof("The used %v is requested to delete their account", req.GetId())

// 	// TODO: Run transaction
// 	err := us.db.DeactivateMyAccount(req.GetId())
// 	if err != nil {
// 		us.log.Errorf("cannot get the profile pic, error: %v", err)
// 		return nil, utils.Errors(err)
// 	}

// 	resp, err := us.db.GetMyProfile(req.GetId())
// 	if err != nil {
// 		switch err {
// 		case sql.ErrNoRows:
// 			us.log.Infof("cannot fine the user with id %v ", req.GetId())
// 			return nil, status.Errorf(codes.NotFound, "failed to find the record, error: %v", err)
// 		case sql.ErrTxDone:
// 			log.Println("The transaction has already been committed or rolled back.")
// 			return nil, status.Errorf(codes.Internal, "failed to find the record, error: %v", err)
// 		case sql.ErrConnDone:
// 			log.Println("The database connection has been closed.")
// 			return nil, status.Errorf(codes.Unavailable, "failed to find the record, error: %v", err)
// 		default:
// 			log.Printf("An internal server error occurred: %v\n", err)
// 			return nil, status.Errorf(codes.Internal, "failed to find the record, error: %v", err)
// 		}
// 	}
// 	// TODO: Set all the users status key in the blog as disabled and not show users blog to the portal.
// 	isvRes, err := us.blogService.SetUserDeactivated(context.Background(), &isv.SetUserDeactivatedReq{Email: resp.Email})
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 		return nil, status.Errorf(codes.Internal, "failed to find the record, error: %v", err)
// 	}

// 	fmt.Printf("isvRes: %v\n", isvRes)

// 	return &pb.DeleteMyAccountRes{
// 		Status: http.StatusOK,
// 		Id:     req.Id,
// 	}, nil
// }
