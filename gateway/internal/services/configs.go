package services

import "project/internal/services/cassandra"

type Configs struct {
	Cassandra cassandra.Configs `json:"cassandra"`
}
