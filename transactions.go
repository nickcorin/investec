package investec

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

// TransactionStatus describes the current state of a transaction.
type TransactionStatus string

const (
	TransactionStatusPosted = "POSTED"
)

// TransactionType describes the kind of the transaction.
type TransactionType string

const (
	TransactionTypeCredit = "CREDIT"
	TransactionTypeDebit  = "DEBIT"
)

// Transaction describes some activity which has occurred on an account.
type Transaction struct {
	AccountID       string            `json:"accountId"`
	Type            TransactionType   `json:"type"`
	Status          TransactionStatus `json:"status"`
	Description     string            `json:"description"`
	CardNumber      string            `json:"cardNumber"`
	PostingDate     time.Time         `json:"postingDate"`
	ValueDate       time.Time         `json:"valueDate"`
	ActionDate      time.Time         `json:"actionDate"`
	TransactionDate time.Time         `json:"transactionDate"`
	Amount          float64           `json:"amount"`
}

// TransactionRequest describes the request parameters available for retrieving
// account transactions. All parameters are requires unless specified.
type TransactionsRequest struct {
	AccountID string `json:"-"`

	// Optional.
	StartDate time.Time `json:"fromDate,omitempty"`

	// Optional.
	EndDate time.Time `json:"toDate,omitempty"`
}

// TransactionsResponse describes the response data returned for retrieving
// account transactions.
type TransactionsResponse struct {
	Data struct {
		Transactions []Transaction `json:"transactions"`
	} `json:"data"`

	Links struct {
		Self string `json:"self"`
	} `json:"links"`

	Metadata struct {
		TotalPages int64 `json:"totalPages"`
	} `json:"meta"`
}

// GetAccountTransactions obtains a specified account's transactions.
func (c *client) GetAccountTransactions(ctx context.Context,
	req *TransactionsRequest) (*TransactionsResponse, error) {

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	_, body, err := c.post(ctx, fmt.Sprintf(
		"/za/pb/v1/accounts/%s/transactions", req.AccountID), nil,
		bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to get account transactions: %w", err)
	}

	data, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var transactions TransactionsResponse
	if err = json.Unmarshal(data, &transactions); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &transactions, nil
}
