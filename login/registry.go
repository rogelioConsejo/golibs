package login

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
