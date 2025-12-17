package cmd

import (
	"context"
	"errors"
	"fmt"

	"github.com/absurek/go-blog-aggregator/internal/application"
	"github.com/absurek/go-blog-aggregator/internal/cli"
	"github.com/absurek/go-blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func AddfeedHandler(app *application.Application, cmd cli.Command, user database.User) error {
	if len(cmd.Args) != 2 {
		return errors.New("addfeed expects exactly 2 arguments")
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	tx, err := app.DB.Begin()
	if err != nil {
		return fmt.Errorf("db tx begin: %w", err)
	}
	defer tx.Rollback()
	qtx := app.DBQueries.WithTx(tx)

	feed, err := qtx.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:     uuid.New(),
		UserID: user.ID,
		Name:   name,
		Url:    url,
	})
	if err != nil {
		return fmt.Errorf("db create feed (username=%s feed=%s url=%s): %w", user.Name, name, url, err)
	}

	_, err = qtx.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:     uuid.New(),
		FeedID: feed.ID,
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("create follow (username=%s url=%s): %w", user.Name, url, err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("db tx commit: %w", err)
	}

	fmt.Println()
	fmt.Println("New feed created:")
	fmt.Println("- id:", feed.ID)
	fmt.Println("- user_id:", feed.UserID)
	fmt.Println("- name:", feed.Name)
	fmt.Println("- url:", feed.Url)
	fmt.Println()
	fmt.Printf("You are now following %s at %s\n", name, url)
	fmt.Println()

	return nil
}
