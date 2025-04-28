package services

import (
	"fmt"
	"github.com/Dikontay/hobbyfind/repository/internal/services/postgres"
)

func Init(configs Configs) error {
	// Init Cassandra

	err := postgres.Init(configs.Postgres)
	if err != nil {
		return fmt.Errorf("failed to init cassandra: %w", err)
	}
	return nil
}
