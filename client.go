package investec

import "github.com/nickcorin/snorlax"

const baseURL = "https://openapi.investec.com"

type client struct {
	clientID         string
	clientSecret     string
	token            string
	tokenAutoRefresh bool
	transport        *snorlax.Client
}

// NewClient returns a Client configured with opts.
func NewClient(opts ...ClientOption) Client {
	c := client{
		tokenAutoRefresh: false,
		transport: snorlax.NewClient(
			snorlax.WithBaseURL(baseURL),
			snorlax.WithRequestOptions(
				snorlax.WithHeader("Accept", "application/json"),
				snorlax.WithHeader("Content-Type", "x-www-form-urlencoded"),
			),
		),
	}

	for _, opt := range opts {
		opt.Apply(&c)
	}

	return &c
}
