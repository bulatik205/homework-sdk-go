package homework

import (
	"context"
	"net/http"
)

func (c *Client) Register(ctx context.Context, req RegisterRequest) (*User, error) {
	resp, err := c.do(ctx, http.MethodPost, "/api/auth/register", nil, req)
	if err != nil {
		return nil, err
	}
	r, err := decodeResponse[User](resp)
	if err != nil {
		return nil, err
	}
	body, err := decodeBody(r)
	if err != nil {
		return nil, err
	}
	return &body, nil
}

func (c *Client) Login(ctx context.Context, username, password string) (*User, error) {
	resp, err := c.do(ctx, http.MethodPost, "/api/auth/login", nil, LoginRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	r, err := decodeResponse[User](resp)
	if err != nil {
		return nil, err
	}
	body, err := decodeBody(r)
	if err != nil {
		return nil, err
	}
	return &body, nil
}

func (c *Client) Logout(ctx context.Context) error {
	resp, err := c.do(ctx, http.MethodPost, "/api/auth/logout", nil, nil)
	if err != nil {
		return err
	}
	r, err := decodeResponse[any](resp)
	if err != nil {
		return err
	}
	if !r.Success {
		return &APIError{StatusCode: r.Code, Message: r.Error}
	}
	return nil
}

func (c *Client) Me(ctx context.Context) (*User, error) {
	resp, err := c.do(ctx, http.MethodGet, "/api/auth/me", nil, nil)
	if err != nil {
		return nil, err
	}
	r, err := decodeResponse[User](resp)
	if err != nil {
		return nil, err
	}
	body, err := decodeBody(r)
	if err != nil {
		return nil, err
	}
	return &body, nil
}
