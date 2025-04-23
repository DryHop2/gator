package commands

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/DryHop2/gator/internal/database"
	"github.com/DryHop2/gator/internal/rss"
	"github.com/DryHop2/gator/internal/state"
	"github.com/google/uuid"
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

	layout := "Mon, 02 Jan 2006 15:04:05 MST"

	for _, item := range parsedFeed.Channel.Item {
		pubTime, err := time.Parse(layout, item.PubDate)
		if err != nil {
			fmt.Println("Could not parse publication time:", err)
			continue
		}

		err = s.DB.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   now,
			UpdatedAt:   now,
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: item.Description != ""},
			PublishedAt: sql.NullTime{Time: pubTime, Valid: true},
			FeedID:      feed.ID,
		})

		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			fmt.Println("Failed to save posit:", err)

		}
	}
}
