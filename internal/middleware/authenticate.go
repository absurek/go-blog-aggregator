package middleware

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/absurek/go-blog-aggregator/internal/application"
	"github.com/absurek/go-blog-aggregator/internal/cli"
)

func Authenticate(handler application.AuthenticatedHandlerFunc) application.HandlerFunc {
	return func(app *application.Application, cmd cli.Command) error {
		username := app.Config.CurrentUsername
		if username == "" {
			return errors.New("login required")
		}

		user, err := app.DBQueries.GetUser(context.Background(), username)
		if err != nil {
			switch {
			case errors.Is(err, sql.ErrNoRows):
				return fmt.Errorf("no user with name '%s'", username)
			default:
				return fmt.Errorf("get user (username=%s cmd=%s): %w", username, cmd, err)
			}
		}

		return handler(app, cmd, user)
	}
}
