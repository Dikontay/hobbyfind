package connection

import (
	"github.com/uptrace/bun"
)

type Service interface {
	GetClient(args ...int) (*bun.DB, error)
}
