package category

import (
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
	c, err := s.CategoryStore.InsertCategoryOnlyExist(mc)
	if err != nil {
		return nil, er.WrapOp(err, op)
	}

	return &Out{
		Category: c,
	}, nil
}

func (s *Service) GetCategories() (*Out, error) {
	op := er.CallerOp()

	categories, err := s.CategoryStore.GetCategories()
	if err != nil {
		return nil, er.WrapOp(err, op)
	}
	return &Out{
		Categories: categories,
	}, nil
}
