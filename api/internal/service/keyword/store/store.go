package store

import (
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/squaaat/nearsfeed/api/internal/model"
)

func GetKeywordByCode(tx *gorm.DB, code string) (*model.Keyword, error) {
	var r model.Keyword
	subTx := tx.Take(&r, "code = ?", code).Scan(&r)
	if subTx.Error != nil {
		if !errors.Is(subTx.Error, gorm.ErrRecordNotFound) {
			return nil, subTx.Error
		}
	}
	if subTx.RowsAffected != 1 {
		return nil, errors.New(fmt.Sprintf("Not found record about code[%s]", code))
	}
	return &r, nil
}

func MustGetKeyword(tx *gorm.DB, name, code string) (*model.Keyword, error) {
	var r model.Keyword
	subTx := tx.Take(&r, "name = ?", name).Scan(&r)
	if subTx.Error != nil {
		if !errors.Is(subTx.Error, gorm.ErrRecordNotFound) {
			return nil, subTx.Error
		}
	}
	if subTx.RowsAffected == 1 {
		return &r, nil
	}

	r = model.Keyword{
		Name: name,
		Code: code,
		DefaultModel: model.DefaultModel{
			Status: model.StatusIdle,
		},
	}

	subTx = tx.Create(&r)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &r, nil
}
