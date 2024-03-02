package domain

import (
	"encoding/json"
	"time"
)

type Transaction struct {
	ID           string    `gorm:"primaryKey"`
	AccountID    string    `gorm:"not null"`
	Date         time.Time `gorm:"not null"`
	DebitAmount  float64
	CreditAmount float64
	Metadata     json.RawMessage `gorm:"not null"`
	CreatedAt    time.Time       `gorm:"autoCreateTime:milli"`
}

func (txn *Transaction) TableName() string {
	return "transaction"
}

type TransactionDTO struct {
	TransactionID uint `json:"transaction_id"`
}
