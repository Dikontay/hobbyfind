package app

type Service interface {
	Start() error
	SetupRoutes()
	Stop() error
}
