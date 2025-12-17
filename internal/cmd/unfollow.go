package cmd

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/absurek/go-blog-aggregator/internal/application"
	"github.com/absurek/go-blog-aggregator/internal/cli"
	"github.com/absurek/go-blog-aggregator/internal/database"
)

func UnfollowHandler(app *application.Application, cmd cli.Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return errors.New("unfollow expects exactly 1 argument")
	}

	url := cmd.Args[0]

	feed, err := app.DBQueries.GetFeedByURL(context.Background(), url)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return fmt.Errorf("no feed found with url '%s' (username=%s cmd=%s)", url, user.Name, cmd)
		default:
			return fmt.Errorf("db get feed (username=%s cmd=%s): %w", user.Name, cmd, err)
		}
	}

	err = app.DBQueries.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("db delete feed (username=%s cmd=%s): %w", user.Name, cmd, err)
	}

	fmt.Printf("You are no longer following %s at %s\n", feed.Name, url)

	return nil
}
