package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID        int64          `gorm:"primaryKey;autoIncrement" json:"user_id"`
	Name      string         `gorm:"type:varchar(300)" json:"name"`
	Email     string         `gorm:"type:varchar(100)" json:"email"`
	Address   string         `gorm:"type:text" json:"address"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Account   Account        `gorm:"foreignKey:UserID"` // Menunjukkan kunci asing
}