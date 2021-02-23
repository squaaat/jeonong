package store

import (
	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"

	_const "github.com/squaaat/nearsfeed/api/internal/const"
	"github.com/squaaat/nearsfeed/api/internal/er"
	"github.com/squaaat/nearsfeed/api/internal/model"
)

type DataManufactures struct {
	Manufactures []*Manufacture `yml:"manufactures"`
}

type Manufacture struct {
	Code                      string `yml:"code"`
	Name                      string `yml:"name"`
	CompanyRegistrationNumber string `yml:"companyRegistrationNumber"`
}

func MustLoadDataAtLocal() (*DataManufactures, error) {
	var op = er.CallerOp()

	data := new(DataManufactures)
	yamlFile, err := ioutil.ReadFile(fmt.Sprintf("%s/data/manufactures.yml", _const.ProjectRootPath))
	if err != nil {
		err = er.WrapOp(err, op)
		return nil, err
	}

	if err = yaml.Unmarshal(yamlFile, data); err != nil {
		err = er.WrapOp(err, op)
		return nil, err
	}

	return data, nil
}

func AddManufactureIfNotExist(tx *gorm.DB, m *Manufacture) (*model.Manufacture, error) {
	man := &model.Manufacture{
		Name: m.Name,
		Code: m.Code,
		CompanyRegistrationNumber: m.CompanyRegistrationNumber,
		DefaultModel: model.DefaultModel{
			Status: model.StatusIdle,
		},
	}

	subTx := tx.Take(man, "name = ? AND code = ?", m.Name, m.Code).Scan(man)
	if subTx.Error != nil {
		if !errors.Is(subTx.Error, gorm.ErrRecordNotFound) {
			return nil, subTx.Error
		}
	}
	if subTx.RowsAffected == 1 {
		return man, nil
	}

	subTx = tx.Create(man)
	if subTx.Error != nil {
		return nil, subTx.Error
	}
	if subTx.RowsAffected != 1 {
		return nil, fmt.Errorf("failed create manufacture : [%v]", man)
	}

	return man, nil
}
