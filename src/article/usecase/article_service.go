package usecase

import (
	"context"
	"github.com/maei/mongodb_mock_hexa/src/domain"
	"log"
)

type serviceArticle struct {
	repoArticle domain.RepositoryArticleInterface
	grpcArticle domain.APIInterfaceGRPCArticle
}

func NewServiceArticle(re domain.RepositoryArticleInterface, gr domain.APIInterfaceGRPCArticle) domain.ServiceArticleInterface {
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
	res, err := s.grpcArticle.SendArticle(ctx, article)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(res)
	return nil
}

func (s *serviceArticle) Find(ctx context.Context, id string) (*domain.Article, error) {
	return nil, nil
}
