package cli

import (
	"errors"
)

type Command struct {
	name string
	args []string
}

func ParseCommand(rawArgs []string) (*Command, error) {
	if len(rawArgs) < 2 {
		return nil, errors.New("insufficient number of arguments")
	}

	return &Command{
		name: rawArgs[1],
		args: rawArgs[2:],
	}, nil
}

func (c Command) String() string {
	return c.name
}
