package persistencia

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

//FUNCIONES PRIVADAS
func conectarABaseDeDatos() (*sql.DB, error) {
	var err error

	credenciales := CredencialesSQL{
		User:     DBUser,
		Password: DBPass,
	}
	BaseDeDatos, errConexion := ConectarMySQL(credenciales, DBName)
	if errConexion != nil {
		mensajeDeError := fmt.Sprintf("error al conectar a la base de datos para registrar una moneda:\n %s\n", errConexion.Error())
		err = errors.New(mensajeDeError)
	}
	return BaseDeDatos, err
}

func cerrarConexion(BaseDeDatos *sql.DB) {
	errCierreConexion := BaseDeDatos.Close()
	if errCierreConexion != nil {
		fmt.Printf("No se pudo cerrar la conexión: %s\n", errCierreConexion.Error())
	}
}

func revisarValorDeId(id uint) error {
	var err error

	if id <= 0 {
		strError := fmt.Sprintf("valor de id inválido (%v), debe ser mayor a 0\n", id)
		err = errors.New(strError)
	}
	return err
}

func extraerCamposYValores(obj Entity) (string, string) {
	campos := reflect.TypeOf(obj)
	valores := reflect.ValueOf(obj)

	strCampos, strValores := extraerCamposYValoresAStrings(campos, valores)
	camposTruncado := quitarComaAlFinal(strCampos)
	valoresTruncado := quitarComaAlFinal(strValores)
	return camposTruncado, valoresTruncado
}

func extraerCamposYValoresAStrings(campos reflect.Type, valores reflect.Value) (string, string) {
	var strCampos string
	var strValores string

	for i := 0; i < campos.NumField(); i++ {
		campo := campos.Field(i)
		valor := valores.Field(i)
		tipo := campo.Type

		if !(valor == reflect.New(tipo) || fmt.Sprint(valor) == "{ false}" || fmt.Sprint(valor) == "[]" || campo.Name == "Id") {
			strCampos = strCampos + fmt.Sprint(campo.Name) + ", "
			strValores = strValores + "'" + fmt.Sprint(valor) + "', "
		}
	}
	return strCampos, strValores
}

func extraerTodosLosCampos(entity Entity) string {
	campos := reflect.TypeOf(entity)
	var strCampos strings.Builder

	for i := 0; i < campos.NumField(); i++ {
		campo := campos.Field(i)

		strCampos.WriteString(fmt.Sprint(campo.Name) + ", ")
	}
	camposTruncado := quitarComaAlFinal(strCampos.String())
	return camposTruncado
}

func parsearCondicionesDeBusqueda(entity Entity, operadorLogico string) string {
	var condiciones string

	campos, valores := extraerCamposYValores(entity)

	arrCampos := strings.Split(campos, ", ")
	arrValores := strings.Split(valores, ", ")

	for indice, campo := range arrCampos {
		if arrValores[indice] != "" && arrValores[indice] != "''" && arrValores[indice] != "0" && arrValores[indice] != "'0'" && arrValores[indice] != "'[]'" {
			condiciones += fmt.Sprintf("%s = %s %s ", campo, arrValores[indice], operadorLogico)
		}
	}

	condiciones = quitarOperadorAlFinal(condiciones, operadorLogico)

	return condiciones
}

func parsearNombresDeColumna(entity Entity) string {
	campos := extraerTodosLosCampos(entity)

	return campos
}

func parsearMapa2CamposYValores(mapa map[string]string) (err error, campos string, valores string) {
	var strbCampos, strbValores strings.Builder
	for campo, valor := range mapa {
		_, err = strbCampos.WriteString(campo + ", ")
		_, err = strbValores.WriteString("'" + valor + "', ")
	}
	if err == nil {
		campos = strbCampos.String()
		campos = quitarComaAlFinal(campos)
		valores = strbValores.String()
		valores = quitarComaAlFinal(valores)
	}
	return err, campos, valores
}

func parsearMapa2Condiciones(mapa map[string]string) string {
	var strbCondiciones strings.Builder

	for columna, valor := range mapa {
		strbCondiciones.WriteString(columna + " ='" + valor + "' " + AND + " ")
	}

	condiciones := strbCondiciones.String()
	condiciones = quitarOperadorAlFinal(condiciones, AND)
	return condiciones
}

func formatearCamposYValoresParaUpdate(campos string, valores string, ignorados []string) string {
	var query string

	fmt.Printf("ignorados: %+v\n", ignorados)
	arrCampos := strings.Split(campos, ", ")
	arrValores := strings.Split(valores, ", ")
	for indice, campo := range arrCampos {
		var ignorar bool
		for _, ignorado := range ignorados {
			if strings.ToLower(campo) == strings.ToLower(ignorado) {
				ignorar = true
			}
		}
		if !ignorar {
			query += fmt.Sprintf("%s = %s, ", campo, arrValores[indice])

		}
	}
	query = quitarComaAlFinal(query)
	return query
}

func quitarComaAlFinal(cadena string) string {
	cadenaRunes := []rune(cadena)
	cadenaTruncada := cadenaRunes[:len(cadenaRunes)-2]

	return string(cadenaTruncada)
}

func quitarOperadorAlFinal(cadena string, operador string) string {
	if cadena != "" {
		cadenaRunes := []rune(cadena)
		operadorRunes := []rune(operador)
		condicionesRunesTruncada := cadenaRunes[:len(cadenaRunes)-(len(operadorRunes)+1)]
		cadena = string(condicionesRunesTruncada)
	}

	return cadena
}
