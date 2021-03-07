package users

type User interface {
	Name() string
}

type user struct {
	name string
}

func NewUser(name string, password string) *user {
	newUser := user{name: name}
	return &newUser
}

func (u *user) Name() string {
	return u.name
}

