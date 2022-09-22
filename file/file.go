package file

import (
	"encoding/json"
	"os"
)

func New(n Name) Persistence {
	return &persistence{fileName: n}
}

type Persistence interface {
	Save(something interface{}) error
	Get(something interface{}) error
}

type persistence struct {
	fileName Name
}

func (p *persistence) Get(something interface{}) error {
	f, err := os.OpenFile(string(p.fileName), os.O_RDONLY|os.O_APPEND|os.O_CREATE, 0644)
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&something)
	if err != nil {
		return err
	}
	return nil
}

func (p *persistence) Save(something interface{}) error {
	f, err := os.OpenFile(string(p.fileName), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	if err != nil {
		return err
	}
	encoder := json.NewEncoder(f)
	err = encoder.Encode(something)
	if err != nil {
		return err
	}

	return nil
}

type Name string
