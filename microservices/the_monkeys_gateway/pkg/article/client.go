package article

import (
	"fmt"

	"github.com/the-monkeys/the_monkeys/config"
	"github.com/the-monkeys/the_monkeys/microservices/the_monkeys_gateway/pkg/article/pb"

	"google.golang.org/grpc"
)

type ArticleServiceClient struct {
	Client pb.ArticleServiceClient
}

func InitArticleServiceClient(c *config.Config) pb.ArticleServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.Microservices.TheMonkeysBlog, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
		return nil
	}

	return pb.NewArticleServiceClient(cc)
}
