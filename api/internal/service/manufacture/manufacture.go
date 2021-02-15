package manufacture

import (
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/squaaat/jeonong/api/internal/app"
	"github.com/squaaat/jeonong/api/internal/er"
	"github.com/squaaat/jeonong/api/internal/model"
	keywordStore "github.com/squaaat/jeonong/api/internal/service/keyword/store"
	manufactureStore "github.com/squaaat/jeonong/api/internal/service/manufacture/store"
)

type Service struct {
	App           *app.Application
	CategoryStore *manufactureStore.Service
}

func New(a *app.Application) *Service {
	return &Service{
		App:           a,
		CategoryStore: manufactureStore.New(a),
	}
}

type In struct {
	Manufacture string
}

type Out struct {
	Manufactures []*model.Manufacture
	Manufacture *model.Manufacture
}

func (s *Service) PutManufacture(man string) (*Out, error) {
	op := er.CallerOp()
	var keyword *model.Keyword
	var manufacture *model.Manufacture

	err := s.App.ServiceDB.DB.Transaction(func(tx *gorm.DB) error {
		k, err := keywordStore.MustGetKeyword(tx, man, "")
		if err != nil {
			return err
		}

		m, err := manufactureStore.AddManufactureIfNotExist(tx, k, "")
		if err != nil {
			return err
		}
		keyword = k
		manufacture = m
		manufacture.Keyword = *k
		return nil
	})
	if err != nil {
		err := er.WrapOp(err, op)
		return nil, err
	}

	return &Out{
		Manufacture: manufacture,
	}, nil
}

func (s *Service) GetManufactures() (*Out, error) {
	op := er.CallerOp()

	var manufactures []*model.Manufacture
	table := &model.Manufacture{}
	s.App.ServiceDB.DB.Name()
	tx := s.App.ServiceDB.DB.
		Model(table).
		Joins("Keyword").
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
		m := new(model.Manufacture)
		err = tx.ScanRows(rows, &m)
		if err != nil {
			return nil, er.WrapOp(err, op)
		}
		manufactures = append(manufactures, m)
	}
	return &Out{
		Manufactures: manufactures,
	}, nil
}
