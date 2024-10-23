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

var (
	ErrEmptyDescirption = errors.New("Empty description")
	ErrInvalidStatus    = errors.New("Invalid status")
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
		return nil, ErrEmptyDescirption
	}

	if err := status.IsValid(); err {
		return nil, ErrInvalidStatus
	}

	return &Task{
		Id:          id,
		Description: description,
		Status:      status,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}, nil
}

func (s Status) IsValid() bool {
	switch s {
	case InProgress, ToDo, Done:
		return true
	default:
		return false
	}
}

func (s Status) String() string {
	return string(s)
}
