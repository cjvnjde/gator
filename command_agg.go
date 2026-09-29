package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/cjvnjde/gator/internal/database"
	"github.com/google/uuid"
)

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return nil
	}
	_, err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return nil
	}
	data, err := fetchFeed(context.Background(), nextFeed.Url)

	for _, item := range data.Channel.Item {
		t, err := time.Parse(time.RFC3339, item.PubDate)
		if err != nil {
			fmt.Println(err)
			t = time.Now()
		}

		s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:    uuid.New(),
			Title: item.Title,
			Url:   item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid:  item.Description != "",
			},
			PublishedAt: t,
			FeedID:      nextFeed.ID,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		})
		fmt.Println(item.Title)
	}

	return nil
}

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("Must hame time between requests")
	}

	timeBetweenReqs := cmd.args[0]

	timeBetweenRequests, err := time.ParseDuration(timeBetweenReqs)
	if err != nil {
		return err
	}
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}
