package commands

import (
	"fmt"

	"github.com/DryHop2/gator/internal/state"
)

func HandlerLogin(s *state.State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("a username is required")
	}

	username := cmd.Args[0]
	err := s.Cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("failed to set user: %w", err)
	}

	fmt.Printf("Current user set to %s\n", username)
	return nil
}
