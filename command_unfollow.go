package main

import (
	"context"
	"fmt"

	"github.com/cjvnjde/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("login needs a username")
	}

	feedURL := cmd.args[0]

	feed, err := s.db.GetFeed(context.Background(), feedURL)
	if err != nil {
		return err
	}

	return s.db.Unfollow(context.Background(), database.UnfollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
}
