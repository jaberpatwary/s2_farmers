package model

import (
	"time"

	"github.com/google/uuid"
)

type OtpToken struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
	User      User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	OtpCode   string    `gorm:"type:varchar(6);not null"`
	Purpose   string    `gorm:"type:varchar(20);not null;check:purpose IN ('login','register','verify')"`
	IsUsed    bool      `gorm:"default:false"`
	ExpiresAt time.Time `gorm:"not null"`
	CreatedAt time.Time
}
