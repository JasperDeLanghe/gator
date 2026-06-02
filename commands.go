package main

import "errors"

type command struct {
	Name string
	Args []string
}

// This will hold all the commands the CLI can handle
type commands struct {
	handlers map[string]func(*state, command) error
}

// This method runs a given command with the provided state if it exists.
func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.handlers[cmd.Name]
	if !ok {
		return errors.New("command not found in registry")
	}

	return handler(s, cmd)
}

// This method registers a new handler function for a command name.
func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}
