package store

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/squaaat/nearsfeed/api/internal/er"
	"github.com/squaaat/nearsfeed/api/internal/model"
)

func (s *Service) InsertCategory(mc *model.Category) (*model.Category, error) {
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

	subTx := s.App.ServiceDB.DB.Create(mc)
	if subTx.Error != nil {
		return nil, er.WrapOp(subTx.Error, op)
	}
	if subTx.RowsAffected != 1 {
		return nil, er.New("failed create category", er.KindInternalServerError,op)
	}

	fullName, err := s.categoryFullNameIds(mc.Depth, mc.Category1ID, mc.Category2ID, mc.Category3ID, mc.Category4ID)
	if err != nil {
		return nil, err
	}

	subTx = s.App.ServiceDB.DB.Model(mc).Update("full_name", fullName).Where("id = ?", mc.ID)
	if subTx.Error != nil {
		if !errors.Is(subTx.Error, gorm.ErrRecordNotFound) {
			return nil, er.WrapOp(subTx.Error, op)
		}
		return nil, subTx.Error
	}
	if subTx.RowsAffected != 1 {
		return nil, er.New("failed update 'full_name' a category", er.KindInternalServerError,op)
	}

	return mc, nil
}

func (s *Service) InsertCategoryIfNotExist(mc *model.Category) (*model.Category, error) {
	op := er.CallerOp()

	if c, err := GetCategoryByModel(s.App.ServiceDB.DB, mc); err != nil {
		return nil, er.WrapOp(err, op)
	} else {
		if c.ID != "" {
			return s.InsertCategory(c)
		}
		return s.InsertCategory(mc)
	}
}

func (s *Service) InsertCategoryOnlyExist(mc *model.Category) (*model.Category, error) {
	op := er.CallerOp()

	if c, err := GetCategoryByModel(s.App.ServiceDB.DB, mc); err != nil {
		return nil, er.WrapOp(err, op)
	} else {
		if c.ID != "" {
			return nil, er.New(fmt.Sprintf("[name=%s, code=%s, depth=%d] is already exists. plz check", c.Name, c.Code, c.Depth), er.KindDubplicated, op)
		}
		return s.InsertCategory(mc)
	}
}
