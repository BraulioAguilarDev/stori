package ports

import "stori/internal/core/domain"

type TransactionServicePort interface {
	Create(accountID, email string, rows [][]string) error
}

type ProfileServicePort interface {
	Create(dto *domain.ProfileDTO) (*domain.ProfileDTO, error)
	GetByID(uuid string) (*domain.ProfileDTO, error)
}

type AccountServicePort interface {
	Create(dto *domain.AccountDTO) (*domain.AccountDTO, error)
	GetByID(uuid string) (*domain.AccountDTO, error)
	GetEmail(accountId string) string
}

type AccountS3ServicePort interface {
	Create(dto *domain.AccountS3DTO) (*domain.AccountS3DTO, error)
	GetFileByAccountID(uuid string) (string, error)
	Find(accountId string) ([]*domain.AccountS3DTO, error)
}
