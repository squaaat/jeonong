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

	for _, cat := range data.Depth1 {
		c, err := AddCategoryIfNotExist(s.App.ServiceDB.DB, cat, 1)
		if err != nil {
			return er.WrapOp(err, op)
		}
		log.Debug().Interface("category", c).Send()
	}
	for _, cat := range data.Depth2 {
		pc, err := CategoryIdsByKeywordCode(s.App.ServiceDB.DB, cat.ParentCode)
		if err != nil {
			return er.WrapOp(err, op)
		}

		c, err := AddCategoryIfNotExist(s.App.ServiceDB.DB, cat, 2, pc...)
		if err != nil {
			return er.WrapOp(err, op)
		}
		log.Debug().Interface("category", c).Send()
	}
	for _, cat := range data.Depth3 {
		pc, err := CategoryIdsByKeywordCode(s.App.ServiceDB.DB, cat.ParentCode)
		if err != nil {
			return er.WrapOp(err, op)
		}

		c, err := AddCategoryIfNotExist(s.App.ServiceDB.DB, cat, 3, pc...)
		if err != nil {
			return er.WrapOp(err, op)
		}
		log.Debug().Interface("category", c).Send()
	}

	return nil
}
