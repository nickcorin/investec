package investec

import "net/http"

var defaultOptions = clientOptions{
	transport: http.DefaultClient,
}

type clientOptions struct {
	baseURL      string
	clientID     string
	clientSecret string
	transport    *http.Client
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

// WithClientID returns a ClientOptionFunc which configures the client's
// clientID used for authentication.
func WithClientID(clientID string) ClientOptionFunc {
	return func(opts *clientOptions) {
		opts.clientID = clientID
	}
}

// WithClientSecret returns a ClientOptionFunc which configures the client's
// clientSecret used for authentications.
func WithClientSecret(secret string) ClientOptionFunc {
	return func(opts *clientOptions) {
		opts.clientSecret = secret
	}
}

// WithTransport returns a ClientOptionFunc which configures the client
// transport.
func WithTransport(transport *http.Client) ClientOptionFunc {
	return func(opts *clientOptions) {
		opts.transport = transport
	}
}

// RequestOption allows for configuration of requests made by the client. The
// configuration is single request scoped.
type RequestOption interface {
	Apply(*http.Request)
}

// RequestOptionFunc is a function to RequestOption adapter.
type RequestOptionFunc func(*http.Request)

// Apply satisfies the RequestOption interface by applying the RequestOptionFunc
// to r.
func (o RequestOptionFunc) Apply(r *http.Request) {
	o(r)
}

// WithBasicAuth returns a RequestOption which sets the request's basic
// authentication.
func WithBasicAuth(username, password string) RequestOptionFunc {
	return func(r *http.Request) {
		r.SetBasicAuth(username, password)
	}
}
