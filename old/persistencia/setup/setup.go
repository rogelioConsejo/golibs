package main

const userNamePlaceholder = "{DBUserName}"
const passwordPlaceholder = "{DBPassword}"
const DBNamePlaceholder = "{DBName}"
const configTemplatePath = "../global.go.template"
const configFilePath = "../global.go"

/*Crea un archivo global.go llenando los datos en global.go.template*/

func main() {
	userNameElegido, passwordElegido, BaseDeDatosAUsar := leerFlagsDeArchivoDeConfiguracion()
	crearArchivoDeConfiguracion(userNameElegido, passwordElegido, BaseDeDatosAUsar)
}




