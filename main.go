package main

import (
	"fmt"
	"os"

	"github.com/absurek/go-blog-aggregator/internal/application"
	"github.com/absurek/go-blog-aggregator/internal/cli"
	"github.com/absurek/go-blog-aggregator/internal/cmd"
	_ "github.com/lib/pq"
)

func main() {
	app, err := application.NewApplication()
	if err != nil {
		fmt.Printf("ERROR: Application startup: %v\n", err)
		os.Exit(1)
	}
	defer app.Close()

	app.RegisterCommand("login", cmd.LoginHandler)
	app.RegisterCommand("register", cmd.RegisterHandler)
	app.RegisterCommand("reset", cmd.ResetHandler)
	app.RegisterCommand("users", cmd.UsersHandler)
	app.RegisterCommand("agg", cmd.AggHandler)
	app.RegisterCommand("addfeed", cmd.AddfeedHandler)
	app.RegisterCommand("feeds", cmd.FeedsHandler)

	cmd, err := cli.ParseCommand(os.Args)
	if err != nil {
		fmt.Printf("Error: invalid command line arguments: %v\n", err)
		os.Exit(1)
	}

	err = app.Execute(*cmd)
	if err != nil {
		fmt.Printf("Error: command %s: %v\n", cmd, err)
		os.Exit(1)
	}
}
