package presenters

import (
	"log"
	"net/http"
	"rss/entities"
	"rss/usecase/ports"

	"github.com/labstack/echo"
)

type NewsPresenter struct {
	ctx echo.Context
}

func NewsOutputPort(ctx echo.Context) ports.NewsOutputPort {
	return &NewsPresenter{
		ctx: ctx,
	}
}

func (newsPresenter *NewsPresenter) OutputNews(news []entities.News) error {
	return newsPresenter.ctx.JSON(http.StatusOK, news)
}

func (newsPresenter *NewsPresenter) OutputError(err error) error {
	log.Fatal(err)
	return newsPresenter.ctx.JSON(http.StatusInternalServerError, err)
}
