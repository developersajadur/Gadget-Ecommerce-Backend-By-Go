package domain

import "time"

type User struct {
	ID        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"-" db:"password"`
	Role      string    `json:"role" db:"role"` // "user" or "admin"
	IsDeleted bool      `json:"is_deleted" db:"is_deleted"`
	IsBlocked bool      `json:"is_blocked" db:"is_blocked"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

const (
	RoleUser  = "user"
	RoleAdmin = "admin"
)

var Role = struct {
	User  string
	Admin string
}{
	User:  RoleUser,
	Admin: RoleAdmin,
}
