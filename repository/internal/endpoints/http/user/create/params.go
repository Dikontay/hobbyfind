package create

import (
	"errors"
	"github.com/Dikontay/hobbyfind/repository/internal/domain/consts"
	"github.com/Dikontay/hobbyfind/repository/internal/domain/models"
)

type Params struct {
	models.User
}

func (p Params) Validate() error {
	if p.Fullname == "" {
		return errors.New(consts.ErrFullnameRequired)
	}
	if len(p.Password) < 8 {
		return errors.New(consts.ErrPasswordTooShort)
	}
	if len(p.Username) < 3 {
		return errors.New(consts.ErrUsernameTooShort)
	}
	if len(p.Email) == 0 {
		return errors.New(consts.ErrEmailRequired)
	}
	if len(p.Phone) == 0 {
		return errors.New(consts.ErrPhoneRequired)
	}
	return nil
}
