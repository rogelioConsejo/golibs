package persistencia

import (
	"database/sql"
	"errors"
	"fmt"
)

const AND = "AND"
const OR = "OR"

type Entity interface {
	GetId() uint
}

func RegistrarEnBaseDeDatos(obj Entity, tabla string) (uint, error) {
	var id int64
	var err error
	var resultado sql.Result
	baseDeDatos, err := conectarABaseDeDatos()
	defer cerrarConexion(baseDeDatos)

	campos, valores := extraerCamposYValores(obj)

	if err == nil {
		query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", tabla, campos, valores)
		resultado, err = baseDeDatos.Exec(query)
	}

	if err == nil {
		id, err = resultado.LastInsertId()
	}

	return uint(id), err
}

func BuscarUnoEnBaseDeDatos(obj Entity, tabla string) (*sql.Row, error) {
	var entrada *sql.Row

	baseDeDatos, err := conectarABaseDeDatos()
	defer cerrarConexion(baseDeDatos)

	condiciones := parsearCondicionesDeBusqueda(obj, AND)

	if condiciones == "" {
		err = errors.New("parámetros de búsqueda inválidos")
	}

	columnas := parsearNombresDeColumna(obj)

	if err == nil && condiciones != "" {
		query := fmt.Sprintf("SELECT %s FROM %s WHERE %s;", columnas, tabla, condiciones)
		entrada = baseDeDatos.QueryRow(query)
	}



	return entrada, err
}

func BuscarEnBaseDeDatos(obj Entity, tabla string) (*sql.Rows, error) {
	var entradas *sql.Rows

	baseDeDatos, err := conectarABaseDeDatos()
	defer cerrarConexion(baseDeDatos)

	condiciones := parsearCondicionesDeBusqueda(obj, OR)
	columnas := parsearNombresDeColumna(obj)

	if err == nil {
		query := fmt.Sprintf("SELECT %s FROM %s WHERE %s;", columnas, tabla, condiciones)
		entradas, err = baseDeDatos.Query(query)
	}

	return entradas, err
}

func LeerEnBaseDeDatos(id uint, modelo Entity, tabla string) (*sql.Row, error) {
	var entrada *sql.Row

	baseDeDatos, err := conectarABaseDeDatos()
	defer cerrarConexion(baseDeDatos)

	columnas := parsearNombresDeColumna(modelo)

	if err == nil {
		query := fmt.Sprintf("SELECT %s FROM %s WHERE Id = %v;", columnas, tabla, id)
		entrada = baseDeDatos.QueryRow(query)

	}

	return entrada, err
}

func ActualizarRegistroEnBaseDeDatos(obj Entity, tabla string, ignorados ...string) error {
	baseDeDatos, err := conectarABaseDeDatos()
	defer cerrarConexion(baseDeDatos)

	err = revisarValorDeId(obj.GetId())

	if err == nil {
		campos, valores := extraerCamposYValores(obj)
		query := fmt.Sprintf("UPDATE %s SET ", tabla)
		query += formatearCamposYValoresParaUpdate(campos, valores, ignorados)
		query += fmt.Sprintf(" WHERE id = %v;", obj.GetId())
		println(query)

		_, err = baseDeDatos.Exec(query)
	}

	return err
}

func BorrarEnBaseDeDatos(id uint, tabla string) error {
	baseDeDatos, err := conectarABaseDeDatos()
	defer cerrarConexion(baseDeDatos)

	err = revisarValorDeId(id)
	if err == nil {
		query := fmt.Sprintf("DELETE FROM %s WHERE Id=%v;", tabla, id)
		_, err = baseDeDatos.Exec(query)
	}

	return err
}

//TODO
func RegistrarMapEnBaseDeDatos(mapa map[string]string, tabla string) (id uint, err error) {
	var baseDeDatos *sql.DB
	var resultado sql.Result
	var id64 int64
	var campos, valores string

	baseDeDatos, err = conectarABaseDeDatos()
	defer cerrarConexion(baseDeDatos)

	if len(mapa) == 0 {
		err = errors.New("no hay datos para registrar")
	}

	if err == nil {
		err, campos, valores = parsearMapa2CamposYValores(mapa)
	}

	if err == nil {
		query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", tabla, campos, valores)
		resultado, err = baseDeDatos.Exec(query)
	}

	if err == nil {
		id64, err = resultado.LastInsertId()
	}

	id = uint(id64)
	return id, err
}

//TODO
func BuscarMapaEnBaseDeDatos(mapa map[string]string, tabla string) (id uint, err error) {
	var baseDeDatos *sql.DB

	baseDeDatos, err = conectarABaseDeDatos()
	defer cerrarConexion(baseDeDatos)

	condiciones := parsearMapa2Condiciones(mapa)

	if condiciones == "" {
		err = errors.New("parámetros de búsqueda inválidos")
	}

	if err == nil && condiciones != "" {
		query := fmt.Sprintf("SELECT id FROM %s WHERE %s;", tabla, condiciones)
		entrada := baseDeDatos.QueryRow(query)
		err = entrada.Scan(&id)
	}



	return id, err
}


