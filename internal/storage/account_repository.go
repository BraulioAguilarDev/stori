package repository

import (
	"stori/internal/core/domain"

	"gorm.io/gorm"
)

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *accountRepository {
	return &accountRepository{
		db: db,
	}
}

func (repo *accountRepository) Create(dto *domain.AccountDTO) (*domain.AccountDTO, error) {
	data := &domain.Account{
		Owner:     dto.Owner,
		Bank:      dto.Bank,
		Type:      dto.Type,
		Number:    dto.Number,
		ProfileID: dto.ProfileID,
	}

	if err := repo.db.Save(data).Error; err != nil {
		return nil, err
	}

	return repo.fromEntity(*data), nil
}

func (repo *accountRepository) GetByID(uuid string) (*domain.AccountDTO, error) {
	var account = domain.Account{ID: uuid}
	if err := repo.db.First(&account).Error; err != nil {
		return nil, err
	}

	return repo.fromEntity(account), nil
}

func (repo *accountRepository) GetEmail(uuid string) string {
	var email string
	query := `select p.email as email from profile p inner join account a on a.profile_id = p.id where a.id = ?`
	if err := repo.db.Raw(query, uuid).Scan(&email).Error; err != nil {
		return ""
	}

	return email
}

// Mapping domain struct to DTO
func (repo *accountRepository) fromEntity(account domain.Account) *domain.AccountDTO {
	return &domain.AccountDTO{
		ID:        account.ID,
		Owner:     account.Owner,
		Bank:      account.Bank,
		Type:      account.Type,
		Number:    account.Number,
		ProfileID: account.ProfileID,
		CreatedAt: account.CreatedAt,
	}
}
