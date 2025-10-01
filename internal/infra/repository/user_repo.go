package repository

import (
	"ecommerce/internal/domain"
	"time"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	Create(user *domain.User) error
	FindByEmail(email string) (*domain.User, error)
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

// Create inserts a new user and returns the inserted ID
func (r *userRepository) Create(user *domain.User) error {
	query := `
		INSERT INTO users (name, email, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	return r.db.QueryRowx(query, user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt).Scan(&user.ID)
}

// FindByEmail fetches a user by email
func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE email=$1", email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
