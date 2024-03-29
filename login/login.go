package login

import (
	"crypto/sha256"
	"fmt"
)

// GetDoorman returns a Doorman that uses the given persistence
func GetDoorman(p TokenPersistence) Doorman {
	return doorman{tokenPersistence: p}
}

// Doorman validates tokens
type Doorman interface {
	ValidateToken(tk Token) (UserName, bool)
}

func hashPassword(s Password) HashedPassword {
	hasher := sha256.New()
	hasher.Write([]byte(s))
	return HashedPassword(fmt.Sprintf("%x", hasher.Sum(nil)))
}

type UserName string
type Password string
type HashedPassword string

type doorman struct {
	tokenPersistence TokenPersistence
}

func (d doorman) ValidateToken(tk Token) (UserName, bool) {
	u := d.tokenPersistence.GetUserName(tk)
	if u == "" {
		return UserName(""), false
	}
	return u, true
}
