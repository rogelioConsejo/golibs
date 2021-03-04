package roles

import (
	"github.com/rogelioConsejo/golibs/auth/roles/permissions"
)

type Role interface {
	Name() string
	Permissions() []permissions.Permission
}

type Configurator interface {
	ChangeName(name string)
	AddPermission(permission permissions.Permission)
	RemovePermission(id string)
}

type Catalog interface {
	CatalogReader
	CatalogWriter
}

type CatalogReader interface {
	Roles() []Role
}

type CatalogWriter interface {
	Creator
	Modifier
}

type Creator interface {
	Add(Role) error
}

type Modifier interface {
	Delete(name string) error
	Update(Role) error
}