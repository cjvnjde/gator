package main

import (
	"context"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("login needs a username")
	}

	username := cmd.args[0]

	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return err
	}

	err = s.config.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Printf("Username %s has been set", username)

	return nil
}
