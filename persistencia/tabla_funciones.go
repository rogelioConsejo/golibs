package persistencia

import "strings"

func parsearTabla(tabla DefinicionTabla) string {
	var queryBuilder strings.Builder
	queryBuilder.WriteString("CREATE TABLE " + tabla.Nombre + " (")
	for nombreColumna, definicionColumna := range tabla.Campos {
		queryBuilder.WriteString(nombreColumna + " " + definicionColumna + ", ")
	}
	queryBuilder.WriteString(")")
	query := queryBuilder.String()
	query = strings.Replace(query, ", )", ");", 1)
	return query
}