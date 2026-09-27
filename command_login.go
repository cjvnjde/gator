package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("login needs a username")
	}

	username := cmd.args[0]
	err := s.config.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("Username %s has been set", username)

	return nil
}
