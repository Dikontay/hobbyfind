package create

import "project/internal/domain/models"

type Response struct {
	User models.User `json:"user"`
}
