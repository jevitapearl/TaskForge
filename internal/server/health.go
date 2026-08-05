package server

import (
	"net/http"
	"encoding/json"

	"github.com/jevitapearl/TaskForge/internal/models"
)

func (s *Server) Health(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(models.Response{Status: "OK", Message: "Health check"})

}