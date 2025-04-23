package commands

import (
	"context"
	"fmt"
	"strconv"

	"github.com/DryHop2/gator/internal/database"
	"github.com/DryHop2/gator/internal/state"
)

func HandlerBrowse(s *state.State, cmd Command, user database.User) error {
	limit := 2
	if len(cmd.Args) >= 1 {
		l, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("invalid limit value: %v", err)
		}
		limit = l
	}

	posts, err := s.DB.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("failed to fetch posts: %w", err)
	}

	if len(posts) == 0 {
		fmt.Println("No posts available.")
		return nil
	}

	for _, p := range posts {
		fmt.Println("Title:			", p.Title)
		fmt.Println("Description:	", p.Description)
		fmt.Println("Published:		", p.PublishedAt.Time)
		fmt.Println("URL:			", p.Url)
		fmt.Println()
	}

	return nil
}
