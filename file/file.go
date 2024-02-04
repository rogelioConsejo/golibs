package file

import (
	"encoding/json"
	"os"
)

// New returns a new Persistence that saves and gets something from the given file
func New(n Name) Persistence {
	return &persistence{fileName: n}
}

// Persistence is a type that can save and get something
// It can be used to save a collection in which case the collection must be handled by the user
type Persistence interface {
	// Save overwrites the file with the given something
	Save(something interface{}) error
	// Get reads the file and stores the content in the given something (which must be a pointer)
	Get(something interface{}) error
}

type persistence struct {
	fileName Name
}

func (p *persistence) Get(something interface{}) error {
	f, err := os.OpenFile(string(p.fileName), os.O_RDONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
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
	if err != nil {
		return err
	}
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
