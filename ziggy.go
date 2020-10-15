package ziggy

import (
	"encoding/json"
	"time"
)

// AccessToken contains a temporary authorization key.
type AccessToken struct {
	Token     string     `json:"access_token"`
	Type      TokenType  `json:"token_type"`
	ExpiresIn int64      `json:"expires_in"`
	Scope     TokenScope `json:"scope"`

	// createdAt is the timestamp of when this token was received. It allows us
	// to calculate whether the token has expired. It is possible that the delay
	// between Intvestec creating the token and the client parsing the data can
	// cause the expiration time to be fuzzy.
	CreatedAt time.Time
}

// IsExpired returns whether the token has expired.
func (token *AccessToken) IsExpired() bool {
	t := time.Now().Add(time.Second * time.Duration(token.ExpiresIn))
	return time.Now().After(t)
}

// Account contains information regarding a single bank account.
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

// Balance contains a breakdown of balances for a single account.
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

// TokenScope describes the the scope of an access token.
type TokenScope string

const (
	TokenScopeAccounts TokenScope = "accounts"
)

// TokenType describes the kind of an access token.
type TokenType string

const (
	TokenTypeBearer TokenType = "Bearer"
)

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
