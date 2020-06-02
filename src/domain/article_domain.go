package domain

import (
	"context"
)

type Article struct {
	Name     string `json:"name" bson:"name"`
	Quantity int64  `json:"quantity" bson:"quantity"`
}

type APIInterfaceGRPCArticle interface {
	SendArticle(ctx context.Context, article *Article) (string, error)
}

type ServiceArticleInterface interface {
	Store(ctx context.Context, article *Article) error
	Find(ctx context.Context, id string) (*Article, error)
}

type RepositoryArticleInterface interface {
	StoreArticle(ctx context.Context, article *Article) error
	FindArticleByID(ctx context.Context, id string) (*Article, error)
}
