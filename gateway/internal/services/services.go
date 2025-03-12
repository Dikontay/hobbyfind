package services

import "project/internal/services/cassandra"

var cass cassandra.Service

func Repository() cassandra.Service {
	return cass
}
