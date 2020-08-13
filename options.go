package investec

import "github.com/nickcorin/snorlax"

// ClientOption allows the configuration of the Client.
type ClientOption interface {
	Apply(*client)
}

// ClientOptionFunc is a function to ClientOption adapter.
type ClientOptionFunc func(*client)

// Apply satisfies the ClientOption interface by applying the ClientOptionFunc
// onto opts.
func (o ClientOptionFunc) Apply(opts *client) {
	o(opts)
}

// WithClientID returns a ClientOptionFunc which configures the client's
// clientID used for authentication.
func WithClientID(clientID string) ClientOptionFunc {
	return func(opts *client) {
		opts.clientID = clientID
	}
}

// WithClientSecret returns a ClientOptionFunc which configures the client's
// clientSecret used for authentications.
func WithClientSecret(secret string) ClientOptionFunc {
	return func(opts *client) {
		opts.clientSecret = secret
	}
}

// WithTransport returns a ClientOptionFunc which configures the client
// transport.
func WithTransport(transport *snorlax.Client) ClientOptionFunc {
	return func(opts *client) {
		opts.transport = transport
	}
}
