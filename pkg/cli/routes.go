package cli

import (
	"fmt"
	"io"

	"github.com/nickcorin/ziggy"
)

// RegisterRoutes registers the routes for the ziggyd CLI application.
func RegisterRoutes(r *CommandRouter) {
	r.Register("accounts", accountsHandler, "some help for accounts")
}

func accountsHandler(c *ziggy.Client, w io.Writer, args ...string) error {
	w.Write([]byte("not yet implemented"))
	return fmt.Errorf("not implemented")
}
