package cli

import (
	"errors"
	"fmt"

	"github.com/absurek/go-blog-aggregator/internal/application"
)

func LoginHandler(app *application.Application, cmd Command) error {
	if len(cmd.args) != 1 {
		return errors.New("login expects a single argument")
	}

	username := cmd.args[0]
	err := app.Config.SetUser(username)
	if err != nil {
		return fmt.Errorf("login set user (username=%s), %w", username, err)
	}

	fmt.Printf("User set to: %s\n", username)

	return nil
}
