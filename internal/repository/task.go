package repository

import (
	"context"

	"github.com/jevitapearl/TaskForge/internal/models"
)

type TaskRepository interface {
	GetAll(ctx context.Context, userID string) ([]models.Task, error)
	GetByID(ctx context.Context, userID string, id string) (models.Task, error)
	Create(ctx context.Context, userID string, task models.Task) error
	Update(ctx context.Context, userID string, id string, new models.Task) error
	Delete(ctx context.Context, userID string, id string) error
	ExistsByTitle(ctx context.Context, title string) bool
}
