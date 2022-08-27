package drivers

import (
	"context"
	"rss/adapters/controllers"
	"rss/adapters/gateways"
	"rss/adapters/presenters"
	"rss/database"
	"rss/usecase/interactors"

	"github.com/labstack/echo"
)

func InitializeNewsDriver(ctx context.Context) (News, error) {
	e := echo.New()
	outputFactory := NewOutputFactory()
	inputFactory := NewInputFactory()
	repositoryFactory := NewRepositoryFactory()
	feed := NewFeedFactory()
	news := controllers.NewNewsController(outputFactory, inputFactory, repositoryFactory, feed)
	newsDriver := NewNewsDriver(e, news)
	return newsDriver, nil
}

func NewFeedFactory() gateways.FeedFactory {
	return &database.MyFeedFactory{}
}

func NewOutputFactory() controllers.OutputFactory {
	return presenters.NewsOutputPort
}

func NewInputFactory() controllers.InputFactory {
	return interactors.KeywordInputPort
}

func NewRepositoryFactory() controllers.RepositoryFactory {
	return gateways.NewsRepository
}
