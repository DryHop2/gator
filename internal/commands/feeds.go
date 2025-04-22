package commands

import (
	"fmt"

	"github.com/DryHop2/gator/internal/state"
)

func HandlerFeeds(s *state.State, cmd Command) error {
	feeds, err := s.DB.GetAllFeedsWithUsers(s.Ctx)
	if err != nil {
		return fmt.Errorf("could not retrieve feeds: %w", err)
	}

	for _, f := range feeds {
		fmt.Printf("Feed: %s\n URL: %s\n Author: %s\n\n", f.FeedName, f.FeedUrl, f.UserName)
	}

	return nil
}
