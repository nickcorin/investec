package investec

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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

// transactionWithoutDate is a type alias for Transaction allowing a custom
// Unmarshaler implementation for only the date values.
type transactionWithoutDate Transaction

// UnmarshalJSON satisfies the json.Unmarshaler interface.
func (t *Transaction) UnmarshalJSON(data []byte) error {
	raw := struct {
		transactionWithoutDate
		PostingDate     string `json:"postingDate"`
		ValueDate       string `json:"valueDate"`
		ActionDate      string `json:"actionDate"`
		TransactionDate string `json:"transactionDate"`
	}{}

	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	*t = Transaction(raw.transactionWithoutDate)

	t.PostingDate, err = time.Parse("2006-01-02", raw.PostingDate)
	if err != nil {
		return err
	}

	t.ValueDate, err = time.Parse("2006-01-02", raw.ValueDate)
	if err != nil {
		return err
	}

	t.ActionDate, err = time.Parse("2006-01-02", raw.ActionDate)
	if err != nil {
		return err
	}

	t.TransactionDate, err = time.Parse("2006-01-02", raw.TransactionDate)
	if err != nil {
		return err
	}

	return nil
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

	res, err := c.opts.Transport.Post(ctx, fmt.Sprintf(
		"/za/pb/v1/accounts/%s/transactions", req.AccountID), nil,
		bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to get account transactions: %w", err)
	}

	var transactions TransactionsResponse
	if err = res.JSON(&transactions); err != nil {
		return nil, fmt.Errorf("failed to unmarshal transactions response: %w",
			err)
	}

	return &transactions, nil
}
