package requester

import (
	"github.com/Dikontay/hobbyfind/entities"
)

type Service interface {
	CreateUser(user entities.User) (*entities.User, error)
}
