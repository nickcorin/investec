package cli

import (
	"fmt"
)

// NewRouter creates a new CommandRouter.
func NewRouter() *CommandRouter {
	return &CommandRouter{
		commands: make(map[string]Command),
	}
}

// CommandRouter stores the command tree.
type CommandRouter struct {
	commands map[string]Command
}

// Register adds a command to the router.
func (router *CommandRouter) Register(command string, fn CommandFn,
	help string) error {

	// Ensure that the commands map is not nil.
	if router.commands == nil {
		router.commands = make(map[string]Command)
	}

	// Check whether a command with the same command string already exists.
	if _, ok := router.commands[command]; ok {
		return fmt.Errorf("command string already exists")
	}

	// Register the command.
	router.commands[command] = Command{
		cmd:       command,
		fn:        fn,
		help:      help,
		subrouter: NewRouter(),
	}

	return nil
}

// Run traverses the router tree and executes the deepest command, passing in
// the remainder of the command string as arguments.
func (router *CommandRouter) Search(commands ...string) (*Command, error) {
	r := router
	for {
		// There are no more commands to search.
		if len(commands) == 0 {
			break
		}

		c, ok := r.commands[commands[0]]
		if !ok {
			break
		}

		if c.subrouter == nil {
			return &c, nil
		}

		r = c.subrouter
		commands = commands[1:]
	}

	return nil, fmt.Errorf("command not found")
}
