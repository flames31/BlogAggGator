package main

import "errors"

type command struct {
	Name string
	Args []string
}

type commands struct {
	regsiteredCommands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.regsiteredCommands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.regsiteredCommands[cmd.Name]
	if !ok {
		return errors.New("command not found brotha")
	}
	return f(s, cmd)
}
