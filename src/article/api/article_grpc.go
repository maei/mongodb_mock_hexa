package api

import (
	"context"
	"errors"
	"github.com/maei/mongodb_mock_hexa/src/article/proto"
	"github.com/maei/mongodb_mock_hexa/src/domain"
	"log"
	"time"
)

type articleGRPC struct {
	art articlepb.ArticleAPIClient
}

func NewArticleGRPC(articleClient articlepb.ArticleAPIClient) domain.APIInterfaceGRPCArticle {
	return &articleGRPC{
		art: articleClient,
	}
}

func (ag *articleGRPC) SendArticle(ctx context.Context, article *domain.Article) (string, error) {
	artPb := &articlepb.ArticleRequest{
		Name: article.Name,
	}
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	res, err := ag.art.StoreArticle(ctx, artPb)
	if err != nil {
		log.Println(err)
		return "", errors.New("error while getting response from grpc-server")
	}

	return res.GetResult(), nil
}
