package persistencia

import (
	"fmt"
	"testing"
)

type resultado struct {
	id      int
	palabra string
}

func TestConectarMySQL(t *testing.T) {
	user := "root"
	pass := ""
	credenciales := CredencialesSQL{
		User:     user,
		Password: pass,
	}

	baseDeDatos, err := ConectarMySQL(credenciales, "erp_v3")
	if err != nil {
		t.Errorf("Error al conectar a la base de datos: %s\n", err.Error())
	} else {
		defer func() {
			errClose := baseDeDatos.Close()
			if errClose != nil {
				t.Errorf("No se pudo cerrar la base de datos: %s", errClose.Error())
			}
		}()

		_, err = baseDeDatos.Exec("SHOW TABLES")
		if err != nil {
			t.Errorf("Error al mostrar tablas de la base de datos: %s\n", err.Error())
		}

		rows, err1 := baseDeDatos.Query("SELECT * FROM prueba")
		if err1 != nil {
			t.Errorf("Error al mostrar leer tablas de la base de datos: %s\n", err1.Error())
		} else {
			defer func() {
				errClose := rows.Close()
				if errClose != nil {
					t.Errorf("Error al cerrar el resultset: %s\n", errClose.Error())
				}
			}()

			var r resultado

			for rows.Next() {
				if err = rows.Scan(&r.id, &r.palabra); err != nil {
					t.Errorf("No se puedieron leer los resultados: %s", err.Error())
				}

				fmt.Printf("palabra: %s\n", r.palabra)
			}

		}

	}
}
