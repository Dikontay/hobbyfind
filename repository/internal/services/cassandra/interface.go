package cassandra

type Service interface {
	Init()

	Stop()
}
