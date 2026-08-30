package server

import (
	"encoding/json"
	"net/http"

	"github.com/jevitapearl/TaskForge/internal/models"
)

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (s *Server) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	response, err := s.service.GetAll(r.Context())
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, response)
}

func (s *Server) GetTask(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	response, err := s.service.GetByID(r.Context(), id)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, response)
}

func (s *Server) CreateTask(w http.ResponseWriter, r *http.Request) {
	var newTask models.TaskPayload
	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if err := s.service.Create(r.Context(), newTask); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusCreated, models.Response{Status: "200 OK", Message: "Created"})

}

func (s *Server) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := s.service.Delete(r.Context(), id); err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	writeJSON(w, http.StatusOK, models.Response{Status: "200 OK", Message: "Deleted"})
}

func (s *Server) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var newDetails models.UpdatePayload

	if err := json.NewDecoder(r.Body).Decode(&newDetails); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := s.service.Update(r.Context(), id, newDetails); err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	writeJSON(w, http.StatusOK, models.Response{Status: "200 OK", Message: "Edited"})

}
