package domain

import "time"

type Account struct {
	ID           string    `gorm:"primaryKey"`
	AccountName  string    `gorm:"not null"`
	AccountEmail string    `gorm:"not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime:milli"`
}

func (txn *Account) TableName() string {
	return "account"
}

type AccountDTO struct {
	ID           string    `json:"account_id"`
	AccountName  string    `json:"name"`
	AccountEmail string    `json:"emai"`
	CreatedAt    time.Time `json:"created_at"`
}
