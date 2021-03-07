package authentication

import (
	"github.com/rogelioConsejo/golibs/auth/authentication/persistence"
	"github.com/rogelioConsejo/golibs/auth/users"
)

type UserAuthentication interface {
	Identify(username string, password string) (*users.User, error)
}

type userAuthentication struct {
}

func NewUserAuthentication() *userAuthentication {
	return &userAuthentication{}
}

func (auth *userAuthentication) Identify(username string, password string) (*users.User, error) {
	repo := persistence.NewRepositoryConnection()
	return repo.IdentifyUser(username, password)
}

