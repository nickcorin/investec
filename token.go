package ziggy

import (
	"bytes"
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/nickcorin/snorlax"
)

// TokenScope describes the access scope of a token.
type TokenScope string

const (
	TokenScopeAccounts TokenScope = "accounts"
)

// TokenType describes the kind of token.
type TokenType string

const (
	TokenTypeBearer TokenType = "Bearer"
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
	createdAt time.Time
}

// IsExpired returns whether the token has expired.
func (token *AccessToken) IsExpired() bool {
	t := time.Now().Add(time.Second * time.Duration(token.ExpiresIn))
	return time.Now().After(t)
}

// GetAccessToken obtains an access token.
func (c *Client) GetAccessToken(ctx context.Context, scope TokenScope) (
	*AccessToken, error) {

	params := make(url.Values)
	params.Set("grant_type", "client_credentials")
	params.Set("scope", string(scope))

	payload := bytes.NewBuffer([]byte(params.Encode()))
	res, err := c.transport.Post(
		ctx, "/identity/v2/oauth2/token", nil, payload,
		snorlax.WithBasicAuth(c.ClientID, c.ClientSecret))
	if err != nil {
		return nil, fmt.Errorf("failed to get access token: %w", err)
	}

	var accessToken AccessToken
	if err = res.JSON(&accessToken); err != nil {
		return nil, fmt.Errorf("failed to unmarshal access token: %w", err)
	}

	// We set this to be able to calculate whether the token is expired at a
	// later stage.
	accessToken.createdAt = time.Now()

	return &accessToken, nil
}

func (c *Client) SetAccessToken(token *AccessToken) *Client {
	c.token = token
	return c
}
