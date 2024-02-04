package login

// TokenPersistence is needed to save and get tokens (to identify logged-in users)
type TokenPersistence interface {
	Save(name UserName, token Token)
	GetUserName(tk Token) UserName
}

// CredentialsPersistence is needed to save and check credentials (usernames and hashed passwords)
type CredentialsPersistence interface {
	Save(user UserName, password HashedPassword)
	CheckCredentials(name UserName, password HashedPassword) bool
}
