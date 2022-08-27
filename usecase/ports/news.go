package ports

import (
	"context"
	"rss/entities"
)

type KeywordInputPort interface {
	GetNews(ctx context.Context, keyword string) error
}

type NewsOutputPort interface {
	OutputNews([]entities.News) error
	OutputError(error) error
}

type NewsRepository interface {
	GetNews(ctx context.Context, keyword string) ([]entities.News, error)
}
