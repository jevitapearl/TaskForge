package repository

import (
	"context"

	"github.com/jevitapearl/TaskForge/internal/models"
)

type TaskRepository interface {
	GetAll(ctx context.Context) []models.Task
	GetByID(ctx context.Context, id string) (models.Task, error)
	Create(ctx context.Context, task models.Task) error
	Update(ctx context.Context, id string, new models.Task) error
	Delete(ctx context.Context, id string) error
}
