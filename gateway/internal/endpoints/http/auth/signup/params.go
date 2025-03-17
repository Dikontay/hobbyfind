package signup

import (
	"fmt"
	"regexp"
)

type Params struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
}

func (p *Params) Validate() error {
	if len(p.Username) == 0 {
		return fmt.Errorf("username is required")
	}
	if len(p.Email) == 0 {
		return fmt.Errorf("email is required")
	}

	if len(p.Password) == 0 {
		return fmt.Errorf("passwrod is required")
	}
	if len(p.Phone) == 0 {
		return fmt.Errorf("phone number is required")
	}
	if len(p.Role) == 0 {
		return fmt.Errorf("role is required")
	}

	return nil
}

func (p *Params) isValidEmail() bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(p.Email)
}

func (p *Params) IsValidPhone() bool {
	re := regexp.MustCompile(`^\+?[0-9]{10,15}$`)
	return re.MatchString(p.Phone)
}
