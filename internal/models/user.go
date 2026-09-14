package models

import "time"

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password"`
	Role         string `json:"role"`
}

type RefreshToken struct {
	UserID    int
	Token     string
	ExpiresAt time.Time
}
