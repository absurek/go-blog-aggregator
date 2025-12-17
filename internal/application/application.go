package application

import (
	"database/sql"
	"fmt"

	"github.com/absurek/go-blog-aggregator/internal/cli"
	"github.com/absurek/go-blog-aggregator/internal/config"
	"github.com/absurek/go-blog-aggregator/internal/database"
)

type HandlerFunc func(*Application, cli.Command) error
type AuthenticatedHandlerFunc func(*Application, cli.Command, database.User) error

type Application struct {
	Config    *config.Config
	DBQueries *database.Queries
	DB        *sql.DB

	handlerMap map[string]HandlerFunc
}

func NewApplication() (*Application, error) {
	cfg, err := config.Read()
	if err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	db, err := sql.Open("postgres", cfg.DataBaseURL)
	if err != nil {
		return nil, fmt.Errorf("connect to database %s: %w", cfg.DataBaseURL, err)
	}

	dbQueries := database.New(db)

	return &Application{
		Config:     cfg,
		DBQueries:  dbQueries,
		DB:         db,
		handlerMap: make(map[string]HandlerFunc),
	}, nil
}

func (a *Application) Execute(cmd cli.Command) error {
	handler, ok := a.handlerMap[cmd.Name]
	if !ok {
		return fmt.Errorf("unkown command: %s", cmd.Name)
	}

	return handler(a, cmd)
}

func (a *Application) RegisterCommand(name string, handler HandlerFunc) {
	a.handlerMap[name] = handler
}

func (a *Application) Close() {
	a.DB.Close()
}
