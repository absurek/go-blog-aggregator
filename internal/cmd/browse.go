package cmd

import (
	"context"
	"fmt"
	"strconv"

	"github.com/absurek/go-blog-aggregator/internal/application"
	"github.com/absurek/go-blog-aggregator/internal/cli"
	"github.com/absurek/go-blog-aggregator/internal/database"
	"github.com/absurek/go-blog-aggregator/internal/utils"
)

func BrowseHandler(app *application.Application, cmd cli.Command, user database.User) error {
	limit := int32(2)
	if len(cmd.Args) == 1 {
		l, err := strconv.ParseInt(cmd.Args[0], 10, 32)
		if err == nil {
			limit = int32(l)
		}
	}

	posts, err := app.DBQueries.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		Name:  user.Name,
		Limit: limit,
	})
	if err != nil {
		return fmt.Errorf("get posts for user (username=%s): %w", user.Name, err)
	}

	fmt.Println()
	fmt.Println("Posts:")
	fmt.Println()
	for _, post := range posts {
		fmt.Println("Title", post.Title.String)
		fmt.Println("Published At", utils.FormatNullTime(post.PublishedAt))
		fmt.Println()
	}

	return nil
}
