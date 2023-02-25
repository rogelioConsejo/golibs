package login

import "testing"

func TestDoorman(t *testing.T) {
	var per CredentialsPersistence = stubCredentialsPersistence{}
	const u UserName = "user"
	const p Password = "password"
	var reg UserRegistry = userRegistry{credentialsPersistence: per}
	reg.AddUser(u, p)

	var tper TokenPersistence = &fakeTokenPersistence{}
	var tkm TokenMaker = tokenMaker{credentialsPersistence: per, tokenPersistence: tper}

	var tk Token = tkm.GenerateToken(u, p)

	var d Doorman = doorman{tokenPersistence: tper}
	var userName UserName
	var isValid bool
	userName, isValid = d.ValidateToken(tk)
	if !isValid {
		t.Error("Token is not valid")
	}
	if userName != u {
		t.Error("Token is not valid")
	}
}

type stubCredentialsPersistence struct {
}

func (f stubCredentialsPersistence) CheckCredentials(name UserName, password HashedPassword) bool {
	return true
}

func (f stubCredentialsPersistence) Save(user UserName, password HashedPassword) {

}

type fakeTokenPersistence struct {
	tokens map[Token]UserName
}

func (f *fakeTokenPersistence) GetUserName(tk Token) UserName {
	return f.tokens[tk]
}

func (f *fakeTokenPersistence) Save(name UserName, token Token) {
	if f.tokens == nil {
		f.tokens = make(map[Token]UserName)
	}
	f.tokens[token] = name
}
