package investec

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	_, body, err := c.get(ctx, "/za/pb/v1/accounts", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get accounts: %w", err)
	}

	data, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var accountsResponse AccountResponse
	if err = json.Unmarshal(data, &accountsResponse); err != nil {
		return nil, fmt.Errorf("failed to parse response body: %w", err)
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
	_, body, err := c.post(ctx, fmt.Sprintf("/za/pb/v1/accounts/%s/balance",
		accountID), nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get account balance %s: %w",
			accountID, err)
	}

	data, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var balanceResponse BalanceResponse
	if err = json.Unmarshal(data, &balanceResponse); err != nil {
		return nil, fmt.Errorf("failed to parse response body: %w", err)
	}

	return &balanceResponse.Data.Balance, nil
}
