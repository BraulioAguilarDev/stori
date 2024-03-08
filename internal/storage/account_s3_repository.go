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

func (repo *accountS3Repository) Find(uuid string) ([]*domain.AccountS3DTO, error) {
	var txns []*domain.AccountS3DTO
	if err := repo.db.Raw(`
		SELECT
			a.owner as account_name,
			a.id as account_id,
			as2.url as url,
			as2.filename as filename
		FROM
			account_s3 as2
		inner join account a on
		as2.account_id = a.id
		WHERE a.id = ?`, uuid).Scan(&txns).Error; err != nil {
		return nil, err
	}

	return txns, nil
}
