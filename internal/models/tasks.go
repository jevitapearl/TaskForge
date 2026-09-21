package models

type Task struct {
	ID     string `json:"task_id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

type TaskPayload struct {
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

type UpdatePayload struct {
	Title  string `json:"title"`
	Status bool   `json:"status"`
}
