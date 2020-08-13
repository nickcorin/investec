package investec

import (
	"context"
	"fmt"
)

type Account struct {
	ID        string `json:"accountId"`
	Name      string `json:"accountName"`
	Number    string `json:"accountNumber"`
	Product   string `json:"productName"`
	Reference string `json:"referenceName"`
}

// AccountResponse describes the response data returned for retrieving a list
// of accounts.
type AccountResponse struct {
	Data struct {
		Accounts []Account `json:"accounts"`
	} `json:"data"`

	Links struct {
		Self string `json:"self"`
	} `json:"links"`

	Metadata struct {
		TotalPages int64 `json:"totalPages"`
	} `json:"meta"`
}

// GetAccounts obtains a list of accounts.
func (c *client) GetAccounts(ctx context.Context) ([]Account, error) {
	res, err := c.transport.Get(ctx, "/za/pb/v1/accounts", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get accounts: %w", err)
	}

	var accountsResponse AccountResponse
	if err = res.JSON(&accountsResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal account response: %w", err)
	}

	return accountsResponse.Data.Accounts, nil
}

type Balance struct {
	AccountID string  `json:"accountId"`
	Available float64 `json:"availableBalance"`
	Currency  string  `json:"currency"`
	Current   float64 `json:"currentBalance"`
}

// BalanceResponse describes the response data returned for retrieving an
// account's balance.
type BalanceResponse struct {
	Data struct {
		Balance
	} `json:"data"`

	Links struct {
		Self string `json:"self"`
	} `json:"links"`

	Metadata struct {
		TotalPages int64 `json:"totalPages"`
	} `json:"meta"`
}

// GetAccountBalance obtains a specified account's balance.
func (c *client) GetAccountBalance(ctx context.Context, accountID string) (
	*Balance, error) {
	res, err := c.transport.Post(ctx, fmt.Sprintf(
		"/za/pb/v1/accounts/%s/balance", accountID), nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get account balance %s: %w",
			accountID, err)
	}

	var balanceResponse BalanceResponse
	if err = res.JSON(&balanceResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal balance response: %w", err)
	}

	return &balanceResponse.Data.Balance, nil
}
