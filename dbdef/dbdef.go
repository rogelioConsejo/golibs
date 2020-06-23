package dbdef

import "github.com/rogelioConsejo/golibs/persistencia"

type Connection struct {
	DbUserName string
	DbPassword string
	DbHost     string
}

func GetDbDefinition(conn *Connection, name string) (definition *persistencia.DefinicionTabla) {
	//TODO
	return
}