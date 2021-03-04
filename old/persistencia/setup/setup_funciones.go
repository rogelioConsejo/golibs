package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"strings"
)

//PRINCIPALES
func crearArchivoDeConfiguracion(userNameElegido string, passwordElegido string, BaseDeDatosAUsar string) {
	var err error
	var plantilla []byte
	var documentoGenerado []byte

	err = validarParametrosNoVacios(userNameElegido, BaseDeDatosAUsar)
	if err == nil {
		plantilla, err = ioutil.ReadFile(configTemplatePath)
	}
	if err == nil {
		documentoGenerado, err = reemplazarValoresEnPlantilla(plantilla, userNameElegido, passwordElegido, BaseDeDatosAUsar)
		err = ioutil.WriteFile(configFilePath, documentoGenerado, 0644)
		if err != nil {
			err = errors.Errorf("error al escribir el archivo: %s\n", err.Error())
		}
	}
	if err != nil {
		fmt.Printf("No se generó el archivo de configuración, existieron errores:\n%s", err)
	}
}

//AUXILIARES
func leerFlagsDeArchivoDeConfiguracion() (userNameElegido string, passwordElegido string, BaseDeDatosAUsar string){
	flag.StringVar(&userNameElegido, "u", "", "Username para acceder a la base de datos.")
	flag.StringVar(&passwordElegido, "p", "", "Password para acceder a la base de datos.")
	flag.StringVar(&BaseDeDatosAUsar, "db", "", "Base de Datos a utilizar.")

	flag.Parse()

	return userNameElegido, passwordElegido, BaseDeDatosAUsar
}

func validarParametrosNoVacios(userNameElegido string, tablaAUsar string) (err error) {
	var errorDeParametros strings.Builder

	if userNameElegido == "" {
		errorDeParametros.WriteString("no se definió el username para la conexión a la base de datos\n")
	}

	if tablaAUsar == "" {
		errorDeParametros.WriteString("no se definió la tabla para la conexión a la base de datos\n")
	}

	if errorDeParametros.String() != "" {
		err = errors.New(errorDeParametros.String())
	}
	return err
}

func reemplazarValoresEnPlantilla(plantilla []byte, userNameElegido string, passwordElegido string, tablaAUsar string) ([]byte, error) {
	var err error
	plantilla = bytes.Replace(plantilla, []byte(userNamePlaceholder), []byte(userNameElegido), -1)
	plantilla = bytes.Replace(plantilla, []byte(passwordPlaceholder), []byte(passwordElegido), -1)
	plantilla = bytes.Replace(plantilla, []byte(DBNamePlaceholder), []byte(tablaAUsar), -1)
	return plantilla, err
}