package models

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Request struct {
	Message string `json:"message"`
}