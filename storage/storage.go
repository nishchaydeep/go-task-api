package storage

import (
	"encoding/json"
	"os"

	"github.com/nishchaydeep15/go-task-api/model"
)

const filePath = "tasks.json"

func SaveTasks(tasks []model.Task) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0644)
}

func LoadTasks() ([]model.Task, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return []model.Task{}, nil // Return empty slice if file does not exist
	}
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var tasks []model.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}
