package usecase

import (
	"context"
	"github.com/maei/mongodb_mock_hexa/src/domain"
	"log"
)

type serviceArticle struct {
	repoArticle domain.DAOInterfaceArticle
	grpcArticle domain.GRPCArticleInterface
}

func NewServiceArticle(re domain.DAOInterfaceArticle, gr domain.GRPCArticleInterface) domain.ServiceArticleInterface {
	return &serviceArticle{
		repoArticle: re,
		grpcArticle: gr,
	}
}

func (s *serviceArticle) Store(ctx context.Context, article *domain.Article) error {
	storeErr := s.repoArticle.StoreArticle(ctx, article)
	if storeErr != nil {
		log.Println(storeErr)
		return storeErr
	}
	return nil
}

func (s *serviceArticle) Find(ctx context.Context, id string) (*domain.Article, error) {
	return nil, nil
}
