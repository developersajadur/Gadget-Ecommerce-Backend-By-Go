package repository

import (
	"ecommerce/internal/domain"
	"ecommerce/pkg/utils"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	Create(user *domain.User) error
	List(page string, limit string, search string) ([]*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	Login(email string, password string) (*domain.User, error)
	GetUserById(id string) (*domain.User, error)
	GetMyUserDetails(id string) (*domain.User, error)
	BlockUserByAdmin(id string) error
	UnblockUserByAdmin(id string) error
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

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

func (r *userRepository) List(page, limit, search string) ([]*domain.User, error) {
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

	// Base query
	query := `
		SELECT *
		FROM users
	`

	args := []interface{}{}
	argCount := 1

	// If search exists, add WHERE
	if search != "" {
		query += fmt.Sprintf(`WHERE name ILIKE $%d OR email ILIKE $%d `, argCount, argCount+1)
		args = append(args, "%"+search+"%", "%"+search+"%")
		argCount += 2
	}

	// Pagination
	query += fmt.Sprintf("ORDER BY created_at DESC LIMIT $%d OFFSET $%d", argCount, argCount+1)
	args = append(args, limitInt, offset)

	var users []domain.User
	if err := r.db.Select(&users, query, args...); err != nil {
		return nil, err
	}

	userPtrs := make([]*domain.User, len(users))
	for i := range users {
		userPtrs[i] = &users[i]
	}

	return userPtrs, nil
}


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

func (r *userRepository) GetMyUserDetails(id string) (*domain.User, error) {

	var user domain.User
	query := "SELECT * FROM users WHERE is_deleted = false AND id = $1"

	err := r.db.Get(&user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) BlockUserByAdmin(id string) error {
	now := time.Now()
	query := "UPDATE users SET is_blocked = true, updated_at = $1 WHERE id = $2"
	_, err := r.db.Exec(query, now, id)
	return err
}

func (r *userRepository) UnblockUserByAdmin(id string) error {
	now := time.Now()
	query := "UPDATE users SET is_blocked = false, updated_at = $1 WHERE id = $2"
	_, err := r.db.Exec(query, now, id)
	return err
}

