package main

import (
	"encoding/json"
	"os"
)

type Storage[T any] struct {
	fileName string
}

func NewStorage[T any](filename string) *Storage[T] {
	return &Storage[T]{fileName: filename}
}


func (s *Storage[T]) Save(data T) error {
	fileData , err := json.MarshalIndent(data , "" , "    ")
	if err != nil {
		return err 
	} 
	return os.WriteFile(s.fileName , fileData, 0644)
}

func (s *Storage[T]) Load(data *T) error {
	fileData , err := os.ReadFile(s.fileName)
	if err != nil {
		return err 
	}

	return json.Unmarshal(fileData  , data)
}


