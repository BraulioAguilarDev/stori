package service

import (
	"stori/internal/core/domain"
	"stori/internal/ports"
)

type accountS3Service struct {
	repo ports.AccountS3RepositoryPort
}

func ProvideAccountS3Service(repo ports.AccountS3RepositoryPort) *accountS3Service {
	return &accountS3Service{
		repo: repo,
	}
}

func (srv *accountS3Service) Create(dto *domain.AccountS3DTO) (*domain.AccountS3DTO, error) {
	return srv.repo.Create(dto)
}
