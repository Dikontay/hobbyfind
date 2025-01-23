package migration

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/cassandra"
	_ "github.com/golang-migrate/migrate/v4/database/cassandra"
)

type service struct {
	configs Configs
}

func NewService(configs Configs) Service {
	return &service{
		configs: configs,
	}
}
func (s service) Start(session *gocql.Session, keyspace string, dbName string) error {
	driver, err := cassandra.WithInstance(session, &cassandra.Config{
		KeyspaceName:    keyspace,
		MigrationsTable: "db_migrations",
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
