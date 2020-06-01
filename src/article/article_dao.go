package article

import "context"

type DAOInterfaceArticle interface {
	StoreArticle(ctx context.Context, article *Article) error
	FindArticleByID(ctx context.Context, id string) (*Article, error)
}
