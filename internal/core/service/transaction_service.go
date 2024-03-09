package service

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"stori/internal/core/domain"
	"stori/internal/ports"
	"stori/pkg/mail"
	"strconv"
	"text/template"
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

type Tx struct {
	Month string
	Count int
}
type Summary struct {
	Total                       float64
	Txs                         []Tx
	AverageDebit, AverageCredit float64
}

func (srv *transactionService) Create(accountID string, rows [][]string) error {
	type Item struct {
		Date   sqltime.Time
		Debit  float64
		Credit float64
	}

	var txns []Item
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

		txns = append(txns, Item{
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

	if err := srv.repo.Create(data); err != nil {
		return err
	}

	txnsMonth, err := srv.repo.TransactionsByMonth(accountID)
	if err != nil {
		panic(err)
	}

	var debitAmount, creditAmount float64
	var debitCount, creditCount int

	var txsmonth = make([]Tx, 0)
	for _, month := range txnsMonth {
		creditAmount += month.CreditAmount
		debitAmount += month.DebitAmount
		creditCount += month.CreditCount
		debitCount += month.DebitCount

		txsmonth = append(txsmonth, Tx{
			Month: month.Month,
			Count: month.CreditCount + month.DebitCount,
		})
	}

	summary := &Summary{
		Total:         roundFloat(creditAmount+debitAmount, 2),
		AverageCredit: roundFloat(creditAmount/float64(creditCount), 2),
		AverageDebit:  roundFloat(debitAmount/float64(debitCount), 2),
		Txs:           txsmonth,
	}

	tmp := template.Must(template.ParseFiles("/usr/share/summary.html"))
	var body bytes.Buffer
	if err := tmp.Execute(&body, summary); err != nil {
		return err
	}

	// TODO: SEND to AWS SES and SQS
	msg := mail.NewEmail(os.Getenv("KEYSG"))
	if err := msg.Send("email@domain", os.Getenv("SENDER"), "Summary", body.String()); err != nil {
		fmt.Printf("Email sending error: %s \n", err.Error())
	}

	return nil
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
