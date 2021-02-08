package model

import (
	"time"

	"gorm.io/gorm"
	"github.com/google/uuid"
)

func Load() []interface{} {
	return []interface{}{
		&Articles{},
		&ArticlesDiscounts{},
		&Sales{},
		&Products{},
		&Categories{},
		&CategoryHierarchy{},
		&Manufactures{},
		&Markets{},
	}
}

type DefaultModel struct {
	ID string `gorm:"primaryKey;type:CHAR(36);not null"`

	CreatedBy string         `gorm:"type:CHAR(36);not null"`
	CreatedAt time.Time      `gorm:"type:timestamp;not null"`
	UpdatedBy string         `gorm:"type:CHAR(36);not null"`
	UpdatedAt time.Time      `gorm:"type:timestamp;not null"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamp;index"`
}

func (m *DefaultModel) BeforeCreate(_ *gorm.DB) error {
	if m.ID == "" {
		m.ID = uuid.NewString()
	}
	m.CreatedBy = m.ID
	m.UpdatedBy = m.ID
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

func (m *DefaultModel) BeforeUpdate(_ *gorm.DB) error {
	m.UpdatedAt = time.Now()
	return nil
}

func (m *DefaultModel) BeforeDelete(_ *gorm.DB) error {
	m.DeletedAt = gorm.DeletedAt{
		Time: time.Now(),
		Valid: true,
	}
	return nil
}