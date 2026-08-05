package models

type Task struct {
	ID        string `json:"task_id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type TaskPayload struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type UpdatePayload struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
