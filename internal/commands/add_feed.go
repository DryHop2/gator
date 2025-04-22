package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/DryHop2/gator/internal/database"
	"github.com/DryHop2/gator/internal/state"
	"github.com/google/uuid"
)

func HandlerAddFeed(s *state.State, cmd Command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("usage: addfeed <name> <url>")
	}

	feedName := cmd.Args[0]
	feedURL := cmd.Args[1]

	user, err := s.DB.GetUserByName(context.Background(), s.Cfg.CurrentUser)
	if err != nil {
		return fmt.Errorf("could not find current user: %w", err)
	}

	newFeed, err := s.DB.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       feedURL,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to create feed: %w", err)
	}

	fmt.Println("Feed added!")
	fmt.Printf("ID: %s\n", newFeed.ID)
	fmt.Printf("Name: %s\n", newFeed.Name)
	fmt.Printf("URL: %s\n", newFeed.Url)
	fmt.Printf("User ID: %s\n", newFeed.UserID)

	return nil
}
