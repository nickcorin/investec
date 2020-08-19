package investec

import (
	"testing"

	"github.com/nickcorin/snorlax"
)

const defaultURL = "https://openapi.investec.com"

type client struct {
	opts *ClientOptions
}

// New returns a Client configured with opts.
func New(opts *ClientOptions) Client {
	if opts == nil {
		opts = &defaultOptions
	}

	if opts.Transport == nil {
		opts.Transport = defaultOptions.Transport
	}

	return &client{opts}
}

// NewForTesting returns a Client to be used for unit testing.
func NewForTesting(_ *testing.T, baseURL string, opts *ClientOptions) Client {
	if opts == nil {
		opts = &defaultOptions
	}

	opts.Transport = snorlax.New(&snorlax.ClientOptions{
		BaseURL: baseURL,
		CallOptions: []snorlax.CallOption{
			snorlax.WithHeader("Accept", "application/json"),
			snorlax.WithHeader("Content-Type",
				"application/x-www-form-urlencoded"),
		},
	})

	return &client{opts}
}
