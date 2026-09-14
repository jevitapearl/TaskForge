package models

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Request struct {
	Message string `json:"message"`
}
