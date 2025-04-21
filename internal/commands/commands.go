package commands

import (
	"fmt"

	"github.com/DryHop2/gator/internal/state"
)

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	handlers map[string]func(*state.State, Command) error
}

func New() *Commands {
	return &Commands{
		handlers: make(map[string]func(*state.State, Command) error),
	}
}

func (c *Commands) Register(name string, handler func(*state.State, Command) error) {
	c.handlers[name] = handler
}

func (c *Commands) Run(s *state.State, cmd Command) error {
	handler, ok := c.handlers[cmd.Name]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.Name)
	}
	return handler(s, cmd)
}
