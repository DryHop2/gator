package commands

import (
	"context"
	"fmt"

	"github.com/DryHop2/gator/internal/rss"
	"github.com/DryHop2/gator/internal/state"
)

func HandlerAgg(s *state.State, cmd Command) error {
	const feedURL = "https://www.wagslane.dev/index.xml"

	feed, err := rss.FetchFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("failed to fetch feed: %w", err)
	}

	// fmt.Printf("Feed Title: %s\n", feed.Channel.Title)
	// fmt.Printf("Feed Description: %s\n", feed.Channel.Description)
	// fmt.Println("Articles:")
	// for _, item := range feed.Channel.Item {
	// 	fmt.Printf("- %s (%s)\n", item.Title, item.PubDate)
	// }

	fmt.Printf("%+v\n", feed)

	return nil
}
