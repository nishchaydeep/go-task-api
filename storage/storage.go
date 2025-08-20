package storage

import (
	"fmt"

	"github.com/nishchaydeep15/go-task-api/config"
	"github.com/nishchaydeep15/go-task-api/model"
)

type Storage interface {
	Add(task *model.Task) error
	Delete(task *model.Task) error
	GetAll() ([]model.Task, error)
	Update(task *model.Task) error
}

var Store Storage

func InitializeStorage() error {
	switch config.Conf.Storage {
	case "file":
		Store = NewFileStore(config.Conf.TaskFilePath)
	case "memory":
		Store = NewMemoryStore()
	default:
		return fmt.Errorf("invalid storage type: %s", config.Conf.Storage)
	}
	return nil
}
