package create

import (
	"github.com/Dikontay/hobbyfind/repository/internal/domain/models"
)

type Response struct {
	User models.User `json:"user"`
}
