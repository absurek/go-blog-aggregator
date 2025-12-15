package cli

import (
	"fmt"

	"github.com/absurek/go-blog-aggregator/internal/application"
)

type HandlerFunc func(*application.Application, Command) error

type Commands struct {
	handlerMap map[string]HandlerFunc
}

func NewCommands() *Commands {
	return &Commands{
		handlerMap: make(map[string]HandlerFunc),
	}
}

func (c *Commands) Run(app *application.Application, cmd Command) error {
	handler, ok := c.handlerMap[cmd.name]
	if !ok {
		return fmt.Errorf("unkown command in run (command=%s)", cmd.name)
	}

	return handler(app, cmd)
}

func (c *Commands) Register(name string, handler HandlerFunc) {
	c.handlerMap[name] = handler
}
