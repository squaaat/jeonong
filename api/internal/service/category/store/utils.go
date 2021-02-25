package store

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/squaaat/nearsfeed/api/internal/er"
	"github.com/squaaat/nearsfeed/api/internal/model"
)

type CategoryNames struct {
	Category1Name string `sql:"category1_name"`
	Category2Name string `sql:"category2_name"`
	Category3Name string `sql:"category3_name"`
	Category4Name string `sql:"category4_name"`
}

func (s *Service) categoryIDsByCode(code string) (string, string, string, string, error) {
	op := er.CallerOp()

	c := &model.Category{}
	tx := s.App.ServiceDB.DB.
		Model(c).
		Where("code = ?", code).
		Scan(c)

	if tx.Error != nil {
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return "", "", "", "", er.WrapOp(tx.Error, op)
		}
	}

	return c.Category1ID, c.Category2ID, c.Category3ID, c.Category4ID, nil
}

func (s *Service) categoryFullNameIds(depth int64, cat1, cat2, cat3, cat4 string) (string, error) {
	op := er.CallerOp()

	if err := checkCategoryValid(depth, cat1, cat2, cat3, cat4); err != nil {
		return "", er.WrapOp(err, op)
	}

	ids := make([]string, 4)
	for i, id := range []string{cat1, cat2, cat3, cat4} {
		ids[i] = id
	}

	categoryTableName := (&model.Category{}).TableName()

	names := &CategoryNames{}
	tx := s.App.ServiceDB.DB.Raw(fmt.Sprintf(`
SELECT
	(CASE
		WHEN ? != ''
		THEN (SELECT name FROM %s as c WHERE depth = 1 AND c.id = ? LIMIT 1)
		ELSE ''
	END) as category1_name
	,(CASE
		WHEN ? != ''
		THEN (SELECT name FROM %s as c WHERE depth = 2 AND c.id = ? LIMIT 1)
		ELSE ''
	END) as category2_name
	,(CASE
		WHEN ? != ''
		THEN (SELECT name FROM %s as c WHERE depth = 3 AND c.id = ? LIMIT 1)
		ELSE ''
	END) as category3_name
	,(CASE
		WHEN ? != ''
		THEN (SELECT name FROM %s as c WHERE depth = 4 AND c.id = ? LIMIT 1)
		ELSE ''
	END) as category4_name
`,
		categoryTableName,
		categoryTableName,
		categoryTableName,
		categoryTableName,
	), ids[0], ids[0], ids[1], ids[1], ids[2], ids[2], ids[3], ids[3]).Scan(names)

	if tx.Error != nil {
		if !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return "", er.WrapOp(tx.Error, op)
		}
	}

	if err := checkCategoryValid(depth, names.Category1Name, names.Category2Name, names.Category3Name, names.Category4Name); err != nil {
		return "", er.WrapOp(err, op)
	}

	fullName := []string{}
	for _, name := range []string{names.Category1Name, names.Category2Name, names.Category3Name, names.Category4Name} {
		if name != "" {
			fullName = append(fullName, name)
		}
	}
	return strings.Join(fullName, ","), nil
}

func GetCategoryByModel(tx *gorm.DB, mc *model.Category) (*model.Category, error) {
	op := er.CallerOp()

	c := &model.Category{}
	subTx := tx.
		Model(c).
		Where(`
depth = ?
AND code = ?
AND name = ?
AND (
	CASE
		WHEN depth = 2 THEN category1_id = ?
		WHEN depth = 3 THEN category1_id = ? AND category2_id = ? 
		WHEN depth = 4 THEN category1_id = ? AND category2_id = ? AND category3_id = ?
		ELSE 1 = 1
	END
)`,
			mc.Depth, mc.Code, mc.Name,
			mc.Category1ID,
			mc.Category1ID, mc.Category2ID,
			mc.Category1ID, mc.Category2ID, mc.Category3ID).
		Scan(c)

	if subTx.Error != nil {
		if errors.Is(subTx.Error, gorm.ErrRecordNotFound) {
			return c, nil
		} else {
			return nil, er.WrapOp(subTx.Error, op)
		}
	}
	return c, nil
}

func checkCategoryValid(depth int64, cat1, cat2, cat3, cat4 string) error {
	if depth == 1 {
		if cat2 != "" || cat3 != "" || cat4 != "" {
			return errors.New("Depth is 1, But, Category2ID, Category3ID or Category4ID is not empty")
		}
	}
	if depth == 2 {
		if cat1 == "" {
			return errors.New("Depth is 2, But, Category1ID is empty")
		}
		if cat3 != "" || cat4 != "" {
			return errors.New("Depth is 2, But, Category3ID, Category4ID is not empty")
		}
	}
	if depth == 3 {
		if cat1 == "" || cat2 == "" {
			return errors.New("Depth is 3, But, Category1ID or Category2ID is empty")
		}
		if cat4 != "" {
			return errors.New("Depth is 3, But, Category4ID is not empty")
		}
	}
	if depth == 4 {
		if cat1 == "" || cat2 == "" || cat3 == "" {
			return errors.New("Depth is 4, But, Category1ID, Category2ID or Category3ID is empty")
		}
	}
	return nil
}
