package store

import (
	"github.com/squaaat/jeonong/api/internal/app"
	"github.com/squaaat/jeonong/api/internal/er"
	"github.com/squaaat/jeonong/api/internal/model"
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
		//
		result := s.App.ServiceDB.DB.Create(&model.Keyword{
			Name: cat.Name,
			Code: cat.Code,
		})
		if result.Error != nil {
			err = er.WrapOp(result.Error, op)
			return result.Error
		}
		if result.RowsAffected == 0 {
			return er.New("Nothing inserted at rows", er.KindInternalServerError, op)
		}
	}

	return nil
}
