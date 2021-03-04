package permissions

type Permission interface {
	Name() string
	Action() Action
}

type Action interface {
}

type Catalog interface {
	CatalogReader
	CatalogWriter
}

type CatalogReader interface {
	Permissions() []Permission
}

type CatalogWriter interface {
	Creator
	Modifier
}

type Creator interface {
	Add(Permission) error
}

type Modifier interface {
	Delete(name string) error
	Update(Permission) error
}
