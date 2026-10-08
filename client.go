package homework

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

type Client struct {
	BaseURL    string
	ProxyURL   string
	Secret     string
	HTTPClient *http.Client
}

type Option func(*Client)

func WithBaseURL(u string) Option {
	return func(c *Client) { c.BaseURL = u }
}

func WithProxy(proxyURL, secret string) Option {
	return func(c *Client) {
		c.ProxyURL = proxyURL
		c.Secret = secret
	}
}

func WithHTTPClient(h *http.Client) Option {
	return func(c *Client) { c.HTTPClient = h }
}

func New(opts ...Option) *Client {
	jar, _ := cookiejar.New(nil)
	c := &Client{
		BaseURL: "https://hw.bulatik.website",
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
			Jar:     jar,
		},
	}
	for _, o := range opts {
		o(c)
	}
	return c
}

func (c *Client) buildURL(path string, params map[string]string) string {
	var raw string

	if c.ProxyURL != "" {
		q := url.Values{}
		if c.Secret != "" {
			q.Set("secret", c.Secret)
		}
		q.Set("path", path)
		raw = c.ProxyURL + "?" + q.Encode()
	} else {
		raw = c.BaseURL + path
	}

	if len(params) > 0 {
		u, err := url.Parse(raw)
		if err == nil {
			q := u.Query()
			for k, v := range params {
				if v != "" {
					q.Set(k, v)
				}
			}
			u.RawQuery = q.Encode()
			raw = u.String()
		}
	}

	return raw
}

func (c *Client) do(ctx context.Context, method, path string, params map[string]string, body any) (*http.Response, error) {
	rawURL := c.buildURL(path, params)

	var buf io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, rawURL, buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return c.HTTPClient.Do(req)
}

func decodeResponse[T any](resp *http.Response) (*Response[T], error) {
	defer resp.Body.Close()

	var out Response[T]
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	if resp.StatusCode >= 400 && out.Error == "" {
		out.Error = resp.Status
	}
	return &out, nil
}

func decodeBody[T any](r *Response[T]) (T, error) {
	var zero T
	if !r.Success {
		return zero, &APIError{StatusCode: r.Code, Message: r.Error}
	}
	return r.Body, nil
}
