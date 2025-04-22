package commands

import (
	"context"
	"fmt"

	"github.com/DryHop2/gator/internal/database"
	"github.com/DryHop2/gator/internal/state"
)

func MiddlewareLoggedIn(
	handler func(s *state.State, cmd Command, user database.User) error,
) func(s *state.State, cmd Command) error {
	return func(s *state.State, cmd Command) error {
		if s.Cfg.CurrentUser == "" {
			return fmt.Errorf("no user is currently logged in")
		}

		user, err := s.DB.GetUserByName(context.Background(), s.Cfg.CurrentUser)
		if err != nil {
			return fmt.Errorf("failed to fetch user: %w", err)
		}

		return handler(s, cmd, user)
	}
}
