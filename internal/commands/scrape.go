package commands

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/DryHop2/gator/internal/database"
	"github.com/DryHop2/gator/internal/rss"
	"github.com/DryHop2/gator/internal/state"
)

func scrapeFeeds(s *state.State) {
	ctx := context.Background()

	feed, err := s.DB.GetNextFeedToFetch(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No feeds are ready for fetching.")
			return
		}
		fmt.Println("Error getting feed:", err)
		return
	}

	now := time.Now()
	err = s.DB.MarkFeedFetched(ctx, database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{Time: now, Valid: true},
		ID:            feed.ID,
	})
	if err != nil {
		fmt.Println("Failed to mark feed as fetched:", err)
		return
	}

	fmt.Printf("Fetching: %s (%s)\n", feed.Name, feed.Url)

	parsedFeed, err := rss.FetchFeed(ctx, feed.Url)
	if err != nil {
		fmt.Println("Error fetching/parsing feed:", err)
		return
	}

	for _, item := range parsedFeed.Channel.Item {
		fmt.Printf("- %s (%s)\n", item.Title, item.PubDate)
	}
}
