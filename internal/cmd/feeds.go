package cmd

import (
	"context"
	"fmt"

	"github.com/absurek/go-blog-aggregator/internal/application"
	"github.com/absurek/go-blog-aggregator/internal/cli"
)

func FeedsHandler(app *application.Application, cmd cli.Command) error {
	feeds, err := app.DBQueries.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("db get feeds (cmd=%s): %w", cmd, err)
	}

	fmt.Println()
	fmt.Println("Feeds:")
	for _, feed := range feeds {
		fmt.Println()
		fmt.Println("- name:", feed.Name)
		fmt.Println("- url:", feed.Url)
		fmt.Println("- user:", feed.UserName)
	}
	fmt.Println()

	return nil
}
