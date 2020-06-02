package api

import (
	"github.com/maei/mongodb_mock_hexa/src/domain"
	"google.golang.org/grpc"
)

type articleGRPC struct {
	client *grpc.ClientConn
}

func NewArticleGRPC(articleClient) domain.GRPCArticleInterface {
	return &articleGRPC{client: client}
}

func (ag *articleGRPC) SendArticle() {
}
