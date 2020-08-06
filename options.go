package investec

import "net/http"

var defaultOptions = clientOptions{
	transport: http.DefaultClient,
}

type clientOptions struct {
	baseURL   string
	transport *http.Client
}

// ClientOption allows the configuration of the Client.
type ClientOption interface {
	Apply(*clientOptions)
}

// ClientOptionFunc is a function to ClientOption adapter.
type ClientOptionFunc func(*clientOptions)

// Apply satisfies the ClientOption interface by applying the ClientOptionFunc
// onto opts.
func (o ClientOptionFunc) Apply(opts *clientOptions) {
	o(opts)
}

// WithBaseURL returns a ClientOptionFunc which configures the client's baseURL
// which prefixes all request paths.
func WithBaseURL(baseURL string) ClientOptionFunc {
	return func(opts *clientOptions) {
		opts.baseURL = baseURL
	}
}

// WithTransport returns a ClientOptionFunc which configures the client
// transport.
func WithTransport(transport *http.Client) ClientOptionFunc {
	return func(opts *clientOptions) {
		opts.transport = transport
	}
}
