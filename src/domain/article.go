package domain

import (
	"context"
)

type Article struct {
	Name     string `json:"name" bson:"name"`
	Quantity int64  `json:"quantity" bson:"quantity"`
}

type GRPCArticleInterface interface {
	SendArticle()
}

type ServiceArticleInterface interface {
	Store(ctx context.Context, article *Article) error
	Find(ctx context.Context, id string) (*Article, error)
}

type DAOInterfaceArticle interface {
	StoreArticle(ctx context.Context, article *Article) error
	FindArticleByID(ctx context.Context, id string) (*Article, error)
}
