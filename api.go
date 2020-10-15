package ziggy

import (
	"context"
)

// Client describes the API for the Investec OpenAPI.
type Client interface {
	// GetAccessToken obtains an access token for a provided access scope.
	GetAccessToken(ctx context.Context, scope TokenScope) (*AccessToken, error)

	// GetAccounts obtains a list of accounts.
	GetAccounts(ctx context.Context) ([]Account, error)

	// GetTransactions obtains a list of transactions for a provided account id.
	GetTransactions(ctx context.Context, r *TransactionsRequest) ([]Transaction,
		error)

	// GetAccountBalance obtains the balance for a provided account id.
	GetAccountBalance(ctx context.Context, accountID string) (*Balance, error)
}
