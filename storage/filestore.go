package storage

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/nishchaydeep15/go-task-api/model"
)

type FileStore struct {
	filename string
	mutex    sync.Mutex
}

func NewFileStore(filename string) *FileStore {
	return &FileStore{filename: filename}
}

func (f *FileStore) Add(task *model.Task) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	tasks, err := f.loadFromFile()
	if err != nil {
		return err
	}
	tasks = append(tasks, *task)
	return f.saveToFile(tasks)
}

func (f *FileStore) Delete(task *model.Task) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	tasks, err := f.loadFromFile()
	if err != nil {
		return err
	}

	var filtered []model.Task
	for _, t := range tasks {
		if t.Name != task.Name {
			filtered = append(filtered, t)
		}
	}
	return f.saveToFile(filtered)
}

func (f *FileStore) GetAll() ([]model.Task, error) {
	return f.loadFromFile()
}

func (f *FileStore) loadFromFile() ([]model.Task, error) {
	if _, err := os.Stat(f.filename); os.IsNotExist(err) {
		return []model.Task{}, nil
	}

	data, err := os.ReadFile(f.filename)
	if err != nil {
		return nil, err
	}

	var tasks []model.Task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

func (f *FileStore) saveToFile(tasks []model.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(f.filename, data, 0644)
}
func (f *FileStore) Update(task *model.Task) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	tasks, err := f.loadFromFile()
	if err != nil {
		return err
	}
	for i, t := range tasks {
		if t.Name == task.Name {
			tasks[i] = *task
			return f.saveToFile(tasks)
		}
	}
	return nil
}
