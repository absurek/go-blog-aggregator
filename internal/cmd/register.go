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

func RegisterHandler(app *application.Application, cmd cli.Command) error {
	if len(cmd.Args) != 1 {
		return errors.New("register expects a single argument")
	}

	user, err := app.DBQueries.CreateUser(context.Background(), database.CreateUserParams{
		ID:   uuid.New(),
		Name: cmd.Args[0],
	})
	if err != nil {
		return fmt.Errorf("db create user: %w", err)
	}

	app.Config.SetUser(user.Name)
	fmt.Println("User created!")
	fmt.Println(user)

	return nil
}
