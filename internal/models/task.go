package models

type Task struct {
	ID     string `json:"id"`
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

type RequestBody struct {
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}
