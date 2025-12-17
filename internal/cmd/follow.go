package cmd

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/absurek/go-blog-aggregator/internal/application"
	"github.com/absurek/go-blog-aggregator/internal/cli"
	"github.com/absurek/go-blog-aggregator/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func FollowHandler(app *application.Application, cmd cli.Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return errors.New("follow expects exactly 1 argument")
	}

	url := cmd.Args[0]

	feed, err := app.DBQueries.GetFeedByURL(context.Background(), url)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return fmt.Errorf("no feed found with url '%s'", url)
		default:
			return fmt.Errorf("get feed (username=%s cmd=%s): %w", user.Name, cmd, err)
		}
	}

	follow, err := app.DBQueries.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:     uuid.New(),
		FeedID: feed.ID,
		UserID: user.ID,
	})
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" { // unique violation
				return errors.New("already following this feed")
			}
		}

		return fmt.Errorf("create follow (username=%s cmd=%s): %w", user.Name, cmd, err)
	}

	fmt.Printf("You are now following %s at %s\n", follow.FeedName, url)

	return nil
}
