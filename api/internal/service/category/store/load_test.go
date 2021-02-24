package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMustLoadDataAtLocal(t *testing.T) {
	data, err := loadCategoryData()
	assert.Empty(t, err)
	assert.NotEmpty(t, data.Depth1)
	assert.NotEmpty(t, data.Depth2)
	assert.NotEmpty(t, data.Depth3)
}
