package investec

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type client struct {
	accessToken string
	opts        clientOptions
}

// NewClient returns a Client configured with opts.
func NewClient(opts ...ClientOption) Client {
	c := client{
		opts: defaultOptions,
	}

	for _, opt := range opts {
		opt.Apply(&c.opts)
	}

	return &c
}

func (c *client) call(ctx context.Context, method, path string,
	params url.Values, body io.Reader, opts ...RequestOption) (int,
	io.Reader, error) {

	uri := fmt.Sprintf("%s%s", c.opts.baseURL, path)
	u, err := url.Parse(uri)
	if err != nil {
		return -1, nil, fmt.Errorf("failed to parse url %s: %w", uri, err)
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), body)
	if err != nil {
		return -1, nil, fmt.Errorf("failed to create http request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	res, err := c.opts.transport.Do(req)
	if err != nil {
		return -1, nil, fmt.Errorf("failed to perform http request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return res.StatusCode, nil, fmt.Errorf("error calling %s: %s", path,
			http.StatusText(res.StatusCode))
	}

	bodyBuffer, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return -1, nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return res.StatusCode, bytes.NewBuffer(bodyBuffer), nil
}

func (c *client) delete(ctx context.Context, path string, params url.Values,
	body io.Reader, opts ...RequestOption) (int, io.Reader, error) {
	return c.call(ctx, http.MethodDelete, path, params, body, opts...)
}

func (c *client) get(ctx context.Context, path string, params url.Values,
	opts ...RequestOption) (int, io.Reader, error) {
	return c.call(ctx, http.MethodGet, path, params, nil, opts...)
}

func (c *client) post(ctx context.Context, path string, params url.Values,
	body io.Reader, opts ...RequestOption) (int, io.Reader, error) {
	return c.call(ctx, http.MethodPost, path, params, body, opts...)
}

func (c *client) put(ctx context.Context, path string, params url.Values,
	body io.Reader, opts ...RequestOption) (int, io.Reader, error) {
	return c.call(ctx, http.MethodPut, path, params, body, opts...)
}
