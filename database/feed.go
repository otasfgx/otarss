package database

import (
	"context"
	"fmt"
	"net/url"

	"github.com/mmcdole/gofeed"
)

type MyFeedFactory struct{}

func (f *MyFeedFactory) NewFeed(ctx context.Context, keyword string) (*gofeed.Feed, error) {
	rssurl := fmt.Sprintf("https://news.google.com/rss/search?q=%s&hl=ja&gl=JP&ceid=JP:ja", url.QueryEscape(keyword))
	feed, err := gofeed.NewParser().ParseURL(rssurl)

	if err != nil {
		return nil, err
	}
	return feed, nil
}
