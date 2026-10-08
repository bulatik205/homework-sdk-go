package homework

import (
	"context"
	"net/http"
)

func (c *Client) Ping(ctx context.Context) error {
	resp, err := c.do(ctx, http.MethodGet, "/api/v1/ping", nil, nil)
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
