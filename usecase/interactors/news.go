package interactors

import (
	"context"
	"rss/usecase/ports"
)

type NewsInteractor struct {
	OutputPort ports.NewsOutputPort
	Repository ports.NewsRepository
}

func KeywordInputPort(outputPort ports.NewsOutputPort, repository ports.NewsRepository) ports.KeywordInputPort {
	return &NewsInteractor{
		OutputPort: outputPort,
		Repository: repository,
	}
}

func (newsInteractor *NewsInteractor) GetNews(ctx context.Context, keyword string) error {
	news, err := newsInteractor.Repository.GetNews(ctx, keyword)

	if err != nil {
		return newsInteractor.OutputPort.OutputError(err)
	}

	return newsInteractor.OutputPort.OutputNews(news)
}
