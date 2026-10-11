package repository

import (
	"context"

	"github.com/jevitapearl/TaskForge/internal/models"
)

type TaskRepository interface {
<<<<<<< HEAD
	GetAll(ctx context.Context, userID string, filter models.TaskFilter) ([]models.Task, error)
	GetByID(ctx context.Context, userID string, id string) (models.Task, error)
	Create(ctx context.Context, userID string, task models.Task) error
	Update(ctx context.Context, userID string, id string, task models.Task) error
	Delete(ctx context.Context, userID string, id string) error
	ExistsByTitle(ctx context.Context, userID string, title string) (bool, error)
=======
	GetAll(ctx context.Context, userID string) ([]models.Task, error)
	GetByID(ctx context.Context, userID string, id string) (models.Task, error)
	Create(ctx context.Context, userID string, task models.Task) error
	Update(ctx context.Context, userID string, id string, new models.Task) error
	Delete(ctx context.Context, userID string, id string) error
	ExistsByTitle(ctx context.Context, title string) bool
>>>>>>> a9364d3794c7d58b367289db3e68c9ff62c5d43c
}
