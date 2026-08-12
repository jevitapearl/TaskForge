package repository

import "github.com/jevitapearl/TaskForge/internal/models"

type TaskRepository interface {
	GetAll() []models.Task
	GetByID(id string) (models.Task, error)
	Create(task models.TaskPayload) error
	Update(id string, new models.UpdatePayload) error
	Delete(id string) error
}
