package domain

type AccountS3 struct {
	AccountID string `gorm:"not null"`

	// S3 url
	URL string `gorm:"not null"`
}

func (a *AccountS3) TableName() string {
	return "account_s3"
}

type AccountS3DTO struct {
	AccountID string `json:"account_id"`
	URL       string `json:"url"`
}
