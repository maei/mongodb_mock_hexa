package article

import (
	"context"
	"log"
)

type ServiceArticleInterface interface {
	Store(ctx context.Context, article *Article) error
	Find(ctx context.Context, id string) (*Article, error)
}

type serviceArticle struct {
	repoArticle DAOInterfaceArticle
	grpcArticle GRPCArticleInterface
}

func NewServiceArticle(re DAOInterfaceArticle, gr GRPCArticleInterface) ServiceArticleInterface {
	return &serviceArticle{
		repoArticle: re,
		grpcArticle: gr,
	}
}

func (s *serviceArticle) Store(ctx context.Context, article *Article) error {
	storeErr := s.repoArticle.StoreArticle(ctx, article)
	if storeErr != nil {
		log.Println(storeErr)
		return storeErr
	}
	return nil
}

func (s *serviceArticle) Find(ctx context.Context, id string) (*Article, error) {
	return nil, nil
}
