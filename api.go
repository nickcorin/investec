package investec

import "context"

type Client interface {
	// GetAccountTransactions obtains a specified account's transactions.
	GetAccountTransactions(ctx context.Context, req *TransactionsRequest) (
		*TransactionsResponse, error)
}
