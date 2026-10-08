package homework

import "fmt"

type APIError struct {
	StatusCode int
	Message    string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("homework api: %d %s", e.StatusCode, e.Message)
}
