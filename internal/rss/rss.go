package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func FetchFeed(ctx context.Context, url string) (*RSSFeed, error) {
	dat, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("feed request: %w", err)
	}

	var feed RSSFeed
	err = xml.NewDecoder(dat.Body).Decode(&feed)
	if err != nil {
		return nil, fmt.Errorf("xml decode: %w", err)
	}

	return &feed, nil
}
