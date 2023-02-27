package login

import "github.com/google/uuid"

// GetTokenMaker returns a TokenMaker that uses the given persistence
func GetTokenMaker(cPer CredentialsPersistence, tPer TokenPersistence) TokenMaker {
	return tokenMaker{credentialsPersistence: cPer, tokenPersistence: tPer}
}

// TokenMaker generates tokens
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
