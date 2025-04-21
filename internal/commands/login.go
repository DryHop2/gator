package commands

import (
	"context"
	"fmt"
	"os"

	"github.com/DryHop2/gator/internal/state"
)

func HandlerLogin(s *state.State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("a username is required")
	}

	username := cmd.Args[0]

	_, err := s.DB.GetUserByName(context.Background(), username)
	if err != nil {
		fmt.Printf("user %s does not exist\n", username)
		os.Exit(1)
	}

	err = s.Cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("failed to set user: %w", err)
	}

	fmt.Printf("Current user set to %s\n", username)
	return nil
}
