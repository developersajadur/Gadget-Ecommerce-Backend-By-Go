package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Otp struct {
	ID         string         `gorm:"primaryKey;type:uuid" json:"id"`
	UserID     string         `gorm:"type:uuid;not null;index" json:"user_id"`
	Email      string         `gorm:"size:100;not null" json:"email"`
	Code       string         `gorm:"size:10;not null" json:"code"`
	ExpiresAt  time.Time      `gorm:"not null" json:"expires_at"`
	Verified   bool           `gorm:"default:false" json:"verified"`
	VerifiedAt *time.Time     `json:"verified_at,omitempty"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

func (o *Otp) BeforeCreate(tx *gorm.DB) (err error) {
	o.ID = uuid.NewString()
	return
}
