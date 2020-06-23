package dbdef

import (
	"github.com/rogelioConsejo/golibs/persistencia"
	"testing"
)

func TestGetDbDefinition(t *testing.T)  {
	var credenciales = persistencia.CredencialesSQL{
		User:     "root",
		Password: "",
		Host:     "localhost:3306",
	}
	GetDbDefinition(&credenciales, "ejemplo")
}
