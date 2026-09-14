package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jevitapearl/TaskForge/internal/auth"
	"github.com/jevitapearl/TaskForge/internal/repository"
)

type AuthService struct {
	repo repository.AuthRepository
}

func NewAuthRepo(repo repository.AuthRepository) *AuthService {
	return &AuthService{repo}
}

func (s *AuthService) Register(ctx context.Context, email string, password string) error {

	hash, err := auth.HashPassword(password)
	if err != nil {
		fmt.Println("service error")
		return err
	}

	fmt.Println("repo called")
	return s.repo.CreateUser(ctx, email, hash)
}

func (s *AuthService) Login(ctx context.Context, email string, password string) (string, string, error) {

	user, err := s.repo.GetUserByEmail(ctx, email)

	if err != nil {
		return "", "", errors.New("invalid credentials")
	}

	if !auth.VerifyPassword(
		password,
		user.PasswordHash,
	) {
		return "", "", errors.New("invalid credentials")
	}

	accessToken, err := auth.GenerateAccessToken(
		user.ID,
		user.Role,
	)

	if err != nil {
		return "", "", err
	}

	refreshToken, err :=
		auth.GenerateRefreshToken(
			user.ID,
		)

	if err != nil {
		return "", "", err
	}

	err = s.repo.StoreRefreshToken(
		ctx,
		user.ID,
		refreshToken,
		time.Now().Add(
			7*24*time.Hour,
		),
	)

	if err != nil {
		return "", "", err
	}

	return accessToken,
		refreshToken,
		nil
}

func (s *AuthService) Refresh(ctx context.Context, refreshToken string) (string, string, error) {

	tokenRecord, err := s.repo.GetRefreshToken(ctx, refreshToken)

	if err != nil {
		return "", "", errors.New("invalid refresh token")
	}

	if time.Now().After(tokenRecord.ExpiresAt) {
		_ = s.repo.DeleteRefreshToken(ctx, refreshToken)
		return "", "", errors.New("refresh token expired")
	}

	user, err := s.repo.GetUserByID(ctx, tokenRecord.UserID)

	if err != nil {
		return "", "", err
	}

	newRefreshToken, err := auth.GenerateRefreshToken(user.ID)

	if err != nil {
		return "", "", err
	}

	err = s.repo.RotateRefreshToken(
		ctx,
		refreshToken,
		newRefreshToken,
		time.Now().Add(7*24*time.Hour),
	)

	if err != nil {
		return "", "", err
	}

	accessToken, err := auth.GenerateAccessToken(user.ID, user.Role)

	if err != nil {
		return "", "", err
	}

	return accessToken,
		newRefreshToken,
		nil
}

func (s *AuthService) Logout(ctx context.Context, refreshToken string) error {
	return s.repo.DeleteRefreshToken(ctx, refreshToken)
}
