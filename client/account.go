package client

import (
	"context"
	"fmt"

	"github.com/nickcorin/ziggy"
)

// GetAccounts satisfies the ziggy.Client interface.
func (c *httpClient) GetAccounts(ctx context.Context) ([]ziggy.Account, error) {
	res, err := c.transport.Get(ctx, "/za/pb/v1/accounts", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get accounts: %w", err)
	}

	var accountsResponse ziggy.AccountResponse
	if err = res.JSON(&accountsResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal account response: %w", err)
	}

	return accountsResponse.Data.Accounts, nil
}

// GetAccountBalance satisfies the ziggy.Client interface.
func (c *httpClient) GetAccountBalance(ctx context.Context, accountID string) (
	*ziggy.Balance, error) {
	res, err := c.transport.Post(ctx, fmt.Sprintf(
		"/za/pb/v1/accounts/%s/balance", accountID), nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get account balance %s: %w",
			accountID, err)
	}

	var balanceResponse ziggy.BalanceResponse
	if err = res.JSON(&balanceResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal balance response: %w", err)
	}

	return &balanceResponse.Data.Balance, nil
}
