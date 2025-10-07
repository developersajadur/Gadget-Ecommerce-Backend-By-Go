package repository

import (
	"ecommerce/internal/domain"
	"ecommerce/pkg/utils"
	"errors"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	Create(user *domain.User) error
	List(page string, limit string) ([]*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	Login(email string, password string) (*domain.User, error)
	GetUserById(id string) (*domain.User, error)
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
		INSERT INTO users (name, email, password, role, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	// Pass all 6 values including role
	return r.db.QueryRowx(query, user.Name, user.Email, user.Password, user.Role, user.CreatedAt, user.UpdatedAt).Scan(&user.ID)
}


func (r *userRepository) List(page, limit string) ([]*domain.User, error) {
	const (
		defaultPage  = 1
		defaultLimit = 20
	)

	toInt := func(s string, def int) int {
		v, err := strconv.Atoi(s)
		if err != nil || v < 1 {
			return def
		}
		return v
	}

	pageInt := toInt(page, defaultPage)
	limitInt := toInt(limit, defaultLimit)
	offset := (pageInt - 1) * limitInt

	query := `
		SELECT *
		FROM users
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	var users []domain.User
	if err := r.db.Select(&users, query, limitInt, offset); err != nil {
		return nil, err
	}

	userPtrs := make([]*domain.User, len(users))
	for i := range users {
		userPtrs[i] = &users[i]
	}

	return userPtrs, nil
}

// FindByEmail fetches a user by email
func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	query := "SELECT * FROM users WHERE email = $1"
	err := r.db.Get(&user, query, email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Login(email string, password string) (*domain.User, error) {
	usr, err := r.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if !utils.CheckPassword(usr.Password, password) {
		return nil, errors.New("invalid password")
	}
	return usr, nil
}

func (r *userRepository) GetUserById(id string) (*domain.User, error) {
	var user domain.User
	query := "SELECT * FROM users WHERE is_deleted = false AND id = $1"

	err := r.db.Get(&user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
