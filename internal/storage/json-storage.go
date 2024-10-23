package storage

import (
	"cli-task-tracker/internal/models"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

var (
	ErrEmptyFilePath  = errors.New("file path is empty")
	ErrTaskNotFound   = errors.New("task not found")
	ErrFileNotFound   = errors.New("file not found")
	ErrFileCreated    = errors.New("file cant be created")
	ErrCantLoadFile   = errors.New("can't load file")
	ErrCantDecodeFile = errors.New("can't decode file")
)

type JsonStorage struct {
	filePath string
	Tasks    map[int]*models.Task
}

func NewJsonStorage(filePath string) (*JsonStorage, error) {

	if filePath == "" {
		return nil, ErrEmptyFilePath
	}

	storage := &JsonStorage{
		filePath: filePath,
		Tasks:    make(map[int]*models.Task),
	}

	_, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		if err := createFile(filePath); err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, ErrFileCreated
	}

	if err := storage.LoadTasks(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrCantLoadFile, err)
	}

	return storage, nil
}

func (storage *JsonStorage) LoadTasks() error {
	file, err := os.Open(storage.filePath)
	if err != nil {
		return fmt.Errorf("failed to open file : %w", err)
	}

	defer file.Close()

	fileInfo, err := file.Stat()

	if err != nil {
		return fmt.Errorf("Failed to get file info: %w", err)
	}

	if fileInfo.Size() == 0 {
		storage.Tasks = make(map[int]*models.Task)
		return nil
	}

	decoder := json.NewDecoder(file)

	var tasks = make(map[int]*models.Task)
	if err := decoder.Decode(&tasks); err != nil {
		return ErrCantDecodeFile
	}

	storage.Tasks = tasks
	return nil

}

//os.O_RDONLY - Open the file for reading only.
//os.O_WRONLY - Open the file for writing only.
//os.O_RDWR - Open the file for both reading and writing.
//os.O_CREATE - Create the file if it doesn’t exist.
//os.O_APPEND - Append data to the file.
//os.O_TRUNC - Truncate the file to zero size if it already exists.

func createFile(filePath string) error {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	return nil
}
