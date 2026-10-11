package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/jevitapearl/TaskForge/internal/models"
	"github.com/jevitapearl/TaskForge/internal/service"
)

func (h *Handler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	limit := 10
	page := 0

	if value := r.URL.Query().Get("limit"); value != "" {
		n, err := strconv.Atoi(value)
		if err != nil || n <= 0 {
			WriteJSON(w, http.StatusBadRequest, map[string]string{
				"error": "invalid limit",
			})
			return
		}
		limit = n
	}

	if value := r.URL.Query().Get("page"); value != "" {
		n, err := strconv.Atoi(value)
		if err != nil || n < 0 {
			WriteJSON(w, http.StatusBadRequest, map[string]string{
				"error": "invalid page",
			})
			return
		}
		page = n
	}

	var completed *bool

	if value := r.URL.Query().Get("completed"); value != "" {
		v, err := strconv.ParseBool(value)
		if err != nil {
			WriteJSON(w, http.StatusBadRequest, map[string]string{
				"error": "invalid completed value",
			})
			return
		}

		completed = &v
	}

	filter := models.TaskFilter{
		Limit:     limit,
		Page:      page,
		Completed: completed,
	}

	userID := getUserID(r)

	response, err := h.service.GetAll(r.Context(), userID, filter)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	WriteJSON(w, http.StatusOK, response)
}


func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	userID := getUserID(r)

	task, err := h.service.GetByID(r.Context(), userID, id)
	if errors.Is(err, sql.ErrNoRows) {
		WriteError(w, http.StatusNotFound, "task not found")
		return
	}

	if err != nil {
		WriteError(w, http.StatusInternalServerError, "internal server error") 
		return
	}

	WriteJSON(w, http.StatusOK, task)
}


func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.TaskPayload

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		WriteError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	userID := getUserID(r)

	err := h.service.Create(r.Context(), userID, task)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrEmptyTitle),
			errors.Is(err, service.ErrTitleTooLong):
			WriteError(w, http.StatusBadRequest, err.Error())

		case errors.Is(err, service.ErrDuplicateTitle):
			WriteError(w, http.StatusConflict, err.Error())

		default:
			WriteError(w, http.StatusInternalServerError,
				"internal server error")
		}
		return
	}

	WriteJSON(w, http.StatusCreated, models.Response{
		Status:  http.StatusCreated,
		Message: "Created",
	})
}

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	err := h.service.Delete(r.Context(), getUserID(r), id)

	if errors.Is(err, sql.ErrNoRows) {
			WriteError(w, http.StatusNotFound, "task not found")
			return
	}

	if err != nil {
			WriteError(w, http.StatusInternalServerError,
					"internal server error")
			return
	}
	WriteJSON(w, http.StatusOK, models.Response{Status: http.StatusOK, Message: "Deleted"})
}

func (h *Handler) UpdateTask(
	w http.ResponseWriter,
	r *http.Request,
) {

	id := r.PathValue("id")
	userID := getUserID(r)

	var payload models.UpdatePayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		WriteError(w, http.StatusBadRequest, "invalid json")
		return
	}

	err := h.service.Update(r.Context(), userID, id, payload)

	if errors.Is(err, service.ErrEmptyTitle) ||
			errors.Is(err, service.ErrTitleTooLong) {
			WriteError(w, http.StatusBadRequest, err.Error())
			return
	}

	if err != nil {
		WriteError(
			w,
			http.StatusInternalServerError,
			"internal server error",
		)
		return
	}

	WriteJSON(
		w,
		http.StatusOK,
		models.Response{
			Status:  http.StatusOK,
			Message: "Edited",
		},
	)
}
