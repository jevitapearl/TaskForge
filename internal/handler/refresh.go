package handler

import (
	"net/http"
)

func (h *Handler) Refresh(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed",	http.StatusMethodNotAllowed)
		return
	}

	cookie, err := r.Cookie("refresh_token")

	if err != nil {
		WriteJSON(w, http.StatusUnauthorized, map[string]string{"error": "missing refresh token"})
		return
	}

	accessToken, refreshToken, err := h.authService.Refresh(r.Context(), cookie.Value)

	if err != nil {
		WriteJSON(w, http.StatusUnauthorized, map[string]string{"error": err.Error()})
		return
	}

	http.SetCookie(w,	&http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Path:     "/",
		MaxAge:   900,
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Path:     "/",
		MaxAge:   7 * 24 * 60 * 60,
		HttpOnly: true,
	})

	WriteJSON(w, http.StatusOK,	map[string]string{"message": "token refreshed"})
}