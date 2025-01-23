package cassandra

import (
	"project/internal/services/cassandra/migration"
	"project/internal/services/cassandra/session"
)

type Configs struct {
	Session   session.Configs   `json:"session"`
	Migration migration.Configs `json:"migration"`
}
