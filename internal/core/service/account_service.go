package service

import (
	"stori/internal/core/domain"
	"stori/internal/ports"
)

type accountService struct {
	repo ports.AccountRepositoryPort
}

func ProvideAccountService(repo ports.AccountRepositoryPort) *accountService {
	return &accountService{
		repo: repo,
	}
}

func (srv *accountService) Create(dto *domain.AccountDTO) (*domain.AccountDTO, error) {
	return srv.repo.Create(dto)
}

func (srv *accountService) GetByID(uuid string) (*domain.AccountDTO, error) {
	return srv.repo.GetByID(uuid)
}
