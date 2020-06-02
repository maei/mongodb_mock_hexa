package controller

import (
	"context"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/maei/mongodb_mock_hexa/src/domain"
	"io/ioutil"
	"log"
	"net/http"
)

type ArticleControllerInterface interface {
	PostArticle(c echo.Context) error
	FindArticle(c echo.Context) error
}

type articleController struct {
	articleService domain.ServiceArticleInterface
}

func NewArticleController(e *echo.Echo, art domain.ServiceArticleInterface) ArticleControllerInterface {
	handler := &articleController{
		articleService: art,
	}
	e.GET("/articles", handler.FindArticle)
	e.POST("/articles", handler.PostArticle)

	return handler
}

func (a *articleController) PostArticle(c echo.Context) error {
	defer c.Request().Body.Close()

	jsonBody := c.Request().Body
	bs, err := ioutil.ReadAll(jsonBody)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})

	}
	art := &domain.Article{}

	artErr := json.Unmarshal(bs, art)
	if artErr != nil {
		log.Println(artErr)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": artErr.Error(),
		})

	}

	stoErr := a.articleService.Store(context.Background(), art)
	if stoErr != nil {
		log.Println(stoErr)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "error writing data to database",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "store article successfully",
	})
}

func (a *articleController) FindArticle(c echo.Context) error {

	return nil
}
