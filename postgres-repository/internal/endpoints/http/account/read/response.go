package read

import (
	"github.com/Dikontay/hobbyfind/repository/internal/domain/models"
)

type Response struct {
	models.Account `json:"Account"`
}
