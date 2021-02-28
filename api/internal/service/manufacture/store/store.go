package store

import (
	"github.com/squaaat/nearsfeed/api/internal/app"
)

type Service struct {
	App *app.Application
}

func New(a *app.Application) *Service {
	return &Service{
		App: a,
	}
}
