package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jevitapearl/TaskForge/internal/models"
)

type AuthRepository interface {
	CreateUser(ctx context.Context, email, hash string) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	StoreRefreshToken(ctx context.Context, userID uuid.UUID, token string, expiresAt time.Time) error
	GetRefreshToken(ctx context.Context, token string) (*models.RefreshToken, error)
	DeleteRefreshToken(ctx context.Context, token string) error
	RotateRefreshToken(ctx context.Context, oldToken string, newToken string, expiresAt time.Time) error
}

func (pr *PostgresRepository) CreateUser(ctx context.Context, email string, hash string) error {

	fmt.Println("Query started")
	_, err := pr.db.ExecContext(
		ctx,
		`
		INSERT INTO users
		(email,password_hash)
		VALUES ($1,$2)
		`,
		email, hash,
	)

	if err != nil {
		fmt.Println("DB error")
	}

	fmt.Println("user created", err)
	return err
}

func (pr *PostgresRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {

	var user models.User

	err := pr.db.QueryRowContext(
		ctx,
		`
		SELECT id,email,password_hash,role
		FROM users
		WHERE email=$1
		`,
		email,
	).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.Role)

	if err != nil {
		fmt.Println("Email not found in query")
		fmt.Println(err)
		return nil, err
	}

	return &user, nil
}

func (pr *PostgresRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*models.User, error) {

	var user models.User

	err := pr.db.QueryRowContext(
		ctx,
		`
		SELECT id,email,password_hash,role
		FROM users
		WHERE id=$1
		`,
		id,
	).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.Role)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (pr *PostgresRepository) StoreRefreshToken(ctx context.Context, userID uuid.UUID, token string, expiresAt time.Time) error {

	_, err := pr.db.ExecContext(
		ctx,
		`
		INSERT INTO refresh_tokens
		(user_id, token, expires_at)
		VALUES ($1,$2,$3)
		`,
		userID, token, expiresAt,
	)

	return err
}

func (pr *PostgresRepository) GetRefreshToken(ctx context.Context, token string) (*models.RefreshToken, error) {

	var rt models.RefreshToken

	err := pr.db.QueryRowContext(
		ctx,
		`
		SELECT user_id,token,expires_at
		FROM refresh_tokens
		WHERE token=$1
		`,
		token,
	).Scan(&rt.UserID, &rt.Token, &rt.ExpiresAt)

	if err != nil {
		return nil, err
	}

	return &rt, nil
}

func (pr *PostgresRepository) DeleteRefreshToken(ctx context.Context, token string) error {

	_, err := pr.db.ExecContext(
		ctx,
		`
		DELETE FROM refresh_tokens
		WHERE token=$1
		`,
		token,
	)

	return err
}

func (pr *PostgresRepository) RotateRefreshToken(ctx context.Context, oldToken string, newToken string, expiresAt time.Time) error {

	tx, err := pr.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	defer tx.Rollback()

	var userID int

	err = tx.QueryRowContext(
		ctx,
		`
		SELECT user_id
		FROM refresh_tokens
		WHERE token=$1
		`,
		oldToken,
	).Scan(&userID)

	if err != nil {
		return err
	}

	_, err = tx.ExecContext(
		ctx,
		`
		DELETE FROM refresh_tokens
		WHERE token=$1
		`,
		oldToken,
	)

	if err != nil {
		return err
	}

	_, err = tx.ExecContext(
		ctx,
		`
		INSERT INTO refresh_tokens
		(user_id, token, expires_at)
		VALUES ($1,$2,$3)
		`,
		userID, newToken, expiresAt,
	)

	if err != nil {
		return err
	}

	return tx.Commit()
}
