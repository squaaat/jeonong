package category

import (
	"github.com/squaaat/nearsfeed/api/internal/container"
	"github.com/squaaat/nearsfeed/api/internal/er"
	"github.com/squaaat/nearsfeed/api/internal/model"
	categoryStore "github.com/squaaat/nearsfeed/api/internal/service/category/store"
)

type Service struct {
	C             *container.Container
	CategoryStore *categoryStore.Service
}

func New(c *container.Container) *Service {
	return &Service{
		C:             c,
		CategoryStore: categoryStore.New(c),
	}
}

type In struct {
	Category *model.Category
}

type Out struct {
	Category   *model.Category
	Categories []*model.Category
}

func (s *Service) PutCategory(mc *model.Category) (*Out, error) {
	op := er.CallerOp()

	mc.Status = model.StatusIdle
	c, err := s.CategoryStore.InsertCategoryOnlyNotExist(mc)
	if err != nil {
		return nil, er.WrapOp(err, op)
	}

	return &Out{
		Category: c,
	}, nil
}

func (s *Service) GetCategories() (*Out, error) {
	op := er.CallerOp()

	categories, err := s.CategoryStore.SelectCategory()
	if err != nil {
		return nil, er.WrapOp(err, op)
	}
	return &Out{
		Categories: categories,
	}, nil
}

func (s *Service) GetCategory(id string) (*Out, error) {
	op := er.CallerOp()

	c, err := s.CategoryStore.SelectOneCategory(id)
	if err != nil {
		return nil, er.WrapOp(err, op)
	}
	return &Out{
		Category: c,
	}, nil
}
