package repository

import (
	"ecommerce/internal/models"
	"ecommerce/pkg/querybuilder"
	"ecommerce/pkg/utils"
	"errors"
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	List(page, limit, search string, filters map[string]interface{}) ([]*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Login(email, password string) (*models.User, error)
	GetUserById(id string) (*models.User, error)
	GetMyUserDetails(id string) (*models.User, error)
	BlockUserByAdmin(id string) error
	UnblockUserByAdmin(id string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// Create inserts a new user
func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

// List returns paginated, filtered, and searched users
func (r *userRepository) List(page, limit, search string, filters map[string]interface{}) ([]*models.User, error) {
	pageInt, limitInt := utils.ParsePagination(page, limit)

	qb := querybuilder.New(r.db.Model(&models.User{}).Where("is_deleted = ?", false))
	qb.SetPagination(pageInt, limitInt)
	qb.SetSearch(search, []string{"name", "email"})
	qb.SetFilters(filters)

	var users []*models.User
	err := qb.Build().Find(&users).Error
	return users, err
}

// FindByEmail fetches a single active user by email
func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ? AND is_deleted = false", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Login validates credentials
func (r *userRepository) Login(email, password string) (*models.User, error) {
	user, err := r.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if !utils.CheckPassword(user.Password, password) {
		return nil, errors.New("invalid password")
	}
	return user, nil
}

// GetUserById fetches a user by ID
func (r *userRepository) GetUserById(id string) (*models.User, error) {
	var user models.User
	err := r.db.Where("id = ? AND is_deleted = false", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetMyUserDetails returns the same as GetUserById
func (r *userRepository) GetMyUserDetails(id string) (*models.User, error) {
	return r.GetUserById(id)
}

// BlockUserByAdmin sets is_blocked to true
func (r *userRepository) BlockUserByAdmin(id string) error {
	return r.db.Model(&models.User{}).Where("id = ?", id).
		Updates(map[string]interface{}{
			"is_blocked": true,
			"updated_at": time.Now(),
		}).Error
}

// UnblockUserByAdmin sets is_blocked to false
func (r *userRepository) UnblockUserByAdmin(id string) error {
	return r.db.Model(&models.User{}).Where("id = ?", id).
		Updates(map[string]interface{}{
			"is_blocked": false,
			"updated_at": time.Now(),
		}).Error
}
