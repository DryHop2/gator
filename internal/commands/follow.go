package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/DryHop2/gator/internal/database"
	"github.com/DryHop2/gator/internal/state"
	"github.com/google/uuid"
)

func HandlerFollow(s *state.State, cmd Command, user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("you must provide a feed url")
	}

	url := cmd.Args[0]

	// user, err := s.DB.GetUserByName(context.Background(), s.Cfg.CurrentUser)
	// if err != nil {
	// 	return fmt.Errorf("failed to find current user: %w", err)
	// }

	feed, err := s.DB.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("failed to find feed: %w", err)
	}

	follow, err := s.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})

	if err != nil {
		return fmt.Errorf("failed to follow feed: %w", err)
	}

	fmt.Printf("User %s is now following feed: %s\n", follow.UserName, follow.FeedName)
	return nil
}
