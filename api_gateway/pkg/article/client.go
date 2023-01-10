package article

import (
	"fmt"

	"github.com/89minutes/the_new_project/api_gateway/config"
	"github.com/89minutes/the_new_project/article_and_post/pkg/pb"
	"google.golang.org/grpc"
)

type ArticleServiceClient struct {
	Client pb.ArticleServiceClient
}

func InitArticleServiceClient(c *config.Config) pb.ArticleServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.ArticleSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewArticleServiceClient(cc)
}
