package store

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/squaaat/nearsfeed/api/internal/er"
	"github.com/squaaat/nearsfeed/api/internal/model"
)

func (s *Service) SelectCategory() ([]*model.Category, error) {
	op := er.CallerOp()

	var categories []*model.Category
	tableName := (&model.Category{}).TableName()
	tx := s.C.ServiceDB.DB.
		Table(tableName).
		Where("status = ?", model.StatusIdle).
		Order("depth ASC").
		Order("sort ASC").
		Order("name ASC")
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return categories, nil
		} else {
			return nil, er.WrapOp(tx.Error, op)
		}
	}
	rows, err := tx.Rows()
	defer rows.Close()
	if err != nil {
		return nil, er.WrapOp(err, op)
	}

	for rows.Next() {
		c := new(model.Category)
		err = tx.ScanRows(rows, &c)
		if err != nil {
			return nil, er.WrapOp(err, op)
		}
		categories = append(categories, c)
	}
	return categories, nil
}
