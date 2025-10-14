package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID            string         `gorm:"primaryKey;type:uuid" json:"id"`
	Name          string         `gorm:"size:150;not null" json:"name"`
	Slug          string         `gorm:"size:150;uniqueIndex;not null" json:"slug"`
	Description   string         `gorm:"type:text" json:"description,omitempty"`
	Price         float64        `gorm:"not null" json:"price"`
	DiscountPrice *float64       `json:"discount_price,omitempty"`
	Stock         int            `gorm:"default:0" json:"stock"`
	CategoryID    *string        `gorm:"type:uuid" json:"category_id,omitempty"`
	IsDeleted     bool           `gorm:"default:false" json:"is_deleted"`
	IsPublished   bool           `gorm:"default:false" json:"is_published"`
	Images        []ProductImage `gorm:"foreignKey:ProductID" json:"images,omitempty"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.NewString()
	return
}

type ProductImage struct {
	ID        string    `gorm:"primaryKey;type:uuid" json:"id"`
	ProductID string    `gorm:"type:uuid;index" json:"product_id"`
	URL       string    `gorm:"not null" json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (pi *ProductImage) BeforeCreate(tx *gorm.DB) (err error) {
	pi.ID = uuid.NewString()
	return
}
