package domain

import "time"

type Account struct {
	ID        string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Owner     string    `gorm:"not null"`
	Bank      string    `gorm:"not null"`
	Type      string    `gorm:"not null"`
	Number    string    `gorm:"not null"`
	ProfileID string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (a *Account) TableName() string {
	return "account"
}

type AccountDTO struct {
	ID        string    `json:"account_id"`
	Owner     string    `json:"owner"`
	Bank      string    `json:"bank"`
	Type      string    `json:"type"`
	Number    string    `json:"number"`
	ProfileID string    `json:"profile_id"`
	CreatedAt time.Time `json:"created_at"`
}
