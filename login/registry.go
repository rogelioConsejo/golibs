package login

// GetUserRegistry returns a UserRegistry that uses the given persistence
func GetUserRegistry(p CredentialsPersistence) UserRegistry {
	return userRegistry{credentialsPersistence: p}
}

// UserRegistry adds users
type UserRegistry interface {
	AddUser(u UserName, p Password)
}

type userRegistry struct {
	credentialsPersistence CredentialsPersistence
}

func (u userRegistry) AddUser(usr UserName, p Password) {
	hashedPassword := hashPassword(p)
	u.credentialsPersistence.Save(usr, hashedPassword)
}
