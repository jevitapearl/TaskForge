package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password"`
	Role         string    `json:"role"`
}

type RefreshToken struct {
	UserID    uuid.UUID
	Token     string
	ExpiresAt time.Time
}
