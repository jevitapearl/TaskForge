package repository

import (
	"context"
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/jevitapearl/TaskForge/internal/models"
)

type MemoryRepository struct {
	tasks []models.Task
	mu    sync.RWMutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		tasks: []models.Task{},
	}
}

func (m *MemoryRepository) findTask(id string) (*models.Task, int) {
	for i, item := range m.tasks {
		if item.ID == id {
			return &m.tasks[i], i
		}
	}
	return nil, -1
}

func (m *MemoryRepository) GetAll(ctx context.Context) []models.Task {
	m.mu.RLock()
	defer m.mu.RUnlock()
	response := make([]models.Task, len(m.tasks))
	copy(response, m.tasks)
	return response
}

func (m *MemoryRepository) GetByID(ctx context.Context, id string) (models.Task, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	entry, i := m.findTask(id)

	if i < 0 {
		return models.Task{}, errors.New("Could not find entry")
	}
	return *entry, nil
}

func (m *MemoryRepository) Create(ctx context.Context, task models.TaskPayload) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.tasks = append(m.tasks, models.Task{
		ID:        uuid.NewString(),
		Title:     task.Title,
		Completed: task.Completed,
	})
	return nil
}

func (m *MemoryRepository) Update(ctx context.Context, id string, new models.UpdatePayload) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, i := m.findTask(id)

	if i < 0 {
		return errors.New("Could not find entry")
	}

	m.tasks[i].Title = new.Title
	m.tasks[i].Completed = new.Completed
	return nil
}

func (m *MemoryRepository) Delete(ctx context.Context, id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	_, i := m.findTask(id)
	if i < 0 {
		return errors.New("Could not find entry")
	}

	m.tasks = append(m.tasks[:i], m.tasks[i+1:]...)
	return nil
}
