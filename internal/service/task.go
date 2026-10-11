package service

import (
	"context"
	"errors"
	"strings"

	"github.com/jevitapearl/TaskForge/internal/models"
	"github.com/jevitapearl/TaskForge/internal/repository"
)

<<<<<<< HEAD
var (
	ErrEmptyTitle     = errors.New("title cannot be empty")
	ErrDuplicateTitle = errors.New("duplicate title")
	ErrTitleTooLong   = errors.New("title too long")
)

=======
>>>>>>> a9364d3794c7d58b367289db3e68c9ff62c5d43c
type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskRepo(repo repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

<<<<<<< HEAD
func (s *TaskService) GetAll(ctx context.Context, userID string, filter models.TaskFilter) ([]models.Task, error) {
	return s.repo.GetAll(ctx, userID, filter)
=======
func (s *TaskService) GetAll(ctx context.Context, userID string) ([]models.Task, error) {
	response, err := s.repo.GetAll(ctx, userID)
	if err != nil {
		return nil, err
	}
	return response, nil
>>>>>>> a9364d3794c7d58b367289db3e68c9ff62c5d43c
}

func (s *TaskService) GetByID(ctx context.Context, userID string, id string) (models.Task, error) {
	return s.repo.GetByID(ctx, userID, id)
}

func (s *TaskService) Create(ctx context.Context, userID string, task models.TaskPayload) error {
<<<<<<< HEAD
	task.Title = strings.TrimSpace(task.Title)

	if task.Title == "" {
=======
	var ErrEmptyTitle = errors.New("title cannot be empty")
	var ErrDuplicateTitle = errors.New("duplicate title")
	var ErrTitleTooLong = errors.New("title too long")
	if strings.TrimSpace(task.Title) == "" {
>>>>>>> a9364d3794c7d58b367289db3e68c9ff62c5d43c
		return ErrEmptyTitle
	}

	if len(task.Title) > 100 {
		return ErrTitleTooLong
	}

<<<<<<< HEAD
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
=======
	if s.ExistsByTitle(ctx, task.Title) {
		return ErrDuplicateTitle
	}
	return s.repo.Create(ctx, userID, models.Task{Title: task.Title, Status: task.Status})
}

func (s *TaskService) Update(ctx context.Context, userID string, id string, task models.UpdatePayload) error {
	return s.repo.Update(ctx, userID, id, models.Task{ID: id, Title: task.Title, Status: task.Status})
>>>>>>> a9364d3794c7d58b367289db3e68c9ff62c5d43c
}

func (s *TaskService) Delete(ctx context.Context, userID string, id string) error {
	return s.repo.Delete(ctx, userID, id)
}

<<<<<<< HEAD
func (s *TaskService) ExistsByTitle(ctx context.Context, userID string, title string) (bool, error) {
	return s.repo.ExistsByTitle(ctx, userID, title)
=======
func (s *TaskService) ExistsByTitle(ctx context.Context, title string) bool {
	return s.repo.ExistsByTitle(ctx, title)
>>>>>>> a9364d3794c7d58b367289db3e68c9ff62c5d43c
}
