package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Request struct {
	Message string `json:"message"`
}

type Task struct {
	ID        string `json:"task_id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// var tasks = []Task{}

func health(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(Response{Status: "OK", Message: "Health check"})

}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(Response{Status: "OK", Message: "Welcome to TaskForge"})
}

func echo(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	json.NewDecoder(r.Body).Decode(&req)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(Response{Status: "OK", Message: req.Message})
}

func AddTask()

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/health", health)
	http.HandleFunc("/echo", echo)
	// http.HandleFunc("/tasks", getTasks)

	log.Println("Server running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}
