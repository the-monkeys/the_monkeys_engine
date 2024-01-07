package article

import (
	"fmt"

	"github.com/the-monkeys/the_monkeys/microservices/api_gateway/config"
	"github.com/the-monkeys/the_monkeys/microservices/article_and_post/pkg/pb"
	"google.golang.org/grpc"
)

type ArticleServiceClient struct {
	Client pb.ArticleServiceClient
}

func InitArticleServiceClient(c *config.Address) pb.ArticleServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.StoryService, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
		return nil
	}

	return pb.NewArticleServiceClient(cc)
}
