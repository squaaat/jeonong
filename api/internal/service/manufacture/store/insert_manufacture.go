package store

import (
	"gorm.io/gorm"

	"github.com/squaaat/nearsfeed/api/internal/er"
	"github.com/squaaat/nearsfeed/api/internal/model"
)

func (s *Service) insertManufacture(m *model.Manufacture) (*model.Manufacture, error) {
	op := er.CallerOp()

	m.DefaultModel.Status = model.StatusIdle
	tx := s.App.ServiceDB.DB.Create(m).Scan(m)
	if tx.Error != nil {
		return nil, er.WrapOp(tx.Error, op)
	}
	if tx.RowsAffected != 1 {
		return nil, er.New("failed create manufacture", er.KindInternalServerError, op)
	}

	return m, nil
}

func (s *Service) InsertManufactureIfNotCategory(m *model.Manufacture) (*model.Manufacture, error) {
	op := er.CallerOp()

	subM, err := GetManufactureByModel(s.App.ServiceDB.DB, m)
	if err != nil {
		if er.Is(err, gorm.ErrRecordNotFound) {
			subM, err = s.insertManufacture(m)
			if err != nil {
				return nil, er.WrapOp(err, op)
			}
			return subM, nil
		}
		return nil, er.WrapOp(err, op)
	}
	subM, err = s.insertManufacture(m)
	if err != nil {
		return nil, er.WrapOp(err, op)
	}
	return subM, nil
}

func (s *Service) InsertManufactureOnlyNotExist(m *model.Manufacture) (*model.Manufacture, error) {
	op := er.CallerOp()

	subM, err := GetManufactureByModel(s.App.ServiceDB.DB, m)
	if err != nil {
		if er.Is(err, gorm.ErrRecordNotFound) {
			subM, err := s.insertManufacture(m)
			if err != nil {
				return nil, er.WrapOp(err, op)
			}
			return subM, nil
		}
		return nil, er.WrapOp(err, op)
	}
	if subM.ID == "" {
		subM, err := s.insertManufacture(m)
		if err != nil {
			return nil, er.WrapOp(err, op)
		}
		return subM, nil
	} else {
		return nil, er.New("code, name is already exist", er.KindDubplicated, op)
	}
}
