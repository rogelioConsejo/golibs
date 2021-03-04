package users

import "github.com/rogelioConsejo/golibs/auth/roles"

type User interface {
	Name() string
	Roles() []roles.Role
}

type Catalog interface {
	CatalogReader
	CatalogWriter
}

type CatalogReader interface {
	Users() []User
}

type CatalogWriter interface {
	Creator
	Modifier
}

type Creator interface {
	Add(User) error
}

type Modifier interface {
	Delete(name string) error
	Update(User) error
}