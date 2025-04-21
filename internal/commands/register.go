package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/DryHop2/gator/internal/database"
	"github.com/DryHop2/gator/internal/state"
	"github.com/google/uuid"
)

func HandlerRegister(s *state.State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("a username is required")
	}

	username := cmd.Args[0]

	user, err := s.DB.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	})

	if err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}

	if err := s.Cfg.SetUser(username); err != nil {
		return fmt.Errorf("failed to update config: %w", err)
	}

	fmt.Printf("User %s created successfully\n", username)
	fmt.Printf("Debug: %+v\n", user)

	return nil
}
