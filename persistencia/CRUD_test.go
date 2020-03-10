package persistencia

import "testing"

func TestBuscarMapaEnBaseDeDatos(t *testing.T) {
	tabla := "Usuarios"
	usuario := map[string]string{
		"nombre": "Usuario de Prueba",
	}

	id, err := RegistrarMapEnBaseDeDatos(usuario, tabla)

	if err != nil {
		t.Errorf("error al registrar el mapa en la Base de Datos: %s\n", err.Error())
	} else {
		id, err = BuscarMapaEnBaseDeDatos(usuario, tabla)

		if err != nil {
			t.Errorf("error al buscar mapa en Base de Datos: %s\n", err.Error())
		} else {
			err = BorrarEnBaseDeDatos(id, tabla)
		}

		if err != nil {
			t.Errorf("error al borrar el mapa de la Base de Datos: %s\n", err.Error())
		}
	}
}
