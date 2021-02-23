package store

import (
	"github.com/rs/zerolog/log"

	"github.com/squaaat/nearsfeed/api/internal/app"
	"github.com/squaaat/nearsfeed/api/internal/er"
)

type Service struct {
	App *app.Application
}

func New(a *app.Application) *Service {
	return &Service{
		App: a,
	}
}

func (s *Service) MustLoadDataAtLocal() error {
	op := er.CallerOp()

	data, err := MustLoadDataAtLocal()
	if err != nil {
		err = er.WrapOp(err, op)
		return err
	}

	for _, man := range data.Manufactures {
		c, err := AddManufactureIfNotExist(s.App.ServiceDB.DB, man)
		if err != nil {
			return er.WrapOp(err, op)
		}
		log.Debug().Interface("category", c).Send()
	}
	return nil
}
