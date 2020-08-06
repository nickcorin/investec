package investec

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type TokenScope int

const (
	TokenScopeAccounts TokenScope = iota
)

type TokenType int

const (
	TokenTypeBearer TokenType = iota
)

type AccessToken struct {
	Token     string
	Type      TokenType
	ExpiresIn int64
	Scope     TokenScope
}

// GetAccessToken obtains an access token.
func (c *client) GetAccessToken(ctx context.Context) (*AccessToken, error) {
	_, body, err := c.get(ctx, "/oauth2/token", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get access token: %w", err)
	}

	data, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body %w", err)
	}

	var accessToken AccessToken
	if err = json.Unmarshal(data, &accessToken); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &accessToken, nil
}
