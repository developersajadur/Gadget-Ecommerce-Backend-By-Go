package domain

import "time"

type User struct {
	ID        int64  `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"-" db:"password"`
	IsDeleted bool   `json:"is_deleted" db:"is_deleted"`
	IsBlocked bool   `json:"is_blocked" db:"is_blocked"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
