package store

import (
	"fmt"
	"io/ioutil"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"

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

func (s *Service) MustLoadDataAtLocal() error {
	op := er.CallerOp()

	data, err := MustLoadDataAtLocal()
	if err != nil {
		err = er.WrapOp(err, op)
		return err
	}

	for _, man := range data.Manufactures {
		c, err := s.InsertManufactureIfNotCategory(&model.Manufacture{
			Name:                      man.Name,
			Code:                      man.Code,
			CompanyRegistrationNumber: man.CompanyRegistrationNumber,
		})
		if err != nil {
			return er.WrapOp(err, op)
		}
		log.Debug().Interface("category", c).Send()
	}
	return nil
}
