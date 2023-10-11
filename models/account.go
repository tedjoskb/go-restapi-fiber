package models

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	ID           int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	Owner        string         `gorm:"type:varchar(300)" json:"owner"`
	Balance      float64        `gorm:"type:DECIMAL(12,2)" json:"balance"`
	Currency     string         `gorm:"type:varchar(100)" json:"currency"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	UserID       int64          `gorm:"index" json:"user_id"`     // Menyimpan kunci asing
	Entries      []Entries      `gorm:"foreignKey:AccountId"`     // Definisikan relasi One-to-Many dengan Entries.
	TransferFrom []Transfer     `gorm:"foreignKey:FromAccountID"` // Relasi ke Transfer yang memiliki FromAccountID
	TransferTo   []Transfer     `gorm:"foreignKey:ToAccountID"`   // Relasi ke Transfer yang memiliki ToAccountID
}
