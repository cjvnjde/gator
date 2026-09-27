package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/cjvnjde/gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("register needs a username")
	}

	username := cmd.args[0]

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	})
	if err != nil {
		return err
	}

	err = s.config.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Printf(`User has been created:
	- UUID: %s,
	- Name: %s, 
	- Created At: %v,
	- Updated At: %v`, user.ID, user.Name, user.CreatedAt, user.UpdatedAt)

	return nil
}
