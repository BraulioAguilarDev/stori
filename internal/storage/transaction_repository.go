package repository

import (
	"stori/internal/core/domain"

	"gorm.io/gorm"
)

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *transactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (repo *transactionRepository) Create(dto []domain.TransactionDTO) error {
	var transactions []*domain.Transaction

	for _, tx := range dto {
		transactions = append(transactions, &domain.Transaction{
			AccountID:    tx.AccountID,
			Date:         tx.Date,
			DebitAmount:  tx.DebitAmount,
			CreditAmount: tx.CreditAmount,
		})
	}

	if err := repo.db.Create(transactions).Error; err != nil {
		return err
	}

	return nil
}
