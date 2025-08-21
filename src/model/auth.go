package model

import (
	"time"

	"github.com/google/uuid"
)

type AuthToken struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID       uuid.UUID `gorm:"type:uuid;not null"`
	User         User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	RefreshToken string    `gorm:"type:text;not null"`
	UserAgent    string    `gorm:"type:text"`
	IPAddress    string    `gorm:"type:varchar(45)"`
	ExpiresAt    time.Time `gorm:"not null"`
	Revoked      bool      `gorm:"default:false"`
	CreatedAt    time.Time
}
