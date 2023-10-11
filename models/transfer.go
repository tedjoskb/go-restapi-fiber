package models

import "time"

type Transfer struct {
	Transfer_id   int64     `gorm:"type:uuid;primaryKey" json:"transfer_id"`
	FromAccountID int64     `json:"from_account_id"`
	ToAccountID   int64     `json:"to_account_id"`
	Amount        float64   `gorm:"type:DECIMAL(12,2)" json:"amount"`
	Created_at    time.Time `json:"created_at"`
}
