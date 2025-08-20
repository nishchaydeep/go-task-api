package storage

import (
	"sync"

	"github.com/nishchaydeep15/go-task-api/model"
)

type MemoryStore struct {
	tasks []model.Task
	mutex sync.Mutex
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		tasks: []model.Task{},
	}
}

func (m *MemoryStore) Add(task *model.Task) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.tasks = append(m.tasks, *task)
	return nil
}

func (m *MemoryStore) Delete(task *model.Task) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	var filtered []model.Task
	for _, t := range m.tasks {
		if t.Name != task.Name {
			filtered = append(filtered, t)
		}
	}
	m.tasks = filtered
	return nil
}

func (m *MemoryStore) GetAll() ([]model.Task, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	copyTasks := make([]model.Task, len(m.tasks))
	copy(copyTasks, m.tasks)
	return copyTasks, nil
}

func (m *MemoryStore) Update(task *model.Task) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	for i, t := range m.tasks {
		if t.Name == task.Name {
			m.tasks[i] = *task
			return nil
		}
	}
	return nil
}
