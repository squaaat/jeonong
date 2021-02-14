package store

import (
	"github.com/rs/zerolog/log"

	"github.com/squaaat/jeonong/api/internal/app"
	"github.com/squaaat/jeonong/api/internal/er"
	keywordStore "github.com/squaaat/jeonong/api/internal/service/keyword/store"
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
		k, err := keywordStore.MustGetKeyword(s.App.ServiceDB.DB, man.Name, man.Code)
		if err != nil {
			return er.WrapOp(err, op)
		}
		c, err := AddManufactureIfNotExist(s.App.ServiceDB.DB, k, man.CompanyRegistrationNumber)
		if err != nil {
			return er.WrapOp(err, op)
		}
		log.Debug().Interface("category", c).Send()
	}
	return nil
}
