package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jevitapearl/TaskForge/internal/models"
)

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(models.Response{Status: http.StatusOK, Message: "Welcome to TaskForge"})
}
