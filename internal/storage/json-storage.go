package storage

import (
	"cli-task-tracker/internal/models"
	"encoding/json"
	"errors"
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

	_, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		if err := createFile(filePath); err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, ErrFileCreated
	}

	if err := storage.LoadTask(); err != nil {
		return nil, ErrCantLoadFile
	}

	return storage, nil
}

func (storage *JsonStorage) LoadTask() error {
	file, err := os.Open(storage.filePath)
	if err != nil {
		return ErrCantLoadFile
	}

	defer file.Close()

	fileInfo, err := file.Stat()

	if err != nil {
		createFile(storage.filePath)
	}

	if fileInfo.Size() == 0 {
		storage.Tasks = make(map[string]*models.Task)
		return nil
	}

	decoder := json.NewDecoder(file)

	var tasks = make(map[string]*models.Task)
	if err := decoder.Decode(&storage.Tasks); err != nil {
		return ErrCantDecodeFile
	}

	storage.Tasks = tasks
	return nil

}

func createFile(filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	return nil
}
