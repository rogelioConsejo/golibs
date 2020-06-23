package dbdef

import (
	"database/sql"
	"fmt"
	"github.com/rogelioConsejo/golibs/persistencia"
	"strings"
)

type DatabaseImage struct {
	name string
	tables []*persistencia.DefinicionTabla
}

func GetDbDefinition(conn *persistencia.CredencialesSQL, name string) (definition *DatabaseImage, err error) {
	var db *sql.DB
	db, err = persistencia.ConectarMySQL(*conn, name)
	var tables *sql.Rows
	if err == nil {
		tables, err = db.Query("SHOW TABLES;")
	}
	if err == nil {
		var dbImage = new(DatabaseImage)
		dbImage.name = name

		for tables.Next() {
			var table = new(persistencia.DefinicionTabla)
			var tableName string
			err = tables.Scan(&tableName)

			if err == nil {
				table.Nombre = tableName
				println(tableName + " ")
				table.Campos, err = getTableFields(db, tableName)
			}
			if err == nil{
				dbImage.tables = append(dbImage.tables, table)
			}
		}
		definition = dbImage
		fmt.Printf("%+v", *dbImage)
		fmt.Printf("%+v", dbImage.tables[0])
		fmt.Printf("%+v", dbImage.tables[1])
	}
	return
}

func getTableFields(db *sql.DB, name string) (fields map[string]string, err error) {
	fields = make(map[string]string)
	var queryResult *sql.Rows

	query := fmt.Sprintf("SELECT COLUMN_NAME, COLUMN_TYPE, COLUMN_KEY, IS_NULLABLE, EXTRA, COLUMN_DEFAULT " +
		"FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = '%s';", name)
	queryResult, err = db.Query(query)

	if err == nil {
		for queryResult.Next(){
			var field string
			var dataType string
			var isNull string
			var key string
			var defaultValue sql.NullString
			var extra string


			var definition string
			var sbDefinition strings.Builder

			err = queryResult.Scan(&field, &dataType,  &key, &isNull, &extra, &defaultValue)
			sbDefinition.WriteString(dataType)
			if key == "PRI" {
				sbDefinition.WriteString(" PRIMARY KEY")
			}
			if isNull == "NO" {
				sbDefinition.WriteString(" NOT NULL")
			}
			sbDefinition.WriteString(";")

			fmt.Printf("%s %s\n", field, sbDefinition.String())


			definition = sbDefinition.String()
			fields[field] = definition
		}
	}

	return
}