package login

import "github.com/google/uuid"

type TokenMaker interface {
	GenerateToken(UserName, Password) Token
}

type Token string

type tokenMaker struct {
	credentialsPersistence CredentialsPersistence
	tokenPersistence       TokenPersistence
}

func (d tokenMaker) GenerateToken(name UserName, password Password) Token {
	if d.credentialsPersistence.CheckCredentials(name, hashPassword(password)) {
		tk := makeToken()
		d.tokenPersistence.Save(name, tk)
		return tk
	}

	return ""
}

func makeToken() Token {
	return Token(uuid.NewString())
}
