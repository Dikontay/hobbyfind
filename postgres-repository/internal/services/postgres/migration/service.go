package migration

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type service struct {
	configs Configs
}

func NewService(configs Configs) Service {
	return &service{
		configs: configs,
	}
}

func (s service) Init(db *sql.DB, dbName string, schemaName string) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{
		MigrationsTable: "schema_migrations",
		DatabaseName:    dbName,
		SchemaName:      schemaName,
	})
	if err != nil {
		return fmt.Errorf("failed to create driver for migration: %v", err)
	}

	migration, err := migrate.NewWithDatabaseInstance("file://migrations", dbName, driver)
	if err != nil {
		return fmt.Errorf("failed to init DB migrations. %+v", err)
	}
	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to migrate DB (Up). %+v", err)
	}
	return nil
}
