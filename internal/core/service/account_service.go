package service

import (
	"stori/internal/core/domain"
	"stori/internal/ports"
)

type accountService struct {
	repo ports.AccountServicePort
}

func ProvideAccountService(repo ports.AccountServicePort) *accountService {
	return &accountService{
		repo: repo,
	}
}

func (srv *accountService) Create(dto *domain.AccountDTO) (*domain.AccountDTO, error) {
	return srv.repo.Create(dto)
}
