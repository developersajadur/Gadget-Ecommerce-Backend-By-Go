package domain

import "time"

type Otp struct {
	ID         string    `json:"id" db:"id"`
	UserId     string    `json:"user_id" db:"user_id"`
	Email      string    `json:"email" db:"email"`
	Code       string    `json:"code" db:"code"`
	ExpiresAt  time.Time `json:"expires_at" db:"expires_at"`
	Verified   bool      `json:"verified" db:"verified"`
	VerifiedAt *time.Time `json:"verified_at,omitempty" db:"verified_at"` 
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}
