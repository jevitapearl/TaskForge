package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jevitapearl/TaskForge/internal/models"
)

func (h *Handler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	response, err := h.service.GetAll(r.Context(), r.Context().Value("userID").(string))
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	WriteJSON(w, http.StatusOK, response)
}

func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	response, err := h.service.GetByID(r.Context(), r.Context().Value("userID").(string), id)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	WriteJSON(w, http.StatusOK, response)
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var newTask models.TaskPayload
	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if err := h.service.Create(r.Context(), r.Context().Value("userID").(string), newTask); err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	WriteJSON(w, http.StatusCreated, models.Response{Status: http.StatusOK, Message: "Created"})

}

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := h.service.Delete(r.Context(), r.Context().Value("userID").(string), id); err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	WriteJSON(w, http.StatusOK, models.Response{Status: http.StatusOK, Message: "Deleted"})
}

func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var newDetails models.UpdatePayload

	if err := json.NewDecoder(r.Body).Decode(&newDetails); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := h.service.Update(r.Context(), r.Context().Value("userID").(string), id, newDetails); err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	WriteJSON(w, http.StatusOK, models.Response{Status: http.StatusOK, Message: "Edited"})

}
