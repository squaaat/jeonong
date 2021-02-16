package config

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	_const "github.com/squaaat/nearsfeed/api/internal/const"
)

func TestMustInit(t *testing.T) {
	env := os.Getenv(_const.KeyEnv)
	cicd, _ := strconv.ParseBool(os.Getenv(_const.KeyCicd))
	cfg := MustInit(env, cicd)
	assert.NotEmpty(t, cfg)
}
