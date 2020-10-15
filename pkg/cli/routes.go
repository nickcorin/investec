package cli

import (
	"context"
	"fmt"
	"io"

	"github.com/nickcorin/ziggy"
)

// RegisterRoutes registers the routes for the ziggyd CLI application.
func RegisterRoutes(r *CommandRouter) {
	r.Register("auth", tokenHandler, "refreshes your ziggy access token")
	r.Register("accounts", accountsHandler, "fetches a list of accounts")
}

func tokenHandler(ctx context.Context, c ziggy.Client, w io.Writer,
	args ...string) error {
	_, err := c.GetAccessToken(ctx, ziggy.TokenScopeAccounts)
	if err != nil {
		fmt.Fprintf(w, "failed to fetch access token :(")
		return nil
	}

	fmt.Fprintf(w, "Successfully authenticated.")

	return nil
}

func accountsHandler(ctx context.Context, c ziggy.Client, w io.Writer,
	args ...string) error {
	_, err := c.GetAccessToken(ctx, ziggy.TokenScopeAccounts)
	if err != nil {
		fmt.Fprintf(w, "failed to fetch access token :(")
		return nil
	}

	accounts, err := c.GetAccounts(ctx)
	if err != nil {
		fmt.Fprintf(w, "failed to fetch accounts :(")
	}

	if len(accounts) == 0 {
		fmt.Fprintf(w, "you have no accounts")
		return nil
	}

	for _, acc := range accounts {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t\n", acc.ID, acc.Name, acc.Number,
		acc.Product, acc.Reference,)
	}

	return nil
}
