package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	_const "github.com/squaaat/nearsfeed/api/internal/const"
)

func TestMustInit(t *testing.T) {
	env := os.Getenv(_const.KeyEnv)
	cfg, err := New(env)
	assert.Empty(t, err)
	assert.NotEmpty(t, cfg)
}
