package migration

import "database/sql"

type Service interface {
	Init(db *sql.DB, dbName string, schemaName string) error
}
