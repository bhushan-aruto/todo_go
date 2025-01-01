package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

type Storage[T any] struct {
	FileName string
}

func NewStoarge[T any](fileName string) *Storage[T] {
	return &Storage[T]{
		FileName: fileName,
	}
}

func (s *Storage[T]) Save(data T) error {
	fileData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return fmt.Errorf("eror occured while json conversion %v", err)
	}
	return os.WriteFile(s.FileName, fileData, 0644)
}

func (s *Storage[T]) Load(data *T) error {
	fileData, err := os.ReadFile(s.FileName)
	if err != nil {
		return fmt.Errorf("error occured while reading data %v", err)
	}
	return json.Unmarshal(fileData, data)
}
