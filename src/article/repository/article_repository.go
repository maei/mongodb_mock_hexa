package repository

import (
	"context"
	"errors"
	"github.com/maei/mongodb_mock_hexa/src/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type mongoArticleRepository struct {
	client     *mongo.Client
	database   string
	timeout    time.Duration
	collection string
}

func NewMongoArticleRepository(client *mongo.Client, mongoTimeout int, mongoDB string, mongoColl string) domain.DAOInterfaceArticle {
	return &mongoArticleRepository{
		client:     client,
		database:   mongoDB,
		timeout:    time.Duration(mongoTimeout) * time.Second,
		collection: mongoColl,
	}
}

func (ma *mongoArticleRepository) StoreArticle(ctx context.Context, article *domain.Article) error {
	log.Println("StoreArticle called")
	ctx, cancel := context.WithTimeout(context.Background(), ma.timeout)
	defer cancel()
	coll := ma.client.Database(ma.database).Collection(ma.collection)
	res, err := coll.InsertOne(ctx, article)
	if err != nil {
		log.Println(err)
		return errors.New("error while writing article to database")
	}
	log.Println(res.InsertedID)
	return nil
}
func (ma *mongoArticleRepository) FindArticleByID(ctx context.Context, id string) (*domain.Article, error) {
	log.Println("FindArticleByID called")

	return nil, nil
}
