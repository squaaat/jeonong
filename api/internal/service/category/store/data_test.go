package store

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"github.com/squaaat/nearsfeed/api/internal/app"
	"github.com/squaaat/nearsfeed/api/internal/config"
)

func TestMustLoadDataAtLocal(t *testing.T) {
	data, err := MustLoadDataAtLocal()
	assert.Empty(t, err)
	assert.NotEmpty(t, data.Depth1)
	assert.NotEmpty(t, data.Depth2)
	assert.NotEmpty(t, data.Depth3)
}

func TestCategoryFullNameIds(t *testing.T) {
	c := config.MustInit(os.Getenv("J_ENV"), false)
	a := app.New(c)

	type args struct {
		tx          *gorm.DB
		categoryIDs []string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "category의 depth가 1, 2, 3, 4 다 있는 경우",
			args: args{
				tx: a.ServiceDB.DB,
				categoryIDs: []string{"1", "2", "3", "4"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CategoryFullNameIds(tt.args.tx, tt.args.categoryIDs...)
			if (err != nil) != tt.wantErr {
				t.Errorf("CategoryFullNameIds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CategoryFullNameIds() got = %v, want %v", got, tt.want)
			}
		})
	}
}