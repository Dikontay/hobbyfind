package read

import "project/internal/domain/models"

type Response struct {
	models.User `json:"user"`
}
