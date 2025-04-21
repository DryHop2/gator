package commands

import (
	"context"
	"fmt"

	"github.com/DryHop2/gator/internal/state"
)

func HandlerReset(s *state.State, cmd Command) error {
	err := s.DB.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to reset users: %w", err)
	}

	fmt.Println("All users have been deleted. Database reset complete.")
	return nil
}
