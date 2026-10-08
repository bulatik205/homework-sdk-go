package homework

import (
	"context"
	"fmt"
	"net/http"
)

func (c *Client) Tasks(ctx context.Context, filter TasksFilter) ([]Task, error) {
	params := map[string]string{}
	if filter.Date != "" {
		params["date"] = filter.Date
	}
	if filter.Subject != "" {
		params["subject"] = filter.Subject
	}
	if filter.Recency != "" {
		params["recency"] = filter.Recency
	}

	resp, err := c.do(ctx, http.MethodGet, "/api/v1/tasks", params, nil)
	if err != nil {
		return nil, err
	}
	r, err := decodeResponse[[]Task](resp)
	if err != nil {
		return nil, err
	}
	return decodeBody(r)
}

func (c *Client) TaskBySubject(ctx context.Context, subject string) (*Task, error) {
	list, err := c.Tasks(ctx, TasksFilter{Subject: subject, Recency: "last"})
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, nil
	}
	return &list[0], nil
}

func (c *Client) TasksByDate(ctx context.Context, date string) ([]Task, error) {
	return c.Tasks(ctx, TasksFilter{Date: date})
}

func (c *Client) Task(ctx context.Context, id int64) (*Task, error) {
	path := fmt.Sprintf("/api/admin/tasks/%d", id)
	resp, err := c.do(ctx, http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, err
	}
	r, err := decodeResponse[Task](resp)
	if err != nil {
		return nil, err
	}
	body, err := decodeBody(r)
	if err != nil {
		return nil, err
	}
	return &body, nil
}
