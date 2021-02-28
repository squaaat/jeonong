package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMustLoadDataAtLocal(t *testing.T) {
	data, err := MustLoadDataAtLocal()
	assert.Empty(t, err)
	assert.NotEmpty(t, data.Manufactures[0])
}
