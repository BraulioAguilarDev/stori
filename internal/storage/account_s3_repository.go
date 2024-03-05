package repository

import (
	"stori/internal/core/domain"

	"gorm.io/gorm"
)

type accountS3Repository struct {
	db *gorm.DB
}

func NewAccountS3Repository(db *gorm.DB) *accountS3Repository {
	return &accountS3Repository{
		db: db,
	}
}

func (repo *accountS3Repository) Create(dto *domain.AccountS3DTO) (*domain.AccountS3DTO, error) {
	data := &domain.AccountS3{
		AccountID: dto.AccountID,
		URL:       dto.URL,
		Filename:  dto.Filename,
	}

	if err := repo.db.Create(data).Error; err != nil {
		return nil, err
	}

	return &domain.AccountS3DTO{
		AccountID: data.AccountID,
		URL:       data.URL,
		Filename:  data.Filename,
	}, nil
}

func (repo *accountS3Repository) GetFileByAccountID(uuid string) (string, error) {
	var accountS3 = domain.AccountS3{AccountID: uuid}
	if err := repo.db.First(&accountS3).Error; err != nil {
		return "", err
	}

	return accountS3.Filename, nil
}
