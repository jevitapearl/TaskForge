package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jevitapearl/TaskForge/internal/repository"
	"github.com/jevitapearl/TaskForge/internal/service"
)

type Handler struct {
	service     *service.TaskService
	authService *service.AuthService
}

func New(repo *repository.PostgresRepository) *Handler {
	return &Handler{
		service:     service.NewTaskRepo(repo),
		authService: service.NewAuthRepo(repo),
	}
}

func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
