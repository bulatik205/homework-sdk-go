package homework

import (
	"context"
	"net/http"
)

func (c *Client) Subjects(ctx context.Context) ([]Subject, error) {
	resp, err := c.do(ctx, http.MethodGet, "/api/v1/subjects", nil, nil)
	if err != nil {
		return nil, err
	}
	r, err := decodeResponse[[]Subject](resp)
	if err != nil {
		return nil, err
	}
	return decodeBody(r)
}
