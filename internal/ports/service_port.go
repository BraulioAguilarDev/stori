package ports

import "stori/internal/core/domain"

type TransactionServicePort interface {
	Create() error
}

type ProfileServicePort interface {
	Create(dto *domain.ProfileDTO) (*domain.ProfileDTO, error)
	GetByID(uuid string) (*domain.ProfileDTO, error)
}

type AccountServicePort interface {
	Create(dto *domain.AccountDTO) (*domain.AccountDTO, error)
}
