package login

type TokenPersistence interface {
	Save(name UserName, token Token)
	GetUserName(tk Token) UserName
}

type CredentialsPersistence interface {
	Save(user UserName, password HashedPassword)
	CheckCredentials(name UserName, password HashedPassword) bool
}
