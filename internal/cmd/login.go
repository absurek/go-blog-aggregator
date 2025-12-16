package cmd

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/absurek/go-blog-aggregator/internal/application"
	"github.com/absurek/go-blog-aggregator/internal/cli"
)

func LoginHandler(app *application.Application, cmd cli.Command) error {
	if len(cmd.Args) != 1 {
		return errors.New("login expects a single argument")
	}

	username := cmd.Args[0]
	_, err := app.DBQueries.GetUser(context.Background(), username)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return fmt.Errorf("unkown user (name=%s)", username)
		default:
			return fmt.Errorf("db get user (name=%s): %w", username, err)
		}
	}

	err = app.Config.SetUser(username)
	if err != nil {
		return fmt.Errorf("set user (username=%s), %w", username, err)
	}

	fmt.Printf("User set to: %s\n", username)

	return nil
}
