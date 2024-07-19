package file

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

// GetPersistence returns a Persistence that saves and gets something from the given file, it is thread safe and returns
// the same instance for the same file
func GetPersistence(n Name) Persistence {
	if p, ok := persistences[n]; ok {
		return p
	}
	p := &persistence{mutex: NewMutex(n)}
	persistences[n] = p
	return p
}

var persistences = make(map[Name]Persistence)

// New (deprecated) returns a new Persistence that saves and gets something from the given file
func New(n Name) Persistence {
	return &persistence{mutex: NewMutex(n)}
}

// Persistence is a type that can save and get something
// It can be used to save a collection in which case the collection must be handled by the user
type Persistence interface {
	// Save overwrites the file with the given something
	Save(something interface{}) error
	// Get reads the file and stores the content in the given something (which must be a pointer)
	Get(something interface{}) error
	Lock()
	Unlock()
}

type persistence struct {
	fileName Name
	mutex    *Mutex
}

func (p *persistence) Lock() {
	p.mutex.Lock()
}

func (p *persistence) Unlock() {
	p.mutex.Unlock()
}

func (p *persistence) Get(something interface{}) error {
	f, err := p.mutex.GetFile(Read)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	decoder := json.NewDecoder(f)
	err = decoder.Decode(something)
	if err != nil && !errors.Is(err, io.EOF) {
		return err
	}
	return nil
}

func (p *persistence) Save(something interface{}) error {
	f, err := p.mutex.GetFile(Write)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	encoder := json.NewEncoder(f)
	err = encoder.Encode(something)
	if err != nil {
		return err
	}

	return nil
}

type Name string
