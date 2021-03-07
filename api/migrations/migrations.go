package migrations

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"

	"github.com/squaaat/nearsfeed/api/internal/container"
	"github.com/squaaat/nearsfeed/api/internal/model"
)

type Syncker struct {
	App          *container.Container
	GormMigrator *gormigrate.Gormigrate
}

func New(a *container.Container) *Syncker {
	s := &Syncker{
		App: a,
	}
	s.GormMigrator = gormigrate.New(
		a.ServiceDB.DB,
		gormigrate.DefaultOptions,
		s.load(),
	)
	s.GormMigrator.InitSchema(func(m *gorm.DB) error {
		return m.AutoMigrate(model.Load()...)
	})
	return s
}

func (s *Syncker) Sync() error {
	return s.GormMigrator.Migrate()
}

func (s *Syncker) Create(v string) error {
	tmpl := versionedTemplate(v)

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	dest := fmt.Sprintf("%s/migrations/migration_%s.go", pwd, v)
	err = ioutil.WriteFile(dest, []byte(tmpl), 0644)
	if err != nil {
		return err
	}

	msg := `
Completed to create migrations file

* Migration File: 
Link => file://%s/migrations/migration_%s.go

* Add Migration method
Link => file://%s/migrations/init.go

exmaple: 
func (s *Syncker) load() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		s.migration_202012251400(),
		// Write your codes here
		... ,
	}
}

`
	fmt.Printf(
		msg,
		pwd,
		v,
		pwd,
	)
	return nil
}

func (s *Syncker) load() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		s.migration_202012251400(),
		// migration script
	}
}
