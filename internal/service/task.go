package service

import (
	"context"
	"errors"
	"strings"

	"github.com/jevitapearl/TaskForge/internal/models"
	"github.com/jevitapearl/TaskForge/internal/repository"
)

var (
	ErrEmptyTitle     = errors.New("title cannot be empty")
	ErrDuplicateTitle = errors.New("duplicate title")
	ErrTitleTooLong   = errors.New("title too long")
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskRepo(repo repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) GetAll(ctx context.Context, userID string, filter models.TaskFilter) ([]models.Task, error) {
	return s.repo.GetAll(ctx, userID, filter)
}

func (s *TaskService) GetByID(ctx context.Context, userID string, id string) (models.Task, error) {
	return s.repo.GetByID(ctx, userID, id)
}

func (s *TaskService) Create(ctx context.Context, userID string, task models.TaskPayload) error {
	task.Title = strings.TrimSpace(task.Title)

	if task.Title == "" {
		return ErrEmptyTitle
	}

	if len(task.Title) > 100 {
		return ErrTitleTooLong
	}

	exists, err := s.ExistsByTitle(ctx, userID, task.Title)
	if err != nil {
		return err
	}

	if exists {
		return ErrDuplicateTitle
	}

	return s.repo.Create(ctx, userID, models.Task{
		Title:  task.Title,
		Status: task.Status,
	})
}

func (s *TaskService) Update(ctx context.Context, userID string, id string, task models.UpdatePayload) error {
	task.Title = strings.TrimSpace(task.Title)

	if task.Title == "" {
		return ErrEmptyTitle
	}

	if len(task.Title) > 100 {
		return ErrTitleTooLong
	}

	return s.repo.Update(ctx, userID, id, models.Task{
		ID:     id,
		Title:  task.Title,
		Status: task.Status,
	})
}

func (s *TaskService) Delete(ctx context.Context, userID string, id string) error {
	return s.repo.Delete(ctx, userID, id)
}

func (s *TaskService) ExistsByTitle(ctx context.Context, userID string, title string) (bool, error) {
	return s.repo.ExistsByTitle(ctx, userID, title)
}
