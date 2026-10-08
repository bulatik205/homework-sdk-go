package homework

import (
	"context"
	"fmt"
	"net/http"
)

func (c *Client) AdminTasks(ctx context.Context, date string) ([]Task, error) {
	params := map[string]string{}
	if date != "" {
		params["date"] = date
	}
	resp, err := c.do(ctx, http.MethodGet, "/api/admin/tasks", params, nil)
	if err != nil {
		return nil, err
	}
	r, err := decodeResponse[[]Task](resp)
	if err != nil {
		return nil, err
	}
	return decodeBody(r)
}

func (c *Client) CreateTask(ctx context.Context, req CreateTaskRequest) (*Task, error) {
	resp, err := c.do(ctx, http.MethodPost, "/api/admin/tasks", nil, req)
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

func (c *Client) UpdateTask(ctx context.Context, id int64, req CreateTaskRequest) (*Task, error) {
	path := fmt.Sprintf("/api/admin/tasks/%d", id)
	resp, err := c.do(ctx, http.MethodPatch, path, nil, req)
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

func (c *Client) DeleteTask(ctx context.Context, id int64) error {
	path := fmt.Sprintf("/api/admin/tasks/%d", id)
	resp, err := c.do(ctx, http.MethodDelete, path, nil, nil)
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

func (c *Client) AdminSubjects(ctx context.Context) ([]Subject, error) {
	resp, err := c.do(ctx, http.MethodGet, "/api/admin/subjects", nil, nil)
	if err != nil {
		return nil, err
	}
	r, err := decodeResponse[[]Subject](resp)
	if err != nil {
		return nil, err
	}
	return decodeBody(r)
}

func (c *Client) ScheduleDay(ctx context.Context, date string) (*ScheduleDay, error) {
	params := map[string]string{"date": date}
	resp, err := c.do(ctx, http.MethodGet, "/api/admin/schedule/day", params, nil)
	if err != nil {
		return nil, err
	}
	r, err := decodeResponse[ScheduleDay](resp)
	if err != nil {
		return nil, err
	}
	body, err := decodeBody(r)
	if err != nil {
		return nil, err
	}
	return &body, nil
}

func (c *Client) ScheduleNext(ctx context.Context, subject string) (*ScheduleNext, error) {
	params := map[string]string{"subject": subject}
	resp, err := c.do(ctx, http.MethodGet, "/api/admin/schedule/next", params, nil)
	if err != nil {
		return nil, err
	}
	r, err := decodeResponse[ScheduleNext](resp)
	if err != nil {
		return nil, err
	}
	body, err := decodeBody(r)
	if err != nil {
		return nil, err
	}
	return &body, nil
}
