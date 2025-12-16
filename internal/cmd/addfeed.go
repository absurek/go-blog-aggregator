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

func AddfeedHandler(app *application.Application, cmd cli.Command) error {
	if len(cmd.Args) != 2 {
		return errors.New("addfeed expects exactly 2 arguments")
	}

	name := cmd.Args[0]
	url := cmd.Args[1]
	currentUsername := app.Config.CurrentUsername

	user, err := app.DBQueries.GetUser(context.Background(), currentUsername)
	if err != nil {
		return fmt.Errorf("db get user (username=%s feed=%s url=%s): %w", currentUsername, name, url, err)
	}

	feed, err := app.DBQueries.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:     uuid.New(),
		UserID: user.ID,
		Name:   name,
		Url:    url,
	})
	if err != nil {
		return fmt.Errorf("db create feed (username=%s feed=%s url=%s): %w", currentUsername, name, url, err)
	}

	fmt.Println("New feed created:")
	fmt.Println("- id:", feed.ID)
	fmt.Println("- user_id:", feed.UserID)
	fmt.Println("- name:", feed.Name)
	fmt.Println("- url:", feed.Url)

	return nil
}
