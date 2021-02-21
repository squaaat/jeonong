package store

import (
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/squaaat/nearsfeed/api/internal/app"
	"github.com/squaaat/nearsfeed/api/internal/er"
	keywordStore "github.com/squaaat/nearsfeed/api/internal/service/keyword/store"
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
		k, err := keywordStore.MustGetKeyword(s.App.ServiceDB.DB, cat.Name, cat.Code)
		if err != nil {
			return er.WrapOp(err, op)
		}
		c, err := AddCategoryIfNotExist(s.App.ServiceDB.DB, k, nil)
		if err != nil {
			return er.WrapOp(err, op)
		}
		log.Debug().Interface("category", c).Send()
	}
	for _, cat := range data.Depth2 {
		k, err := keywordStore.MustGetKeyword(s.App.ServiceDB.DB, cat.Name, cat.Code)
		if err != nil {
			return er.WrapOp(err, op)
		}
		pk, err := keywordStore.GetKeywordByCode(s.App.ServiceDB.DB, cat.ParentCode)
		if err != nil {
			return er.WrapOp(err, op)
		}
		fmt.Println(k.Code, pk.Code)
		c, err := AddCategoryIfNotExist(s.App.ServiceDB.DB, k, pk)
		if err != nil {
			return er.WrapOp(err, op)
		}
		log.Debug().Interface("category", c).Send()
	}
	for _, cat := range data.Depth3 {
		k, err := keywordStore.MustGetKeyword(s.App.ServiceDB.DB, cat.Name, cat.Code)
		if err != nil {
			return er.WrapOp(err, op)
		}
		pk, err := keywordStore.GetKeywordByCode(s.App.ServiceDB.DB, cat.ParentCode)
		if err != nil {
			return er.WrapOp(err, op)
		}
		c, err := AddCategoryIfNotExist(s.App.ServiceDB.DB, k, pk)
		if err != nil {
			return er.WrapOp(err, op)
		}
		log.Debug().Interface("category", c).Send()
	}

	return nil
}
