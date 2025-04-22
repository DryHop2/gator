package commands

import (
	"context"
	"fmt"

	"github.com/DryHop2/gator/internal/state"
)

func HanlderFollowing(s *state.State, cmd Command) error {
	user, err := s.DB.GetUserByName(context.Background(), s.Cfg.CurrentUser)
	if err != nil {
		return fmt.Errorf("could not get current user: %w", err)
	}

	follows, err := s.DB.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("could not get followed feeds: %w", err)
	}

	if len(follows) == 0 {
		fmt.Println("You're not following any feeds yet.")
		return nil
	}

	fmt.Println("You are following:")
	for _, follow := range follows {
		fmt.Printf("* %s - %s\n", follow.FeedName, follow.FeedUrl)
	}

	return nil
}
