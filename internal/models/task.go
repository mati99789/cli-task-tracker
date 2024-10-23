package models

import (
	"errors"
	"time"
)

type Status string

const (
	InProgress Status = "InProgress"
	ToDo       Status = "ToDo"
	Done       Status = "Done"
)

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func CreateTask(id int, description string, status Status) (*Task, error) {

	if len(description) == 0 {
		return nil, errors.New("description is required")
	}

	if err := status.ValidStatus(); err {
		return nil, errors.New("invalid status")
	}

	return &Task{
		Id:          id,
		Description: description,
		Status:      status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func (s Status) ValidStatus() bool {
	switch s {
	case InProgress, ToDo, Done:
		return true
	default:
		return false
	}
}
