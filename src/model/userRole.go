package model

import (
	"time"

	"github.com/google/uuid"
)

type UserRole struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
	User      User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Role      string    `gorm:"type:varchar(50);not null;check:role IN ('farmer','expert','admin','support')"`
	CreatedAt time.Time
}
