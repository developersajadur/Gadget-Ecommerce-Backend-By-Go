package repository

import (
	"ecommerce/internal/domain"
	"ecommerce/pkg/querybuilder"
	"ecommerce/pkg/utils"
	"errors"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	Create(user *domain.User) error
	List(page string, limit string, search string, filters map[string]string) ([]*domain.User, error)
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

func (r *userRepository) List(page, limit, search string, filters map[string]string) ([]*domain.User, error) {
	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)

	qb := querybuilder.New("SELECT * FROM users")
	qb.SetPagination(pageInt, limitInt)
	qb.SetSearch(search, []string{"name", "email"})
	qb.AddFilters(filters)

	query, args := qb.Build()

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



