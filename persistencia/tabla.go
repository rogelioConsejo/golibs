package persistencia

import "fmt"

//STRUCT

type DefinicionTabla struct {
	Nombre string
	Campos map[string]string
}

//FUNCIONES PÚBLICAS
func CrearTabla(tabla DefinicionTabla) error {
	BaseDatos, err := conectarABaseDeDatos()
	defer cerrarConexion(BaseDatos)

	if err == nil {
		query := parsearTabla(tabla)
		_, err = BaseDatos.Exec(query)
	}

	return err
}

func BorrarTabla(nombreTabla string) error{
	baseDeDatos, err := conectarABaseDeDatos()
	defer cerrarConexion(baseDeDatos)

	if err == nil{
		query := fmt.Sprintf("DROP TABLE %s;", nombreTabla)
		_, err = baseDeDatos.Exec(query)
	}

	return err
}
