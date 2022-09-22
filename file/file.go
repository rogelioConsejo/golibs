package file

import (
	"encoding/json"
	"fmt"
	"os"
)

func New(n Name) (Persistence, error) {
	f, err := os.OpenFile(string(n), os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, fmt.Errorf("could not create file persistence access: %w", err)
	}
	return persistence{file: f}, nil
}

type Persistence interface {
	Save(something interface{}) error
	Get(something interface{}) error
	Close() error
}

type persistence struct {
	file *os.File
}

func (p persistence) Close() error {
	err := p.file.Close()
	if err != nil {
		return err
	}
	return nil
}

func (p persistence) Get(something interface{}) error {
	decoder := json.NewDecoder(p.file)
	err := decoder.Decode(&something)
	if err != nil {
		return err
	}
	return nil
}

func (p persistence) Save(something interface{}) error {
	encoder := json.NewEncoder(p.file)
	err := encoder.Encode(something)
	if err != nil {
		return err
	}
	return nil
}

type Name string
