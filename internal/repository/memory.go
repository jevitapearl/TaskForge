package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jevitapearl/TaskForge/internal/models"
)

type MemoryRepository struct {
	tasks []models.Task
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

func (m *MemoryRepository) GetAll() []models.Task {
	return m.tasks
}

func (m *MemoryRepository) GetByID(id string) (*models.Task, error) {
	entry, i := m.findTask(id)
	if i < 0 {
		return &models.Task{}, errors.New("Could not find entry")
	}
	return entry, nil
}

func (m *MemoryRepository) Create(task models.TaskPayload) error {
	m.tasks = append(m.tasks, models.Task{
		ID:        uuid.NewString(),
		Title:     task.Title,
		Completed: task.Completed,
	})
	return nil
}

func (m *MemoryRepository) Update(id string, new models.UpdatePayload) error {
	_, i := m.findTask(id)

	if i < 0 {
		return errors.New("Could not find entry")
	}
	m.tasks[i].Title = new.Title
	m.tasks[i].Completed = new.Completed
	return nil
}

func (m *MemoryRepository) Delete(id string) error {
	_, i := m.findTask(id)
	if i < 0 {
		return errors.New("Could not find entry")
	}

	m.tasks = append(m.tasks[:i], m.tasks[i+1:]...)
	return nil
}
