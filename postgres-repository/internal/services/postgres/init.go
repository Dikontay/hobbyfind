package postgres

import (
	"fmt"
	"github.com/Dikontay/hobbyfind/repository/internal/services/postgres/connection"
	"github.com/Dikontay/hobbyfind/repository/internal/services/postgres/migration"
)

var cn connection.Service
var mg migration.Service

func ConnectionService() connection.Service {
	return cn
}

func MigrationService() migration.Service {
	return mg
}

func Init(configs Configs) error {
	cn = connection.NewService(configs.Connection)
	mg = migration.NewService(configs.Migration)

	db, err := cn.GetClient()
	if err != nil {
		return fmt.Errorf("failed to get connection: %v", err)
	}

	err = mg.Init(db.DB, configs.Connection.DB, configs.Connection.Schema)

	if err != nil {
		return fmt.Errorf("failed to create db %v", err)
	}
	return nil
}
