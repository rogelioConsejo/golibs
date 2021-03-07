package persistence

import (
	"errors"
	"github.com/rogelioConsejo/golibs/auth/users"
)

type RepositoryConnection interface {
	IdentifyUser(username string, password string) (*users.User, error)
}

type repositoryConnection struct {
}

func NewRepositoryConnection() RepositoryConnection {
	return &repositoryConnection{}
}

//TODO: Mock
func (repo *repositoryConnection) IdentifyUser(username string, password string) (*users.User, error) {
	return nil, errors.New("not implemented")
}

