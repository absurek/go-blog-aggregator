package cmd

import (
	"context"
	"fmt"

	"github.com/absurek/go-blog-aggregator/internal/application"
	"github.com/absurek/go-blog-aggregator/internal/cli"
)

func ResetHandler(app *application.Application, cmd cli.Command) error {
	err := app.DBQueries.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("db delete all users: %w", err)
	}

	fmt.Println("Users table reset!")

	return nil
}
