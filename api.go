package investec

import "context"

type Client interface {
	// AccessToken obtains an access token.
	GetAccessToken(ctx context.Context, scope TokenScope) (*AccessToken, error)

	// GetAccounts obtains a list of accounts.
	GetAccounts(ctx context.Context) ([]Account, error)

	// GetAccountBalance obtains a specified account's balance.
	GetAccountBalance(ctx context.Context, accountID string) (*Balance, error)

	// GetAccountTransactions obtains a specified account's transactions.
	GetAccountTransactions(ctx context.Context, req *TransactionsRequest) (
		*TransactionsResponse, error)
}
