package commands

import (
	"fmt"
	"time"

	"github.com/DryHop2/gator/internal/state"
)

func HandlerAgg(s *state.State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: agg <duration>")
	}

	duration, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}

	fmt.Printf("Collecting feeds every %s\n", duration)

	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	scrapeFeeds(s)

	for range ticker.C {
		scrapeFeeds(s)
	}

	return nil
}
