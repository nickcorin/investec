package investec

import (
	"bytes"
	"context"
	"fmt"
	"net/url"

	"github.com/nickcorin/snorlax"
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

	payload := bytes.NewBuffer([]byte(params.Encode()))
	res, err := c.opts.Transport.Post(
		ctx, "/identity/v2/oauth2/token", nil, payload,
		snorlax.WithBasicAuth(c.opts.ClientID, c.opts.ClientSecret))
	if err != nil {
		return nil, fmt.Errorf("failed to get access token: %w", err)
	}

	var accessToken AccessToken
	if err = res.JSON(&accessToken); err != nil {
		return nil, fmt.Errorf("failed to unmarshal access token: %w", err)
	}

	return &accessToken, nil
}
