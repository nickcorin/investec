package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/nickcorin/ziggy"
)

// GetTransactions satisfies the ziggy.Client interface.
func (c *httpClient) GetTransactions(ctx context.Context,
	req *ziggy.TransactionsRequest) ([]ziggy.Transaction, error) {

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	res, err := c.transport.Post(ctx, fmt.Sprintf(
		"/za/pb/v1/accounts/%s/transactions", req.AccountID), nil,
		bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to get account transactions: %w", err)
	}

	var transactions ziggy.TransactionsResponse
	if err = res.JSON(&transactions); err != nil {
		return nil, fmt.Errorf("failed to unmarshal transactions response: %w",
			err)
	}

	return transactions.Data.Transactions, nil
}
