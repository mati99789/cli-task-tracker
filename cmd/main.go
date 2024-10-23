package main

import (
	"cli-task-tracker/internal/storage"
	"log"
)

func main() {

	_, err := storage.NewJsonStorage("tasks.json")

	if err != nil {
		log.Fatalf("Failed to initialize storage: %v", err)
	}

}
