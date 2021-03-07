package store

import (
	"github.com/squaaat/nearsfeed/api/internal/er"
	"github.com/squaaat/nearsfeed/api/internal/model"
)

func (s *Service) SelectOneManufacture(id string) (*model.Manufacture, error) {
	op := er.CallerOp()

	m := &model.Manufacture{}
	tx := s.C.ServiceDB.DB.
		Model(m).
		Where("id = ?", id).
		Scan(m)

	if tx.Error != nil {
		return nil, er.WrapOp(tx.Error, op)
	}

	return m, nil
}
