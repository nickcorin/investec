package client

import (
	"bytes"
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/nickcorin/ziggy"

	"github.com/nickcorin/snorlax"
)

// GetAccessToken satisfies the ziggy.Client interface.
func (c *httpClient) GetAccessToken(ctx context.Context,
	scope ziggy.TokenScope) (*ziggy.AccessToken, error) {

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

	if !res.IsSuccess() {
		return nil, fmt.Errorf("failed to fetch access token: %d status " +
		"code received", res.StatusCode)
	}

	var accessToken ziggy.AccessToken
	if err = res.JSON(&accessToken); err != nil {
		return nil, fmt.Errorf("failed to unmarshal access token: %w", err)
	}

	// We set this to be able to calculate whether the token is expired at a
	// later stage.
	accessToken.CreatedAt = time.Now()

	c.SetToken(&accessToken)

	return &accessToken, nil
}

func (c *httpClient) SetAccessToken(token *ziggy.AccessToken) *httpClient {
	c.token = token
	return c
}
