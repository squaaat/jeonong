package container

import (
	"github.com/squaaat/nearsfeed/api/internal/config"
	"github.com/squaaat/nearsfeed/api/internal/db"
	"github.com/squaaat/nearsfeed/api/internal/er"
)

type Container struct {
	Config    *config.Config
	ServiceDB *db.Client
}

func New(cfg *config.Config) (*Container, error) {
	op := er.CallerOp()
	c := &Container{
		Config: cfg,
	}

	client, err := db.New(db.ParseConfig(cfg))
	if err != nil {
		return nil, er.WrapOp(err, op)
	}
	c.ServiceDB = client

	return c, nil
}
