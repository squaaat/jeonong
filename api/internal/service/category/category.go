package category

import (
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/squaaat/nearsfeed/api/internal/app"
	"github.com/squaaat/nearsfeed/api/internal/er"
	"github.com/squaaat/nearsfeed/api/internal/model"
	categoryStore "github.com/squaaat/nearsfeed/api/internal/service/category/store"
	keywordStore "github.com/squaaat/nearsfeed/api/internal/service/keyword/store"
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
	Categories []string
}

type Out struct {
	Categories []*model.Category
}

func (s *Service) PutCategory(in []string) (*Out, error) {
	op := er.CallerOp()
	var keywords = make([]*model.Keyword, len(in))
	var categories = make([]*model.Category, len(in))

	err := s.App.ServiceDB.DB.Transaction(func(tx *gorm.DB) error {
		for i, keyword := range in {
			k, err := keywordStore.MustGetKeyword(tx, keyword, keyword)
			if err != nil {
				return err
			}
			keywords[i] = k
		}

		for i, _ := range keywords {
			if i == (len(keywords) - 1) {
				break
			}
			if i == 0 {
				cat, err := categoryStore.AddCategoryIfNotExist(tx, keywords[i], keywords[i])
				if err != nil {
					return err
				}
				categories[i] = cat
			}
			cat, err := categoryStore.AddCategoryIfNotExist(tx, keywords[i+1], keywords[i])
			if err != nil {
				return err
			}
			categories[i+1] = cat
		}
		return nil
	})
	if err != nil {
		err := er.WrapOp(err, op)
		return nil, err
	}

	return &Out{
		Categories: categories,
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
