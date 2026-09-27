package main

import (
	"fmt"
)

type command struct {
	name string
	args []string
}

type commands struct {
	commands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	currentCommand, ok := c.commands[cmd.name]

	if !ok {
		return fmt.Errorf("command %s doesn't exist", cmd.name)
	}

	return currentCommand(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commands[name] = f
}
