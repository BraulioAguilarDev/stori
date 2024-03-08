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

/*
month|credit_amount|debit_amount|debit_count|credit_count|
-----+-------------+------------+-----------+------------+
aug  |        10.00|      -20.46|          1|           1|
jul  |        60.50|      -10.30|          1|           1|
*/
func (r *transactionRepository) TransactionsByMonth(uuid string) ([]*domain.TransactionDTO, error) {
	var tnx []*domain.TransactionDTO
	averageQuery := `
	SELECT to_char(date,'mon') AS month,
		SUM(t.credit_amount) AS credit_amount, 
		SUM(t.debit_amount) AS debit_amount,
		COUNT(case when t.debit_amount != 0 then t.debit_amount end) AS debit_count,
		COUNT(case when t.credit_amount != 0 then t.credit_amount end) AS credit_count
	FROM transaction t where t.account_id = ? GROUP BY 1;`

	if err := r.db.Raw(averageQuery, uuid).Scan(&tnx).Error; err != nil {
		return nil, err
	}

	return tnx, nil
}
