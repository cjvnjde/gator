package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Printf("name: %s\nurl: %s\ncreated by: %s\n", feed.Name, feed.Url, feed.User.Name)
	}
	return nil
}
