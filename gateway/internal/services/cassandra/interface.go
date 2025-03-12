package cassandra

import "github.com/Dikontay/hobbyfind/gateway/internal/domain/models"

type Service interface {
	Init() error

	CreateUser(user models.User) (*models.User, error)
	ReadUser(id string) (*models.User, error)
	UpdateUser(user models.User) (*models.User, error)
	DeleteUser(id string) error
	ListUsers(user models.User) ([]*models.User, error)

	Stop()
}
