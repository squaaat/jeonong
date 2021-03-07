package store

import (
	"github.com/squaaat/nearsfeed/api/internal/container"
)

type Service struct {
	C *container.Container
}

func New(c *container.Container) *Service {
	return &Service{
		C: c,
	}
}
