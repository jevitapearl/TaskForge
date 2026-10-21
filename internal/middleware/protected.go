package middleware

import (
	"net/http"
)

func Protected(next http.Handler) http.Handler {
	return AuthMiddleware(
		LoggingMiddleware(next),
	)
}
