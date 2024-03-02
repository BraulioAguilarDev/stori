package repository

import "gorm.io/gorm"

type transactionRepository struct {
	db *gorm.DB
}

func NewEventsToSentRepository(db *gorm.DB) *transactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (repo *transactionRepository) Create() error {
	return nil
}
