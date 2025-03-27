package login

import (
	"fmt"
	"regexp"
)

type Params struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (p *Params) isValidEmail() bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(p.Email)
}
func (p *Params) Validate() error {

	if p.isValidEmail() {
		return fmt.Errorf("email is required")
	}

	if len(p.Password) == 0 {
		return fmt.Errorf("passwrod is required")
	}

	return nil
}
