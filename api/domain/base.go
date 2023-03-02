package domain

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	Id        uint      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) error {
	b.CreatedAt = time.Now()
	b.UpdatedAt = b.CreatedAt
	return nil
}

func (b *Base) BeforeUpdate(tx *gorm.DB) error {
	b.UpdatedAt = time.Now()
	return nil
}
