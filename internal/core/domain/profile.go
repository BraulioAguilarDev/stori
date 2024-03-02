package domain

import (
	"github.com/SamuelTissot/sqltime"
)

type Profile struct {
	ID        string       `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string       `gorm:"not null"`
	Email     string       `gorm:"not null"`
	Firebase  string       `gorm:"not null"`
	CreatedAt sqltime.Time `gorm:"autoCreateTime:milli"`
}

func (u *Profile) TableName() string {
	return "profile"
}

type ProfileDTO struct {
	ID        string       `json:"profile_id"`
	Name      string       `json:"name"`
	Email     string       `json:"email"`
	Firebase  string       `json:"firebase_id"`
	CreatedAt sqltime.Time `json:"created_at"`
}
