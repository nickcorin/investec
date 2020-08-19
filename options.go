package investec

import (
	"github.com/nickcorin/snorlax"
)

var defaultOptions = ClientOptions{
	Transport: snorlax.New(&snorlax.ClientOptions{
		BaseURL: "https://openapi.investec.com",
		CallOptions: []snorlax.CallOption{
			snorlax.WithHeader("Accept", "application/json"),
			snorlax.WithHeader("Content-Type", "x-www-form-urlencoded"),
		},
	}),
}

// ClientOptions defines the configurable attributes of the client.
type ClientOptions struct {
	AuthToken    string
	ClientID     string
	ClientSecret string
	Transport    *snorlax.Client
}
