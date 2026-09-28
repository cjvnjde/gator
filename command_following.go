package main

import (
	"context"
	"fmt"

	"github.com/cjvnjde/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	followings, err := s.db.GetFeedFollowForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, following := range followings {
		fmt.Println(following.Feed.Name)
	}

	return nil
}
