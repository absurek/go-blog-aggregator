package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/absurek/go-blog-aggregator/internal/application"
	"github.com/absurek/go-blog-aggregator/internal/cli"
	"github.com/absurek/go-blog-aggregator/internal/database"
	"github.com/absurek/go-blog-aggregator/internal/rss"
	"github.com/absurek/go-blog-aggregator/internal/utils"
	"github.com/google/uuid"
)

func scrapeFeed(app *application.Application) error {
	tx, err := app.DB.Begin()
	if err != nil {
		return fmt.Errorf("tx begin: %w", err)
	}
	defer tx.Rollback()

	qtx := app.DBQueries.WithTx(tx)

	feed, err := qtx.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("get next feed to fecht (feed_id=%s feed_name=%s): %w", feed.ID, feed.Name, err)
	}

	log.Printf("Fetching %s from %s", feed.Name, feed.Url)
	rssFeed, err := rss.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("fetch feed (feed_id=%s feed_name=%s): %w", feed.ID, feed.Name, err)
	}

	err = qtx.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return fmt.Errorf("mark feed fetched (feed_id=%s feed_name=%s): %w", feed.ID, feed.Name, err)
	}

	for _, item := range rssFeed.Channel.Item {
		err = qtx.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			FeedID:      feed.ID,
			Url:         item.Link,
			Title:       utils.ParseNullString(item.Title),
			Description: utils.ParseNullString(item.Description),
			PublishedAt: utils.ParseNullTime(item.PubDate),
		})
		if err != nil {
			return fmt.Errorf("create post (feed_id=%s feed_name=%s): %w", feed.ID, feed.Name, err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("tx commit (feed_id=%s feed_name=%s): %w", feed.ID, feed.Name, err)
	}

	return nil
}

func AggHandler(app *application.Application, cmd cli.Command) error {
	if len(cmd.Args) != 1 {
		return errors.New("agg expects exactly 1 argument")
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("parse time between requests (cmd=%s): %w", cmd, err)
	}

	log.Printf("Collecting feeds every %s\n", timeBetweenReqs)

	ticker := time.NewTicker(timeBetweenReqs)
	for ; ; <-ticker.C {
		err = scrapeFeed(app)
		if err != nil {
			log.Printf("Error: scrape feed: %v\n", err)
		}
	}
}
