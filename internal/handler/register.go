package handler

import (
	"encoding/json"
	"net/http"
)

type RegisterPayload struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req RegisterPayload

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if err := h.authService.Register(r.Context(), req.Email, req.Password); err != nil {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	WriteJSON(
		w,
		http.StatusCreated,
		map[string]string{
			"message": "registered",
		},
	)
}
