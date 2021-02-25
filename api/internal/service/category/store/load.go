package store

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"

	_const "github.com/squaaat/nearsfeed/api/internal/const"
	"github.com/squaaat/nearsfeed/api/internal/er"
	"github.com/squaaat/nearsfeed/api/internal/model"
)

type DataCategories struct {
	Depth1 []*Category `yaml:"depth1"`
	Depth2 []*Category `yaml:"depth2"`
	Depth3 []*Category `yaml:"depth3"`
}

type Category struct {
	Sort       int64  `yaml:"sort"`
	Name       string `yaml:"name"`
	Code       string `yaml:"code"`
	ParentCode string `yaml:"parentCode"`
}

func loadCategoryData() (*DataCategories, error) {
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

func (s *Service) MustLoadDataAtLocal() error {
	op := er.CallerOp()

	data, err := loadCategoryData()
	if err != nil {
		err = er.WrapOp(err, op)
		return err
	}

	for _, cat := range data.Depth1 {
		mc := &model.Category{
			Name:   cat.Name,
			Code:   cat.Code,
			Status: model.StatusIdle,
			Sort:   cat.Sort,
			Depth:  1,
		}
		_, err = s.InsertCategoryIfNotExist(mc)
		if err != nil {
			return er.WrapOp(err, op)
		}
	}
	for _, cat := range data.Depth2 {
		c1, _, _, _, err := s.categoryIDsByCode(cat.ParentCode)
		if err != nil {
			return er.WrapOp(err, op)
		}

		mc := &model.Category{
			Name:        cat.Name,
			Code:        cat.Code,
			Status:      model.StatusIdle,
			Sort:        cat.Sort,
			Depth:       2,
			Category1ID: c1,
		}
		_, err = s.InsertCategoryIfNotExist(mc)
		if err != nil {
			return er.WrapOp(err, op)
		}
	}
	for _, cat := range data.Depth3 {
		c1, c2, _, _, err := s.categoryIDsByCode(cat.ParentCode)
		if err != nil {
			return er.WrapOp(err, op)
		}

		mc := &model.Category{
			Name:        cat.Name,
			Code:        cat.Code,
			Status:      model.StatusIdle,
			Sort:        cat.Sort,
			Depth:       3,
			Category1ID: c1,
			Category2ID: c2,
		}
		_, err = s.InsertCategoryIfNotExist(mc)
		if err != nil {
			return er.WrapOp(err, op)
		}
	}

	return nil
}
