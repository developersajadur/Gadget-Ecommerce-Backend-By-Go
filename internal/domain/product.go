package domain

import "time"

type Product struct {
	ID            string     `json:"id" db:"id"`
	Name          string     `json:"name" db:"name"`
	Slug          string     `json:"slug" db:"slug"`
	Description   string     `json:"description,omitempty" db:"description"`
	Price         uint32     `json:"price" db:"price"`
	DiscountPrice *uint32    `json:"discount_price,omitempty" db:"discount_price"`
	Stock         int        `json:"stock" db:"stock"`
	CategoryID    *string    `json:"category_id,omitempty" db:"category_id"`
	Images        []string   `json:"images,omitempty" db:"images"`
	IsDeleted     bool       `json:"is_deleted" db:"is_deleted"`
	IsPublished   bool       `json:"is_published" db:"is_published"`
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at" db:"updated_at"`
}
