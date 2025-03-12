package services

import "github.com/Dikontay/hobbyfind/repository/internal/services/cassandra"

var cass cassandra.Service

func Repository() cassandra.Service {
	return cass
}
