package domain

import "time"

type Account struct {
	ID        string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string    `gorm:"not null"`
	Type      string    `gorm:"not null"`
	Number    int       `gorm:"not null"`
	UserID    string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime:milli"`
}

func (a *Account) TableName() string {
	return "account"
}

type AccountDTO struct {
	ID        string    `json:"account_id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Number    int       `json:"number"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
