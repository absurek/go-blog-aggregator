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

func FollowingHandler(app *application.Application, cmd cli.Command, user database.User) error {
	follows, err := app.DBQueries.GetFeedFollowsForUser(context.Background(), user.Name)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return fmt.Errorf("no follows found for user %s", user.Name)
		default:
			return fmt.Errorf("get follows (username=%s cmd=%s): %w", user.Name, cmd, err)
		}
	}

	fmt.Println()
	fmt.Println("Feeds you are following:")
	for _, follow := range follows {
		fmt.Println("-", follow.FeedName)
	}
	fmt.Println()

	return nil
}
