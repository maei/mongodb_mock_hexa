package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/maei/mongodb_mock_hexa/src/article/api"
	"github.com/maei/mongodb_mock_hexa/src/article/controller"
	"github.com/maei/mongodb_mock_hexa/src/article/proto"
	"github.com/maei/mongodb_mock_hexa/src/article/repository"
	"github.com/maei/mongodb_mock_hexa/src/article/usecase"
	"github.com/maei/mongodb_mock_hexa/src/clients/grpc"
	"github.com/maei/mongodb_mock_hexa/src/clients/mongodb"
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
	grpcURL      = "localhost:50051"
)

func main() {
	mclient, err := mongodb.NewMongoClient(mongoURL, mongoTimeout)
	if err != nil {
		log.Println("error while connecting to mongodb")
		panic(0)
	}
	conn, grpcErr := grpc.NewGRPCClient(grpcURL)
	if grpcErr != nil {
		log.Println("error while connecting to gRPC server")
		panic(0)
	}
	defer conn.Close()
	e := echo.New()

	articleGRPCConn := articlepb.NewArticleAPIClient(conn)
	articleGRPC := api.NewArticleGRPC(articleGRPCConn)
	articleRepo := repository.NewArticleRepository(mclient, mongoTimeout, mongoDB, mongoColl)
	articleUseCase := usecase.NewServiceArticle(articleRepo, articleGRPC)
	controller.NewArticleController(e, articleUseCase)

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
	mclient.Disconnect(context.Background())
}

func httpPort() string {
	port := "8000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return fmt.Sprintf(":%s", port)
}
