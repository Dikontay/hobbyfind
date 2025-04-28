package read

import (
	"fmt"
	"github.com/Dikontay/hobbyfind/repository/utils"
)

type Params struct {
	Id string `json:"-" params:"id"`
}

func (p *Params) Validate() error {
	if !utils.CheckUuid(p.Id) {
		return fmt.Errorf("invalid id: %s", p.Id)
	}
	return nil
}
