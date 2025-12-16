package cmd

import (
	"context"
	"fmt"

	"github.com/absurek/go-blog-aggregator/internal/application"
	"github.com/absurek/go-blog-aggregator/internal/cli"
	"github.com/absurek/go-blog-aggregator/internal/rss"
)

func AggHandler(app *application.Application, cmd cli.Command) error {
	const url = "https://www.wagslane.dev/index.xml"
	feed, err := rss.FetchFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("fetch feed: %w", err)
	}

	fmt.Println(feed)

	return nil
}
