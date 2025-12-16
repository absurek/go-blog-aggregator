package cli

import (
	"errors"
)

type Command struct {
	Name string
	Args []string
}

func ParseCommand(rawArgs []string) (*Command, error) {
	if len(rawArgs) < 2 {
		return nil, errors.New("insufficient number of arguments")
	}

	return &Command{
		Name: rawArgs[1],
		Args: rawArgs[2:],
	}, nil
}

func (c Command) String() string {
	return c.Name
}
