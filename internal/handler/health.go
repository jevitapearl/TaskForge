package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jevitapearl/TaskForge/internal/models"
)

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(models.Response{Status: http.StatusAccepted, Message: "Health check"})

}
