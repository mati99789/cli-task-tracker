package storage

import (
	"cli-task-tracker/internal/models"
	"errors"
)

var (
	ErrEmptyFilePath = errors.New("file path is empty")
	ErrTaskNotFound  = errors.New("task not found")
)

type JsonStorage struct {
	filePath string
	Tasks    map[string]*models.Task
}

func NewJsonStorage(filePath string) (*JsonStorage, error) {

	if filePath == "" {
		return nil, ErrEmptyFilePath
	}

	storage := &JsonStorage{
		filePath: filePath,
		Tasks:    make(map[string]*models.Task),
	}

	return storage, nil
}
