package service

import (
    "context"

    "github.com/jevitapearl/TaskForge/internal/models"
    "github.com/jevitapearl/TaskForge/internal/repository"
)

type TaskService struct {
  repo repository.TaskRepository
}

func New(repo repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) GetAll(ctx context.Context) []models.Task {
  return s.repo.GetAll(ctx)
}

func (s *TaskService) GetByID(ctx context.Context, id string) (models.Task, error) {
  return s.repo.GetByID(ctx, id)
}

func (s *TaskService) Create(ctx context.Context, task models.TaskPayload) error {
	return s.repo.Create(ctx, models.Task{Title: task.Title, Completed: task.Completed})
}

func (s *TaskService) Update(ctx context.Context, id string, task models.UpdatePayload) error {
  return s.repo.Update(ctx, id, models.Task{ID: id, Title: task.Title, Completed: task.Completed})
}

func (s *TaskService) Delete(ctx context.Context, id string) error {
  return s.repo.Delete(ctx, id)
}