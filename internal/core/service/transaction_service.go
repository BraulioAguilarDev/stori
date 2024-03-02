package service

import "stori/internal/ports"

type transactionService struct {
	repo ports.TransactionRepositoryPort
}

func ProvideTransactionService(repo ports.TransactionRepositoryPort) *transactionService {
	return &transactionService{
		repo: repo,
	}
}

// func (srv *transactionService) Create() error {
// 	return srv.repo.Create()
// }
