package category

import (
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/squaaat/nearsfeed/api/internal/app"
	"github.com/squaaat/nearsfeed/api/internal/er"
	"github.com/squaaat/nearsfeed/api/internal/model"
	categoryStore "github.com/squaaat/nearsfeed/api/internal/service/category/store"
)

type Service struct {
	App           *app.Application
	CategoryStore *categoryStore.Service
}

func New(a *app.Application) *Service {
	return &Service{
		App:           a,
		CategoryStore: categoryStore.New(a),
	}
}

type In struct {
	Category *model.Category
}

type Out struct {
	Category *model.Category
	Categories []*model.Category
}

func (s *Service) PutCategory(mc *model.Category) (*Out, error) {
	op := er.CallerOp()

	mc.Status = model.StatusIdle
	c, err := s.CategoryStore.AddCategoryOnlyExist(mc)
	if err != nil {
		return nil, er.WrapOp(err, op)
	}

	return &Out{
		Category: c,
	}, nil
}

func (s *Service) GetCategories() (*Out, error) {
	op := er.CallerOp()

	var categories []*model.Category
	table := &model.Category{}
	s.App.ServiceDB.DB.Name()
	tx := s.App.ServiceDB.DB.
		Model(table).
		Joins("Keyword").
		Joins("ParentKeyword").
		Where(fmt.Sprintf("%s.status = 'IDLE'", table.TableName()))
	if tx.Error != nil {
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
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
	return &Out{
		Categories: categories,
	}, nil

}
