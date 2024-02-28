package toblog

import (
	"context"
	"log"

	"github.com/sirupsen/logrus"
	"github.com/the-monkeys/the_monkeys/apis/interservice/blogs/pb"
	"google.golang.org/grpc"
)

type BlogClient struct {
	blogServiceClient pb.BlogServiceClient
}

func NewClient(conn *grpc.ClientConn) *BlogClient {
	return &BlogClient{blogServiceClient: pb.NewBlogServiceClient(conn)}
}

// TODO: Update error handling and return error
func (bc *BlogClient) UpdateBlogsUserDeactivated(email string) {
	logrus.Infof("Requesting to update user's blog status for: %s", email)
	// send a request to the server
	res, err := bc.blogServiceClient.SetUserDeactivated(context.Background(), &pb.SetUserDeactivatedReq{
		Email: email,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// print the response from the server
	log.Printf("Greeting: %s", res.Message)
}
