package controllers

import (
	"context"
	"rss/adapters/gateways"
	"rss/usecase/ports"

	"github.com/labstack/echo"
)

type News interface {
	GetNews(ctx context.Context) func(c echo.Context) error
}

type OutputFactory func(echo.Context) ports.NewsOutputPort
type InputFactory func(ports.NewsOutputPort, ports.NewsRepository) ports.KeywordInputPort
type RepositoryFactory func(gateways.FeedFactory) ports.NewsRepository

type NewsController struct {
	outputFactory     OutputFactory
	inputFactory      InputFactory
	repositoryFactory RepositoryFactory
	feedFactory       gateways.FeedFactory
}

func NewNewsController(outputFactory OutputFactory, inputFactory InputFactory, repositoryFactory RepositoryFactory, feedFactory gateways.FeedFactory) News {
	return &NewsController{
		outputFactory:     outputFactory,
		inputFactory:      inputFactory,
		repositoryFactory: repositoryFactory,
		feedFactory:       feedFactory,
	}
}

func (newsController *NewsController) GetNews(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		keyword := c.FormValue("keyword")
		return newsController.newInputPort(c).GetNews(ctx, keyword)
	}
}

func (newsController *NewsController) newInputPort(c echo.Context) ports.KeywordInputPort {
	outputPort := newsController.outputFactory(c)
	repository := newsController.repositoryFactory(newsController.feedFactory)
	return newsController.inputFactory(outputPort, repository)
}
