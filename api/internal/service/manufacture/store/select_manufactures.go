package store

import (
	"github.com/squaaat/nearsfeed/api/internal/er"
	"github.com/squaaat/nearsfeed/api/internal/model"
)

func (s *Service) SelectManufactures() ([]*model.Manufacture, error) {
	op := er.CallerOp()

	m := &model.Manufacture{}
	tx := s.App.ServiceDB.DB.
		Model(m).
		Where("status = ?", model.StatusIdle).
		Order("name ASC")

	if tx.Error != nil {
		return nil, er.WrapOp(tx.Error, op)
	}

	rows, err := tx.Rows()
	defer rows.Close()
	if err != nil {
		return nil, er.WrapOp(err, op)
	}

	var r []*model.Manufacture
	for rows.Next() {
		item := new(model.Manufacture)
		err = tx.ScanRows(rows, &item)
		if err != nil {
			return nil, er.WrapOp(err, op)
		}
		r = append(r, item)
	}

	return r, nil
}
