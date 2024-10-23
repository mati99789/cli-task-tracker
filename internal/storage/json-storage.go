package storage

import (
	"cli-task-tracker/internal/models"
	"errors"
)

type JsonStorage struct {
	Tasks map[string]*models.Task
}

func NewJsonStorage(filePath string) (*JsonStorage, error) {

	if filePath == "" {
		return nil, errors.New("filePath is empty")
	}

	return &JsonStorage{}, nil
}
