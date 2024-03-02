package domain

import (
	"github.com/SamuelTissot/sqltime"
)

type Transaction struct {
	ID            string       `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Date          sqltime.Time `gorm:"not null"`
	DebitAmount   float64
	CreditAmount  float64
	BankAccountID string       `gorm:"not null"`
	CreatedAt     sqltime.Time `gorm:"autoCreateTime:milli"`
}

func (t *Transaction) TableName() string {
	return "transaction"
}

type TransactionDTO struct {
	TransactionID string       `json:"transaction_id"`
	AccountID     string       `json:"account_id"`
	Date          sqltime.Time `json:"date"`
	DebitAmount   float64      `json:"debit_amount"`
	CreditAmount  float64      `json:"credit_amount"`
	BankAccountID string       `json:"bank_id"`
	CreatedAt     sqltime.Time `json:"created_at"`
}
