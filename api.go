package investec

import "context"

type Client interface {
	// AccessToken obtains an access token.
	GetAccessToken(ctx context.Context, scope TokenScope) (*AccessToken, error)

	// GetAccountTransactions obtains a specified account's transactions.
	GetAccountTransactions(ctx context.Context, req *TransactionsRequest) (
		*TransactionsResponse, error)
}
