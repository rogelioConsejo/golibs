package persistencia

import (
"database/sql"
"fmt"
_ "github.com/go-sql-driver/mysql"
)

//CredencialesSQL: el nombre de usuario y password para conectarse a la base de datos.
type CredencialesSQL struct {
	User     string
	Password string
	Host string
}

//ConectarMySQL crea una conexión a una base de datos MySQL.
func ConectarMySQL(credenciales CredencialesSQL, nombreDeBaseDeDatos string) (*sql.DB, error) {
	datosParaConectarse := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		credenciales.User, credenciales.Password, credenciales.Host, nombreDeBaseDeDatos)

	baseDeDatos, err := sql.Open("mysql", datosParaConectarse)
	if err != nil {
		mensajeDeError := fmt.Sprintf("No se pudo conectar a la base de datos (%s)", nombreDeBaseDeDatos)
		return nil, fmt.Errorf(mensajeDeError)
	}

	return baseDeDatos, nil
}
