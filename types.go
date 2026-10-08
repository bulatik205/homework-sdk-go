package homework

type Response[T any] struct {
	Success bool   `json:"success"`
	Body    T      `json:"body,omitempty"`
	Code    int    `json:"code,omitempty"`
	Error   string `json:"error,omitempty"`
}

type Task struct {
	ID        int64   `json:"id"`
	Subject   string  `json:"subject"`
	Task      string  `json:"task"`
	DateFrom  *string `json:"date_from,omitempty"`
	DateTo    *string `json:"date_to,omitempty"`
	Status    *string `json:"status,omitempty"`
	InsteadOf *string `json:"instead_of,omitempty"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
}

type Subject struct {
	Code    string `json:"code"`
	Display string `json:"display"`
}

type ScheduleDay struct {
	Date    string         `json:"date"`
	Day     int            `json:"day"`
	DayName string         `json:"day_name"`
	Lessons []ScheduleItem `json:"lessons"`
}

type ScheduleItem struct {
	Code    string `json:"code"`
	Display string `json:"display"`
}

type ScheduleNext struct {
	Date    string `json:"date"`
	Day     int    `json:"day"`
	DayName string `json:"day_name"`
}

type CreateTaskRequest struct {
	Subject   string  `json:"subject"`
	Task      string  `json:"task"`
	DateFrom  *string `json:"date_from,omitempty"`
	DateTo    *string `json:"date_to,omitempty"`
	Status    *string `json:"status,omitempty"`
	InsteadOf *string `json:"instead_of,omitempty"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Key      string `json:"key"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TasksFilter struct {
	Date    string
	Subject string
	Recency string
}
