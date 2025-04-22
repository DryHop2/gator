package commands

import (
	"context"
	"fmt"

	"github.com/DryHop2/gator/internal/state"
)

func HandlerUsers(s *state.State, cmd Command) error {
	users, err := s.DB.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("could not fetch users: %w", err)
	}

	current := s.Cfg.CurrentUser

	for _, u := range users {
		if u.Name == current {
			fmt.Printf("* %s (current)\n", u.Name)
		} else {
			fmt.Printf("* %s\n", u.Name)
		}
	}

	return nil
}
