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

func AddCategoryIfNotExist(tx *gorm.DB, keyword, parentKeyword *model.Keyword) (*model.Category, error) {
	if keyword == nil {
		return nil, errors.New("'keyword' is primary")
	}

	parentKeywordID := keyword.ID
	parentKeywordName := keyword.Name
	if parentKeyword != nil {
		parentKeywordID = parentKeyword.ID
		parentKeywordName = parentKeyword.Name
	}

	c := &model.Category{
		KeywordID:       keyword.ID,
		ParentKeywordID: parentKeywordID,
		DefaultModel: model.DefaultModel{
			Status: model.StatusIdle,
		},
	}

	subTx := tx.
		Preload("Keyword").
		Preload("ParentKeyword").
		Take(c, "keyword_id = ? AND parent_keyword_id = ?", c.KeywordID, c.ParentKeywordID).
		Scan(c)
	if subTx.Error != nil {
		if !errors.Is(subTx.Error, gorm.ErrRecordNotFound) {
			return nil, subTx.Error
		}
	}
	if subTx.RowsAffected == 1 {
		return c, nil
	}

	subTx = tx.Create(c)
	if subTx.Error != nil {
		return nil, subTx.Error
	}
	if subTx.RowsAffected != 1 {
		return nil, fmt.Errorf("failed create categroy [child: %s/parent: %s]", keyword.Name, parentKeywordName)
	}

	return c, nil
}
