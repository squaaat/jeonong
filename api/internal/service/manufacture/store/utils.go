package store

import (
	"gorm.io/gorm"

	"github.com/squaaat/nearsfeed/api/internal/er"
	"github.com/squaaat/nearsfeed/api/internal/model"
)

func GetManufactureByModel(tx *gorm.DB, mc *model.Manufacture) (*model.Manufacture, error) {
	op := er.CallerOp()

	subTx := tx.
		Model(mc).
		First(mc, "name = ? AND code = ?", mc.Name, mc.Code).
		Scan(mc)
	if subTx.Error != nil {
		return nil, er.WrapOp(subTx.Error, op)
	}

	return mc, nil
}
