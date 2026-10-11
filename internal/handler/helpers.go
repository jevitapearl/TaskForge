package handler

import (
	"context"
	"net/http"
)

type contextKey string

const userIDKey contextKey = "userID"

func getUserID(r *http.Request) string {
	userID, _ := r.Context().Value(userIDKey).(string)
	return userID
}

func setUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}
