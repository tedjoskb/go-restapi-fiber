package models

import (
	"time"
)

type Entries struct {
	Entries_id string    `gorm:"type:uuid;primaryKey" json:"entries_id"`
	AccountId  int64     `json:"account_id"` // Ini adalah kunci asing yang mengacu ke Account
	Amount     float64   `json:"amount"`
	Created_at time.Time `json:"created_at"`
}
