package list

import (
	"github.com/Dikontay/hobbyfind/repository/internal/domain/models"
)

type Params struct {
	models.Account
	PageNum  int `json:"page_num" default:"1"`
	PageSize int `json:"page_size" default:"10"`
}

func (p *Params) Validate() error {
	if p.PageNum <= 0 {
		p.PageNum = 1
	}
	if p.PageSize <= 0 {
		p.PageSize = 10
	}
	return nil
}
