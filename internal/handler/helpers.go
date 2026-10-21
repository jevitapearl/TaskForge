package handler

import (
	"net/http"

	"github.com/jevitapearl/TaskForge/internal/contextx"
)

func getUserID(r *http.Request) string {
	return contextx.UserID(r.Context())
}

func getUserRole(r *http.Request) string {
	return contextx.UserRole(r.Context())
}
