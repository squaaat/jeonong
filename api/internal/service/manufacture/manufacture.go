package manufacture

import (
	"github.com/squaaat/nearsfeed/api/internal/container"
	"github.com/squaaat/nearsfeed/api/internal/er"
	"github.com/squaaat/nearsfeed/api/internal/model"
	manufactureStore "github.com/squaaat/nearsfeed/api/internal/service/manufacture/store"
)

type Service struct {
	C                *container.Container
	ManufactureStore *manufactureStore.Service
}

func New(c *container.Container) *Service {
	return &Service{
		C:                c,
		ManufactureStore: manufactureStore.New(c),
	}
}

type In struct {
	Manufacture *model.Manufacture
}

type Out struct {
	Manufactures []*model.Manufacture
	Manufacture  *model.Manufacture
}

func (s *Service) PutManufacture(m *model.Manufacture) (*Out, error) {
	op := er.CallerOp()

	man, err := s.ManufactureStore.InsertManufactureOnlyNotExist(m)
	if err != nil {
		return nil, er.WrapOp(err, op)
	}
	return &Out{
		Manufacture: man,
	}, nil
}

func (s *Service) GetManufactures() (*Out, error) {
	op := er.CallerOp()

	var manufactures []*model.Manufacture

	manufactures, err := s.ManufactureStore.SelectManufactures()
	if err != nil {
		return nil, er.WrapOp(err, op)
	}

	return &Out{
		Manufactures: manufactures,
	}, nil
}

func (s *Service) GetManufacture(id string) (*Out, error) {
	op := er.CallerOp()

	man, err := s.ManufactureStore.SelectOneManufacture(id)
	if err != nil {
		return nil, er.WrapOp(err, op)
	}

	return &Out{
		Manufacture: man,
	}, nil
}
