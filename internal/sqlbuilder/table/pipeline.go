//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var Pipeline = newPipelineTable("public", "pipeline", "")

type pipelineTable struct {
	postgres.Table

	// Columns
	TenantID        postgres.ColumnString
	PipelineID      postgres.ColumnString
	Name            postgres.ColumnString
	SourceID        postgres.ColumnString
	DestinationID   postgres.ColumnString
	Config          postgres.ColumnString
	CreatedByUserID postgres.ColumnString
	UpdatedByUserID postgres.ColumnString
	CreatedAt       postgres.ColumnTimestamp
	UpdatedAt       postgres.ColumnTimestamp

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
	DefaultColumns postgres.ColumnList
}

type PipelineTable struct {
	pipelineTable

	EXCLUDED pipelineTable
}

// AS creates new PipelineTable with assigned alias
func (a PipelineTable) AS(alias string) *PipelineTable {
	return newPipelineTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new PipelineTable with assigned schema name
func (a PipelineTable) FromSchema(schemaName string) *PipelineTable {
	return newPipelineTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new PipelineTable with assigned table prefix
func (a PipelineTable) WithPrefix(prefix string) *PipelineTable {
	return newPipelineTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new PipelineTable with assigned table suffix
func (a PipelineTable) WithSuffix(suffix string) *PipelineTable {
	return newPipelineTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newPipelineTable(schemaName, tableName, alias string) *PipelineTable {
	return &PipelineTable{
		pipelineTable: newPipelineTableImpl(schemaName, tableName, alias),
		EXCLUDED:      newPipelineTableImpl("", "excluded", ""),
	}
}

func newPipelineTableImpl(schemaName, tableName, alias string) pipelineTable {
	var (
		TenantIDColumn        = postgres.StringColumn("tenant_id")
		PipelineIDColumn      = postgres.StringColumn("pipeline_id")
		NameColumn            = postgres.StringColumn("name")
		SourceIDColumn        = postgres.StringColumn("source_id")
		DestinationIDColumn   = postgres.StringColumn("destination_id")
		ConfigColumn          = postgres.StringColumn("config")
		CreatedByUserIDColumn = postgres.StringColumn("created_by_user_id")
		UpdatedByUserIDColumn = postgres.StringColumn("updated_by_user_id")
		CreatedAtColumn       = postgres.TimestampColumn("created_at")
		UpdatedAtColumn       = postgres.TimestampColumn("updated_at")
		allColumns            = postgres.ColumnList{TenantIDColumn, PipelineIDColumn, NameColumn, SourceIDColumn, DestinationIDColumn, ConfigColumn, CreatedByUserIDColumn, UpdatedByUserIDColumn, CreatedAtColumn, UpdatedAtColumn}
		mutableColumns        = postgres.ColumnList{NameColumn, SourceIDColumn, DestinationIDColumn, ConfigColumn, CreatedByUserIDColumn, UpdatedByUserIDColumn, CreatedAtColumn, UpdatedAtColumn}
		defaultColumns        = postgres.ColumnList{}
	)

	return pipelineTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		TenantID:        TenantIDColumn,
		PipelineID:      PipelineIDColumn,
		Name:            NameColumn,
		SourceID:        SourceIDColumn,
		DestinationID:   DestinationIDColumn,
		Config:          ConfigColumn,
		CreatedByUserID: CreatedByUserIDColumn,
		UpdatedByUserID: UpdatedByUserIDColumn,
		CreatedAt:       CreatedAtColumn,
		UpdatedAt:       UpdatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
		DefaultColumns: defaultColumns,
	}
}
