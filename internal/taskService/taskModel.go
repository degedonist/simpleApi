package taskService

type Task struct {
	ID     string `json:"id"`
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

type TaskRequest struct {
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}
