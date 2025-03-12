package services

import (
	"fmt"
	"project/internal/services/cassandra"
)

func Init(configs Configs) error {
	// Init Cassandra
	cass = cassandra.NewService(configs.Cassandra)
	err := cass.Init()
	if err != nil {
		return fmt.Errorf("failed to init cassandra: %w", err)
	}
	return nil
}
