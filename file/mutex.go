package file

import (
	"errors"
	"os"
	"sync"
)

func NewMutex(fileName Name) *Mutex {
	return &Mutex{fileName: fileName}
}

type Mutex struct {
	fileName Name
	mutex    sync.Mutex
}

func (fm *Mutex) Lock() {
	fm.mutex.Lock()
}

func (fm *Mutex) Unlock() {
	fm.mutex.Unlock()
}

func (fm *Mutex) GetFile(action Action) (*os.File, error) {
	switch action {
	case Read:
		return os.OpenFile(string(fm.fileName), os.O_RDONLY|os.O_APPEND|os.O_CREATE, 0644)
	case Write:
		return os.OpenFile(string(fm.fileName), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	default:
		return nil, errors.New("invalid action")
	}
}

type Action string

const (
	Read  Action = "read"
	Write Action = "write"
)
