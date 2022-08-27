package gateways

import (
	"context"
	"errors"
	"fmt"
	"rss/entities"
	"rss/usecase/ports"

	"github.com/mmcdole/gofeed"
)

type FeedFactory interface {
	NewFeed(ctx context.Context, keyword string) (*gofeed.Feed, error)
}

type NewsGateway struct {
	feedFactory FeedFactory
}

func NewsRepository(feedFactory FeedFactory) ports.NewsRepository {
	return &NewsGateway{
		feedFactory: feedFactory,
	}
}

func (gateway *NewsGateway) GetNews(ctx context.Context, keyword string) ([]entities.News, error) {
	if keyword == "" {
		return nil, errors.New("request parameter is insufficient")
	}
	feed, err := gateway.feedFactory.NewFeed(ctx, keyword)
	if err != nil {
		return nil, fmt.Errorf("failed to get google news feed")
	}

	return getResult(ctx, feed, keyword)
}

func getResult(ctx context.Context, feed *gofeed.Feed, keyword string) ([]entities.News, error) {
	news := []entities.News{}
	for idx, item := range feed.Items {
		if idx > 0 {
			break
		}
		news = append(news, entities.News{
			Title:   item.Title,
			Link:    item.Link,
			Keyword: keyword,
			Source:  "Google News",
		})
	}
	return news, nil
}
