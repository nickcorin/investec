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
	opts clientOptions
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
	params url.Values, body io.Reader) (int, io.Reader, error) {

	u, err := url.Parse(fmt.Sprintf("%s%s", c.opts.baseURL, path))
	if err != nil {
		return -1, nil, fmt.Errorf("failed to parse url %s: %w", u.String(),
			err)
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), body)
	if err != nil {
		return -1, nil, fmt.Errorf("failed to create http request: %w", err)
	}

	res, err := c.opts.transport.Do(req)
	if err != nil {
		return -1, nil, fmt.Errorf("failed to perform http request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return res.StatusCode, nil, fmt.Errorf("error calling %s: %d", path,
			res.StatusCode)
	}

	bodyBuffer, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return -1, nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return res.StatusCode, bytes.NewBuffer(bodyBuffer), nil
}

func (c *client) delete(ctx context.Context, path string, params url.Values,
	body io.Reader) (int, io.Reader, error) {
	return c.call(ctx, http.MethodDelete, path, params, body)
}

func (c *client) get(ctx context.Context, path string, params url.Values) (
	int, io.Reader, error) {
	return c.call(ctx, http.MethodGet, path, params, nil)
}

func (c *client) post(ctx context.Context, path string, params url.Values,
	body io.Reader) (int, io.Reader, error) {
	return c.call(ctx, http.MethodPost, path, params, body)
}

func (c *client) put(ctx context.Context, path string, params url.Values,
	body io.Reader) (int, io.Reader, error) {
	return c.call(ctx, http.MethodPut, path, params, body)
}
