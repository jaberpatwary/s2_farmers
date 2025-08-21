package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	PhoneNumber string    `gorm:"type:varchar(15);uniqueIndex;not null"`
	FullName    string    `gorm:"type:varchar(100)"`
	UserType    string    `gorm:"type:varchar(20);not null;check:user_type IN ('farmer','expert','admin')"`
	IsVerified  bool      `gorm:"default:false"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
