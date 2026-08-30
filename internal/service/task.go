package service

import (
	"context"
	"errors"
	"strings"

	"github.com/jevitapearl/TaskForge/internal/models"
	"github.com/jevitapearl/TaskForge/internal/repository"
)

type TaskService struct {
	repo repository.TaskRepository
}

func New(repo repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) GetAll(ctx context.Context) ([]models.Task, error) {
	response, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *TaskService) GetByID(ctx context.Context, id string) (models.Task, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *TaskService) Create(ctx context.Context, task models.TaskPayload) error {
	var ErrEmptyTitle = errors.New("title cannot be empty")
	var ErrDuplicateTitle = errors.New("duplicate title")
	var ErrTitleTooLong = errors.New("title too long")
	if strings.TrimSpace(task.Title) == "" {
		return ErrEmptyTitle
	}

	if len(task.Title) > 100 {
		return ErrTitleTooLong
	}

	if s.ExistsByTitle(ctx, task.Title) {
		return ErrDuplicateTitle
	}
	return s.repo.Create(ctx, models.Task{Title: task.Title, Completed: task.Completed})
}

func (s *TaskService) Update(ctx context.Context, id string, task models.UpdatePayload) error {
	return s.repo.Update(ctx, id, models.Task{ID: id, Title: task.Title, Completed: task.Completed})
}

func (s *TaskService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *TaskService) ExistsByTitle(ctx context.Context, title string) bool {
	return s.repo.ExistsByTitle(ctx, title)
}
