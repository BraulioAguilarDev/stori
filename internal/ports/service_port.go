package ports

import "stori/internal/core/domain"

type TransactionServicePort interface {
	Create() error
}

type ProfileServicePort interface {
	Create(dto *domain.ProfileDTO) (*domain.ProfileDTO, error)
}
