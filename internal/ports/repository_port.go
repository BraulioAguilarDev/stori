package ports

import "stori/internal/core/domain"

type TransactionRepositoryPort interface {
	Create(dto []domain.TransactionDTO) error
	TransactionsByMonth(uuid string) ([]*domain.TransactionDTO, error)
}

type ProfileRepositoryPort interface {
	Create(dto *domain.ProfileDTO) (*domain.ProfileDTO, error)
	GetByID(uuid string) (*domain.ProfileDTO, error)
}

type AccountRepositoryPort interface {
	Create(dto *domain.AccountDTO) (*domain.AccountDTO, error)
	GetByID(uuid string) (*domain.AccountDTO, error)
}

type AccountS3RepositoryPort interface {
	Create(dto *domain.AccountS3DTO) (*domain.AccountS3DTO, error)
	GetFileByAccountID(uuid string) (string, error)
	Find(uuid string) ([]*domain.AccountS3DTO, error)
}
