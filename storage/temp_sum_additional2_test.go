package storage

import (
	"testing"
	"time"

	"github.com/nishchaydeep15/go-task-api/model"
	"github.com/stretchr/testify/assert"
)

func TestNewMemoryStore(t *testing.T) {
	store := NewMemoryStore()
	assert.NotNil(t, store)
	assert.NotNil(t, store.tasks)
	tasks, err := store.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, 0, len(tasks))
}

func TestMemoryStore_Add(t *testing.T) {
	store := NewMemoryStore()
	task := &model.Task{
		Name:        "Test Task",
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		AccessesAt:  time.Now(),
		Description: "Test description",
		Category:    "test",
		Important:   false,
	}

	err := store.Add(task)
	assert.NoError(t, err)

	tasks, err := store.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(tasks))
	assert.Equal(t, "Test Task", tasks[0].Name)
}

func TestMemoryStore_AddMultiple(t *testing.T) {
	store := NewMemoryStore()

	task1 := &model.Task{Name: "Task 1", Completed: false, CreatedAt: time.Now()}
	task2 := &model.Task{Name: "Task 2", Completed: true, CreatedAt: time.Now()}

	err := store.Add(task1)
	assert.NoError(t, err)
	err = store.Add(task2)
	assert.NoError(t, err)

	tasks, err := store.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(tasks))
}

func TestMemoryStore_Delete(t *testing.T) {
	store := NewMemoryStore()

	task := &model.Task{Name: "To Delete", Completed: false, CreatedAt: time.Now()}
	err := store.Add(task)
	assert.NoError(t, err)

	err = store.Delete(task)
	assert.NoError(t, err)

	tasks, err := store.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, 0, len(tasks))
}

func TestMemoryStore_DeleteNonExistent(t *testing.T) {
	store := NewMemoryStore()

	task1 := &model.Task{Name: "Task 1", Completed: false, CreatedAt: time.Now()}
	err := store.Add(task1)
	assert.NoError(t, err)

	task2 := &model.Task{Name: "Task 2", Completed: false, CreatedAt: time.Now()}
	err = store.Delete(task2)
	assert.NoError(t, err)

	tasks, err := store.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(tasks))
}

func TestMemoryStore_GetAll(t *testing.T) {
	store := NewMemoryStore()

	tasks, err := store.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, 0, len(tasks))

	task := &model.Task{Name: "Task", Completed: false, CreatedAt: time.Now()}
	err = store.Add(task)
	assert.NoError(t, err)

	tasks, err = store.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(tasks))
}

func TestMemoryStore_Update(t *testing.T) {
	store := NewMemoryStore()

	task := &model.Task{Name: "Original", Completed: false, Description: "Old", CreatedAt: time.Now()}
	err := store.Add(task)
	assert.NoError(t, err)

	updatedTask := &model.Task{Name: "Original", Completed: true, Description: "New", CreatedAt: time.Now()}
	err = store.Update(updatedTask)
	assert.NoError(t, err)

	tasks, err := store.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(tasks))
	assert.True(t, tasks[0].Completed)
	assert.Equal(t, "New", tasks[0].Description)
}

func TestMemoryStore_UpdateNonExistent(t *testing.T) {
	store := NewMemoryStore()

	task := &model.Task{Name: "NonExistent", Completed: true, CreatedAt: time.Now()}
	err := store.Update(task)
	assert.NoError(t, err)

	tasks, err := store.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, 0, len(tasks))
}

func TestMemoryStore_GetAllReturnsDeepCopy(t *testing.T) {
	store := NewMemoryStore()

	task := &model.Task{Name: "Task", Completed: false, CreatedAt: time.Now()}
	err := store.Add(task)
	assert.NoError(t, err)

	tasks1, err := store.GetAll()
	assert.NoError(t, err)

	tasks2, err := store.GetAll()
	assert.NoError(t, err)

	assert.Equal(t, len(tasks1), len(tasks2))
}
