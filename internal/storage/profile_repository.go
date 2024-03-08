package repository

import (
	"stori/internal/core/domain"

	"gorm.io/gorm"
)

type profileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) *profileRepository {
	return &profileRepository{
		db: db,
	}
}

func (repo *profileRepository) Create(dto *domain.ProfileDTO) (*domain.ProfileDTO, error) {
	data := &domain.Profile{
		Name:  dto.Name,
		Email: dto.Email,
	}

	if err := repo.db.Save(data).Error; err != nil {
		return nil, err
	}

	return repo.fromEntity(*data), nil
}

func (repo *profileRepository) GetByID(uuid string) (*domain.ProfileDTO, error) {
	var profile = domain.Profile{ID: uuid}
	if err := repo.db.First(&profile).Error; err != nil {
		return nil, err
	}

	return repo.fromEntity(profile), nil
}

// Mapping domain struct to DTO
func (repo *profileRepository) fromEntity(profile domain.Profile) *domain.ProfileDTO {
	return &domain.ProfileDTO{
		ID:        profile.ID,
		Name:      profile.Name,
		Email:     profile.Email,
		CreatedAt: profile.CreatedAt,
	}
}
