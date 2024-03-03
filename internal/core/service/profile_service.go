package service

import (
	"stori/internal/core/domain"
	"stori/internal/ports"
)

type profileService struct {
	repo ports.ProfileServicePort
}

func ProvideProfileService(repo ports.ProfileServicePort) *profileService {
	return &profileService{
		repo: repo,
	}
}

func (srv *profileService) Create(dto *domain.ProfileDTO) (*domain.ProfileDTO, error) {
	return srv.repo.Create(dto)
}

func (srv *profileService) GetByID(uuid string) (*domain.ProfileDTO, error) {
	return srv.repo.GetByID(uuid)
}
