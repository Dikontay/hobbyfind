package create

import (
	"github.com/Dikontay/hobbyfind/repository/internal/domain/models"
)

type Response struct {
	Account models.Account `json:"Account"`
}
