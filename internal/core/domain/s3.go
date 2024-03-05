package domain

type AccountS3 struct {
	AccountID string `gorm:"not null"`
	URL       string `gorm:"not null"`
	Filename  string
}

func (s *AccountS3) TableName() string {
	return "account_s3"
}

type AccountS3DTO struct {
	AccountID string `json:"account_id"`
	URL       string `json:"url"`
	Filename  string `json:"filename"`
}
