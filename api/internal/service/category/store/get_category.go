package store

import (
	"github.com/squaaat/nearsfeed/api/internal/er"
	"github.com/squaaat/nearsfeed/api/internal/model"
)

func (s *Service) GetCategory(id string) (*model.Category, error) {
	op := er.CallerOp()

	mc := &model.Category{DefaultModel: model.DefaultModel{ID: id}}
	tx := s.App.ServiceDB.DB.First(mc, "id = ? AND status = ?", mc.ID, model.StatusIdle)

	if tx.Error != nil {
		return nil, er.WrapKindAndOp(tx.Error, er.KindNotFound, op)
	}

	return mc, nil
}
