package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID         string         `gorm:"primaryKey;type:uuid" json:"id"`
	Name       string         `gorm:"size:100;not null" json:"name"`
	Email      string         `gorm:"size:100;uniqueIndex;not null" json:"email"`
	Password   string         `gorm:"not null" json:"-"`
	Role       string         `gorm:"size:20;default:'user'" json:"role"` // "user" or "admin"
	IsDeleted  bool           `gorm:"default:false" json:"is_deleted"`
	IsBlocked  bool           `gorm:"default:false" json:"is_blocked"`
	IsVerified bool           `gorm:"default:false" json:"is_verified"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	return
}

const (
	RoleUser  = "user"
	RoleAdmin = "admin"
)
