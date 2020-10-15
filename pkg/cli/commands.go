package cli

import (
	"io"
	"strings"

	"github.com/nickcorin/ziggy"
)

// CommandFn describes the signature for an executable ziggyd command. It takes
// in a client, a Writer to which all output will be streamed, and the arguments
// for the command.
type CommandFn func(c *ziggy.Client, w io.Writer, args ...string) error

// Command describes a node in the command tree.
type Command struct {
	cmd       string
	fn        CommandFn
	help      string
	subrouter *CommandRouter
}

// Run executes the command with the provided arguments.
func (c *Command) Run(client *ziggy.Client, w io.Writer, args ...string) error {
	newArgs := make([]string, 0)
	for i, arg := range args {
		if !strings.EqualFold(c.cmd, arg) {
			continue
		}
		newArgs = append(newArgs, args[i:]...)
		break
	}

	return c.fn(client, w, newArgs...)
}

// RegisterSubcommand adds a command to the Command's subrouter.
func (c *Command) RegisterSubcommand(command string, fn CommandFn,
	help string) error {
	if c.subrouter == nil {
		c.subrouter = NewRouter()
	}

	return c.subrouter.Register(command, fn, help)
}

func (c *Command) printHelp(w io.Writer) {
	w.Write([]byte(c.help))
}
