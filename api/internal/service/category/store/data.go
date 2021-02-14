package store

import (
	"fmt"
	"io/ioutil"

	"github.com/ghodss/yaml"

	_const "github.com/squaaat/jeonong/api/internal/const"
	"github.com/squaaat/jeonong/api/internal/er"
)

type DataCategories struct {
	Depth1 []*Category `yml:"depth1"`
	Depth2 []*Category `yml:"depth2"`
	Depth3 []*Category `yml:"depth3"`
}

type Category struct {
	ParentName string `yml:"parentName"`
	Sort       int64  `yml:"sort"`
	Name       string `yml:"name"`
	Code       string `yml:"code"`
}

func MustLoadDataAtLocal() (*DataCategories, error) {
	var op = er.CallerOp()

	data := new(DataCategories)
	yamlFile, err := ioutil.ReadFile(fmt.Sprintf("%s/data/categories.yml", _const.ProjectRootPath))
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
