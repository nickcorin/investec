package client

import (
	"net/url"
	"testing"

	"github.com/nickcorin/ziggy"

	"github.com/sirupsen/logrus"
	"github.com/nickcorin/snorlax"
)

// httpClient defines a stateful REST client wrapper for the Investec Open API.
type httpClient struct {
	ClientID     string
	ClientSecret string

	baseURL   string
	proxyURL  *url.URL
	token     *ziggy.AccessToken
	transport snorlax.Client
}

// New returns a Client configured with opts.
func NewHTTP(clientID, clientSecret string) ziggy.Client {
	return &httpClient{
		ClientID:     clientID,
		ClientSecret: clientSecret,

		baseURL:  ziggy.DefaultURL,
		proxyURL: nil,
		token:    nil,
		transport: snorlax.DefaultClient.
			AddRequestHook(snorlax.WithHeader("Accept", "application/json")).
			AddRequestHook(snorlax.WithHeader("Content-Type",
				"application/x-www-form-urlencoded")).
			SetLogLevel(logrus.TraceLevel).
			SetBaseURL(ziggy.DefaultURL),
	}
}

// NewHTTPForTesting returns a Ziggy client with a custom baseURL.
func NewHTTPForTesting(_ *testing.T, baseURL string) *httpClient {
	client := NewHTTP("", "")
	return client.(*httpClient).setBaseURL(baseURL)
}

func (c *httpClient) setBaseURL(u string) *httpClient {
	baseURL, err := url.Parse(u)
	if err != nil {
		// TODO: Add logs to indicate that this failed.
		return c
	}

	c.transport.SetBaseURL(baseURL.String())
	c.baseURL = baseURL.String()
	return c
}

// SetToken sets a temporary authorization token on the Ziggy client.
func (c *httpClient) SetToken(token *ziggy.AccessToken) *httpClient {
	c.token = token
	return c
}

// SetProxy sets the proxy URL for the Ziggy client.
func (c *httpClient) SetProxy(u string) *httpClient {
	proxyURL, err := url.Parse(u)
	if err != nil {
		// TODO: Add logs to indicate that this failed.
		return c
	}
	c.proxyURL = proxyURL
	c.transport.SetProxy(u)

	return c
}
