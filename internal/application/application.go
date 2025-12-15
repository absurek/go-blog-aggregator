package application

import "github.com/absurek/go-blog-aggregator/internal/config"

type Application struct {
	Config *config.Config
}

func NewApplication(config *config.Config) *Application {
	return &Application{
		Config: config,
	}
}
