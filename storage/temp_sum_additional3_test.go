package storage

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/nishchaydeep15/go-task-api/model"
	"github.com/stretchr/testify/assert"
)

func TestNewFileStore(t *testing.T) {
	store := NewFileStore("test.json")
	assert.NotNil(t, store)
	assert.Equal(t, "test.json", store.filename)
}

func TestFileStore_GetAllEmptyFile(t *testing.T) {
	filename := "test_empty.json"
	defer os.Remove(filename)

	store := NewFileStore(filename)
	tasks, err := store.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, 0, len(tasks))
}

func TestFileStore_AddAndGetAll(t *testing.T) {
	filename := "test_add.json"
	defer os.Remove(filename)

	store := NewFileStore(filename)
	task := &model.Task{
		Name:        "Test Task",
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Description: "Test",
	}

	err := store.Add(task)
	assert.NoError(t, err)

	tasks, err := store.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(tasks))
	assert.Equal(t, "Test Task", tasks[0].Name)
}

func TestFileStore_Delete(t *testing.T) {
	filename := "test_delete.json"
	defer os.Remove(filename)

	store := NewFileStore(filename)
	task := &model.Task{Name: "Task1", Completed: false, CreatedAt: time.Now()}
	err := store.Add(task)
	assert.NoError(t, err)

	err = store.Delete(task)
	assert.NoError(t, err)

	tasks, err := store.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, 0, len(tasks))
}

func TestFileStore_DeleteNonExistent(t *testing.T) {
	filename := "test_delete_nonexist.json"
	defer os.Remove(filename)

	store := NewFileStore(filename)
	task1 := &model.Task{Name: "Task1", Completed: false, CreatedAt: time.Now()}
	err := store.Add(task1)
	assert.NoError(t, err)

	task2 := &model.Task{Name: "Task2", Completed: false, CreatedAt: time.Now()}
	err = store.Delete(task2)
	assert.NoError(t, err)

	tasks, err := store.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(tasks))
}

func TestFileStore_Update(t *testing.T) {
	filename := "test_update.json"
	defer os.Remove(filename)

	store := NewFileStore(filename)
	task := &model.Task{Name: "Task", Completed: false, Description: "Old", CreatedAt: time.Now()}
	err := store.Add(task)
	assert.NoError(t, err)

	updatedTask := &model.Task{Name: "Task", Completed: true, Description: "New", CreatedAt: time.Now()}
	err = store.Update(updatedTask)
	assert.NoError(t, err)

	tasks, err := store.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(tasks))
	assert.True(t, tasks[0].Completed)
	assert.Equal(t, "New", tasks[0].Description)
}

func TestFileStore_UpdateNonExistent(t *testing.T) {
	filename := "test_update_nonexist.json"
	defer os.Remove(filename)

	store := NewFileStore(filename)
	task := &model.Task{Name: "NonExistent", Completed: true, CreatedAt: time.Now()}
	err := store.Update(task)
	assert.NoError(t, err)

	tasks, err := store.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, 0, len(tasks))
}

func TestFileStore_LoadFromExistingFile(t *testing.T) {
	filename := "test_existing.json"
	defer os.Remove(filename)

	tasks := []model.Task{
		{Name: "Task1", Completed: false, CreatedAt: time.Now()},
		{Name: "Task2", Completed: true, CreatedAt: time.Now()},
	}
	data, err := json.MarshalIndent(tasks, "", "  ")
	assert.NoError(t, err)
	err = os.WriteFile(filename, data, 0644)
	assert.NoError(t, err)

	store := NewFileStore(filename)
	loadedTasks, err := store.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(loadedTasks))
	assert.Equal(t, "Task1", loadedTasks[0].Name)
	assert.Equal(t, "Task2", loadedTasks[1].Name)
}

func TestFileStore_AddMultiple(t *testing.T) {
	filename := "test_add_multiple.json"
	defer os.Remove(filename)

	store := NewFileStore(filename)
	task1 := &model.Task{Name: "Task1", Completed: false, CreatedAt: time.Now()}
	task2 := &model.Task{Name: "Task2", Completed: true, CreatedAt: time.Now()}

	err := store.Add(task1)
	assert.NoError(t, err)
	err = store.Add(task2)
	assert.NoError(t, err)

	tasks, err := store.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(tasks))
}
