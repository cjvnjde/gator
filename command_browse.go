package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cjvnjde/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	var limit int32 = 2

	if len(cmd.args) > 0 {
		i, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			limit = int32(i)
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	})
	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Println(post.Title)
	}

	return nil
}
