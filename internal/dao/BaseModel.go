package dao

import (
	"time"
)

type BaseModel struct {
	ID        int       `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
