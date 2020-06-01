package controller

import (
	"context"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/maei/mongodb_mock_hexa/src/article"
	"io/ioutil"
	"log"
	"net/http"
)

type ArticleController struct {
	ArticleService article.ServiceArticleInterface
}

func NewArticleController(e *echo.Echo, art article.ServiceArticleInterface) {
	handler := &ArticleController{
		ArticleService: art,
	}
	e.GET("/articles", handler.FindArticle)
	e.POST("/articles", handler.PostArticle)

}

func (a *ArticleController) PostArticle(c echo.Context) error {
	defer c.Request().Body.Close()

	jsonBody := c.Request().Body
	bs, err := ioutil.ReadAll(jsonBody)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})

	}
	art := &article.Article{}

	artErr := json.Unmarshal(bs, art)
	if artErr != nil {
		log.Println(artErr)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": artErr.Error(),
		})

	}

	stoErr := a.ArticleService.Store(context.Background(), art)
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

func (a *ArticleController) FindArticle(c echo.Context) error {

	return nil
}
