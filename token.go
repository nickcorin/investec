package investec

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

type TokenScope string

const (
	TokenScopeAccounts TokenScope = "accounts"
)

type TokenType string

const (
	TokenTypeBearer TokenType = "Bearer"
)

type AccessToken struct {
	Token     string     `json:"access_token"`
	Type      TokenType  `json:"token_type"`
	ExpiresIn int64      `json:"expires_in"`
	Scope     TokenScope `json:"scope"`
}

// GetAccessToken obtains an access token.
func (c *client) GetAccessToken(ctx context.Context, scope TokenScope) (
	*AccessToken, error) {

	params := make(url.Values)
	params.Set("grant_type", "client_credentials")
	params.Set("scope", string(scope))

	_, body, err := c.post(ctx, "/identity/v2/oauth2/token", nil,
		bytes.NewBuffer([]byte(params.Encode())), WithBasicAuth(c.opts.clientID,
			c.opts.clientSecret))
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
