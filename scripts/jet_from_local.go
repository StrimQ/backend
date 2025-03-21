package main

import (
	"github.com/go-jet/jet/v2/generator/metadata"
	"github.com/go-jet/jet/v2/generator/postgres"
	"github.com/go-jet/jet/v2/generator/template"
	postgres2 "github.com/go-jet/jet/v2/postgres"
	_ "github.com/lib/pq"
)

var dbConnection = postgres.DBConnection{
	Host:       "localhost",
	Port:       5432,
	User:       "postgresql",
	Password:   "strimqadmin_1234",
	DBName:     "postgres",
	SchemaName: "public",
	SslMode:    "disable",
}

func main() {
	const newModelPath = "../../models"
	const newSQLBuilderPath = "../../sqlbuilder"

	err := postgres.Generate(
		"./internal",
		dbConnection,
		template.Default(postgres2.Dialect).
			UseSchema(func(schemaMetaData metadata.Schema) template.Schema {
				return template.DefaultSchema(schemaMetaData).
					UseModel(template.DefaultModel().UsePath(newModelPath)).
					UseSQLBuilder(template.DefaultSQLBuilder().UsePath(newSQLBuilderPath))
			}),
	)
	if err != nil {
		panic(err)
	}
}
