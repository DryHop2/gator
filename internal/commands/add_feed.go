package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/DryHop2/gator/internal/database"
	"github.com/DryHop2/gator/internal/state"
	"github.com/google/uuid"
)

func HandlerAddFeed(s *state.State, cmd Command, user database.User) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("usage: addfeed <name> <url>")
	}

	feedName := cmd.Args[0]
	feedURL := cmd.Args[1]

	// user, err := s.DB.GetUserByName(context.Background(), s.Cfg.CurrentUser)
	// if err != nil {
	// 	return fmt.Errorf("could not find current user: %w", err)
	// }

	now := time.Now()

	feed, err := s.DB.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      feedName,
		Url:       feedURL,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to create feed: %w", err)
	}

	fmt.Printf("Feed created:\n- Name: %s\n\n- URL: %s\n", feed.Name, feed.Url)

	follow, err := s.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("could not follow feed: %w", err)
	}

	fmt.Printf("Now following feed \"%s\" as user \"%s\".\n", follow.FeedName, follow.UserName)

	return nil
}
