package create

import (
	"github.com/Dikontay/hobbyfind/repository/internal/domain/models"
)

type Params struct {
	models.Account
}

func (p Params) Validate() error {
	return nil
}
