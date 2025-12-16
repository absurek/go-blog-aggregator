package cmd

import (
	"context"
	"fmt"

	"github.com/absurek/go-blog-aggregator/internal/application"
	"github.com/absurek/go-blog-aggregator/internal/cli"
)

func UsersHandler(app *application.Application, cmd cli.Command) error {
	users, err := app.DBQueries.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("db get users: %w", err)
	}

	for _, user := range users {
		if app.Config.CurrentUsername == user.Name {
			fmt.Printf("%s (current)\n", user.Name)
		} else {
			fmt.Println(user.Name)
		}
	}

	return nil
}
