package store

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/squaaat/nearsfeed/api/internal/er"
	"github.com/squaaat/nearsfeed/api/internal/model"
)

func (s *Service) insertCategory(mc *model.Category) (*model.Category, error) {
	op := er.CallerOp()

	if err := checkCategoryValid(mc.Depth, mc.Category1ID, mc.Category2ID, mc.Category3ID, mc.Category4ID); err != nil {
		return nil, err
	}

	mc.ID = uuid.NewString()
	if mc.Depth == 1 {
		mc.Category1ID = mc.ID
	}
	if mc.Depth == 2 {
		mc.Category2ID = mc.ID
	}
	if mc.Depth == 3 {
		mc.Category3ID = mc.ID
	}
	if mc.Depth == 4 {
		mc.Category4ID = mc.ID
	}

	tx := s.C.ServiceDB.DB.Create(mc).Scan(mc)
	if tx.Error != nil {
		return nil, er.WrapOp(tx.Error, op)
	}
	if tx.RowsAffected != 1 {
		return nil, er.New("failed create category", er.KindInternalServerError, op)
	}

	fullName, err := s.categoryFullNameIds(mc.Depth, mc.Category1ID, mc.Category2ID, mc.Category3ID, mc.Category4ID)
	if err != nil {
		return nil, err
	}

	tx = s.C.ServiceDB.DB.Model(mc).Update("full_name", fullName).Where("id = ?", mc.ID)
	if tx.Error != nil {
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, er.WrapOp(tx.Error, op)
		}
		return nil, tx.Error
	}
	if tx.RowsAffected != 1 {
		return nil, er.New("failed update 'full_name' a category", er.KindInternalServerError, op)
	}

	return mc, nil
}

func (s *Service) InsertCategoryIfNotExist(c *model.Category) (*model.Category, error) {
	op := er.CallerOp()

	subCat, err := GetCategoryByModel(s.C.ServiceDB.DB, c)
	if err != nil {
		if er.Is(err, gorm.ErrRecordNotFound) {
			subCat, err = s.insertCategory(c)
			if err != nil {
				return nil, er.WrapOp(err, op)
			}
			return subCat, nil
		}
		return nil, er.WrapOp(err, op)
	}
	subCat, err = s.insertCategory(c)
	if err != nil {
		return nil, er.WrapOp(err, op)
	}
	return subCat, nil
}

func (s *Service) InsertCategoryOnlyNotExist(c *model.Category) (*model.Category, error) {
	op := er.CallerOp()

	subCat, err := GetCategoryByModel(s.C.ServiceDB.DB, c)
	if err != nil {
		if er.Is(err, gorm.ErrRecordNotFound) {
			subCat, err := s.insertCategory(c)
			if err != nil {
				return nil, er.WrapOp(err, op)
			}
			return subCat, nil
		}
		return nil, er.WrapOp(err, op)
	}
	if subCat.ID == "" {
		subCat, err := s.insertCategory(c)
		if err != nil {
			return nil, er.WrapOp(err, op)
		}
		return subCat, nil
	} else {
		return nil, er.New("code, name is already exist", er.KindDubplicated, op)
	}
}
