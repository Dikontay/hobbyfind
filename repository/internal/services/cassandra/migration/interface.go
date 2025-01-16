package migration

type Service interface {
	Init()

	Stop()
}
