package domain

import "time"

type Category struct {
	ID          string    `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Slug        string    `json:"slug" db:"slug"`
	Description string    `json:"description,omitempty" db:"description"`
	Image       *string   `json:"image,omitempty" db:"image"`
	IsDeleted   bool      `json:"is_deleted" db:"is_deleted"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
