package ports

import "stori/internal/core/domain"

type TransactionRepositoryPort interface {
	Create() error
}

type RegisterRepositoryPort interface {
	Create(dto *domain.ProfileDTO) (*domain.ProfileDTO, error)
}
