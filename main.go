package main

import (
	"fmt"
	"os"

	"github.com/absurek/go-blog-aggregator/internal/application"
	"github.com/absurek/go-blog-aggregator/internal/cli"
	"github.com/absurek/go-blog-aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error: could not read config: %v\n", err)
		os.Exit(1)
	}

	app := application.NewApplication(cfg)
	commands := cli.NewCommands()

	commands.Register("login", cli.LoginHandler)

	cmd, err := cli.ParseCommand(os.Args)
	if err != nil {
		fmt.Printf("Error: invalid command line arguments: %v\n", err)
		os.Exit(1)
	}

	err = commands.Run(app, *cmd)
	if err != nil {
		fmt.Printf("Error: command %s: %v\n", cmd, err)
		os.Exit(1)
	}
}
