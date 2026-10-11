package middleware

import (
	"context"
	"net/http"

	"github.com/jevitapearl/TaskForge/internal/auth"
	"github.com/jevitapearl/TaskForge/internal/handler"
	"github.com/jevitapearl/TaskForge/internal/models"
)

<<<<<<< HEAD
func AuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
=======
type contextKey string

const (
	UserIDKey contextKey = "userID"
	RoleKey   contextKey = "role"
)

func AuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

>>>>>>> a9364d3794c7d58b367289db3e68c9ff62c5d43c
		cookie, err := r.Cookie("access_token")

		if err != nil {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}

		claims, err := auth.ValidateToken(cookie.Value)

		if err != nil {
			handler.WriteJSON(w, http.StatusUnauthorized, models.Response{Status: http.StatusUnauthorized, Message: "Unauthorized"})
			return
		}

<<<<<<< HEAD
		ctx = context.WithValue(ctx, "userID", claims.UserID)
		ctx = context.WithValue(ctx, "role", claims.Role)
=======
		ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
		ctx = context.WithValue(ctx, RoleKey, claims.Role)
>>>>>>> a9364d3794c7d58b367289db3e68c9ff62c5d43c

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
