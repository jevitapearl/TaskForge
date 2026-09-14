package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jevitapearl/TaskForge/internal/models"
)

func (h *Handler) Echo(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(models.Response{Status: http.StatusAccepted, Message: req.Message})
}
