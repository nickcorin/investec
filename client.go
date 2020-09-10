package ziggy

import (
	"log"
	"net/url"
	"testing"

	"github.com/nickcorin/snorlax"
)

const DefaultURL = "https://openapi.investec.com"

// Client defines a stateful REST client wrapper for the Investec Open API.
type Client struct {
	ClientID     string
	ClientSecret string

	baseURL   string
	proxyURL  *url.URL
	token     *AccessToken
	transport *snorlax.Client
}

// New returns a Client configured with opts.
func NewClient(clientID, clientSecret string) *Client {
	return &Client{
		ClientID:     clientID,
		ClientSecret: clientSecret,

		baseURL:  DefaultURL,
		proxyURL: nil,
		token:    nil,
		transport: snorlax.DefaultClient.
			AddRequestHook(snorlax.WithBasicAuth(clientID, clientSecret)).
			AddRequestHook(snorlax.WithHeader("Accept", "application/json")).
			AddRequestHook(snorlax.WithHeader("Content-Type",
				"x-www-form-urlencoded")).
			SetBaseURL(DefaultURL),
	}
}

// NewClientForTesting returns a Ziggy client with a custom baseURL.
func NewClientForTesting(_ *testing.T, baseURL string) *Client {
	client := NewClient("", "")
	return client.setBaseURL(baseURL)
}

func (c *Client) setBaseURL(u string) *Client {
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
func (c *Client) SetToken(token *AccessToken) *Client {
	c.token = token
	return c
}

// SetProxy sets the proxy URL for the Ziggy client.
func (c *Client) SetProxy(u string) *Client {
	proxyURL, err := url.Parse(u)
	if err != nil {
		log.Println("I DIDN'T SET THE PROXY!")
		// TODO: Add logs to indicate that this failed.
		return c
	}
	log.Println("I DID SET THE PROXY!")
	c.proxyURL = proxyURL
	c.transport.SetProxy(u)

	return c
}
