package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	_articleService "github.com/maei/mongodb_mock_hexa/src/article"
	"github.com/maei/mongodb_mock_hexa/src/clients/mongodb"
	_articleHttpDelivery "github.com/maei/mongodb_mock_hexa/src/controller"
	_articleRepository "github.com/maei/mongodb_mock_hexa/src/repository"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	mongoURL     = "mongodb://localhost:27017"
	mongoTimeout = 30
	mongoDB      = "article"
	mongoColl    = "article_coll"
)

func main() {
	client, err := mongodb.NewMongoClient(mongoURL, mongoTimeout)
	if err != nil {
		log.Println("error while connecting to mongodb")
		panic(0)
	}
	e := echo.New()

	articleRepo := _articleRepository.NewMongoArticleRepository(client, mongoTimeout, mongoDB, mongoColl)
	articleUseCase := _articleService.NewServiceArticle(articleRepo, nil)
	_articleHttpDelivery.NewArticleController(e, articleUseCase)

	errs := make(chan error, 2)
	go func() {
		log.Println("listening on port :8000")
		errs <- e.Start(httpPort())
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	log.Printf("Terminated %s", <-errs)
	client.Disconnect(context.Background())
}

func httpPort() string {
	port := "8000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return fmt.Sprintf(":%s", port)
}
