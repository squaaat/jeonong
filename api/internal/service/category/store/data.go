package store

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"

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

func AddCategoryIfNotExist(tx *gorm.DB, cat *Category, depth int64, categoriesIDs... string) (*model.Category, error) {
	op := er.CallerOp()

	c := &model.Category{}
	subTx := tx.
		Model(c).
		Where("depth = ? AND code = ? AND name = ?", depth, cat.Code, cat.Name).
		Scan(c)

	if subTx.Error != nil {
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, er.WrapOp(subTx.Error, op)
		}
	}

	if c.ID != "" {
		return c, nil
	}


	ids := make([]string, 4)
	for i, id := range categoriesIDs {
		ids[i] = id
	}
	if ids[depth-1] == "" {
		ids[depth-1] = uuid.NewString()
	}

	c = &model.Category{
		DefaultModel: model.DefaultModel{
			ID: ids[depth-1],
		},
		Status: model.StatusIdle,
		Name: cat.Name,
		Code: cat.Code,
		Sort: cat.Sort,
		Depth: depth,
		Category1ID: ids[0],
		Category2ID: ids[1],
		Category3ID: ids[2],
		Category4ID: ids[3],
	}
	subTx = tx.Create(c)
	if subTx.Error != nil {
		return nil, subTx.Error
	}
	if subTx.RowsAffected != 1 {
		return nil, fmt.Errorf("failed create categroy [input: %v]", c)
	}

	names, err := CategoryFullNameIds(tx, ids...)
	if err != nil {
		return nil, err
	}

	fullName := []string{}
	for _, name := range names {
		if name != "" {
			fullName = append(fullName, name)
		}
	}

	subTx = tx.Model(c).Update("full_name", strings.Join(fullName, ",")).Where("id = ?", c.ID)
	if subTx.Error != nil {
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, er.WrapOp(subTx.Error, op)
		}
		return nil, subTx.Error
	}
	if subTx.RowsAffected != 1 {
		return nil, fmt.Errorf("failed create categroy [input: %v]", c)
	}

	return c, nil
}


type CategoryNames struct {
	Category1Name string `sql:"category1_name"`
	Category2Name string `sql:"category2_name"`
	Category3Name string `sql:"category3_name"`
	Category4Name string `sql:"category4_name"`
}

func CategoryIdsByKeywordCode(tx *gorm.DB, code string) ([]string, error) {
	op := er.CallerOp()

	c := &model.Category{}
	subTx := tx.
		Model(c).
		Where("code = ?", code).
		Scan(c)

	if subTx.Error != nil {
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, er.WrapOp(subTx.Error, op)
		}
	}

	var r []string
	if c.Category1ID != "" {
		r = append(r, c.Category1ID)
	}
	if c.Category2ID != "" {
		r = append(r, c.Category2ID)
	}
	if c.Category3ID != "" {
		r = append(r, c.Category3ID)
	}
	if c.Category4ID != "" {
		r = append(r, c.Category4ID)
	}

	return r, nil
}

func CategoryFullNameIds(tx *gorm.DB, categoryIDs... string) ([]string, error){
	op := er.CallerOp()

	ids := make([]string, 4)
	for i, id := range categoryIDs {
		ids[i] = id
	}

	categoryTableName := (&model.Category{}).TableName()

	names := &CategoryNames{}
	subTx := tx.Raw(fmt.Sprintf(`
SELECT
	(CASE
		WHEN ? != ''
		THEN (SELECT name FROM %s as c WHERE c.category1_id = ? LIMIT 1)
		ELSE ''
	END) as category1_name
	,(CASE
		WHEN ? != ''
		THEN (SELECT name FROM %s as c WHERE c.category2_id = ? LIMIT 1)
		ELSE ''
	END) as category2_name
	,(CASE
		WHEN ? != ''
		THEN (SELECT name FROM %s as c WHERE c.category3_id = ? LIMIT 1)
		ELSE ''
	END) as category3_name
	,(CASE
		WHEN ? != ''
		THEN (SELECT name FROM %s as c WHERE c.category4_id = ? LIMIT 1)
		ELSE ''
	END) as category4_name
`,
		categoryTableName,
		categoryTableName,
		categoryTableName,
		categoryTableName,
	), ids[0], ids[0], ids[1], ids[1], ids[2], ids[2], ids[3], ids[3]).Scan(names)

	if subTx.Error != nil {
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, er.WrapOp(tx.Error, op)
		}
	}

	return []string{names.Category1Name, names.Category2Name, names.Category3Name, names.Category4Name}, nil
}