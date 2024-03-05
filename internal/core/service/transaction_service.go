package service

import (
	"stori/internal/core/domain"
	"stori/internal/ports"
	"strconv"
	"time"

	"github.com/SamuelTissot/sqltime"
)

type transactionService struct {
	repo ports.TransactionRepositoryPort
}

func ProvideTransactionService(repo ports.TransactionRepositoryPort) *transactionService {
	return &transactionService{
		repo: repo,
	}
}

func (srv *transactionService) Create(accountID string, rows [][]string) error {
	type Tx struct {
		Date   sqltime.Time
		Debit  float64
		Credit float64
	}

	var txns []Tx
	for _, item := range rows[1:] {
		date, err := time.Parse("2006-01-02", item[1])
		if err != nil {
			return err
		}

		var debit, credit float64
		mode := item[2][0:1]
		if mode == "-" {
			debit, _ = strconv.ParseFloat(item[2][1:], 64)
		} else {
			credit, _ = strconv.ParseFloat(item[2][1:], 64)
		}

		txns = append(txns, Tx{
			Date: sqltime.Time{
				Time: date,
			},
			Debit:  -debit,
			Credit: credit,
		})
	}

	var data []domain.TransactionDTO
	for _, t := range txns {
		data = append(data, domain.TransactionDTO{
			AccountID:    accountID,
			Date:         t.Date,
			DebitAmount:  t.Debit,
			CreditAmount: t.Credit,
		})
	}

	return srv.repo.Create(data)
}
