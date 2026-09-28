package main

import (
	"context"
	"fmt"
	"time"

	"github.com/cjvnjde/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("url is required")
	}

	feedURL := cmd.args[0]

	feed, err := s.db.GetFeed(context.Background(), feedURL)
	if err != nil {
		return err
	}

	follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return err
	}

	fmt.Printf("feed: %s\nuser: %s", follow.Feed.Name, follow.User.Name)

	return nil
}
