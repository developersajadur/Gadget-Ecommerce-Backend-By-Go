package usecase

import (
	"database/sql"
	"ecommerce/internal/config"
	"ecommerce/internal/domain"
	"ecommerce/internal/infra/repository"
	"ecommerce/pkg/utils"
	"ecommerce/pkg/utils/jwt"
	"errors"
)

type UserUsecase interface {
	Create(name, email, password string) (*domain.User, error)
	List(page string, limit string) ([]*domain.User, error)
	Login(email, password string) (string, error)
	GetUserById(id string) (*domain.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo}
}

func (uc *userUsecase) Create(name, email, password string) (*domain.User, error) {
	existing, err := uc.userRepo.FindByEmail(email)
	if err == nil && existing != nil {
		return nil, errors.New("user already exists")
	}

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	user := &domain.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hashedPassword

	err = uc.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}


func (uc *userUsecase) List(page string, limit string) ([]*domain.User, error) {
	users, err := uc.userRepo.List(page, limit)
	if err != nil {
		return nil, errors.New("internal error")
	}
	return users, nil
}

func (uc *userUsecase) Login(email, password string) (string, error) {
	usr, err := uc.userRepo.Login(email, password)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if usr == nil {
		return "", errors.New("user not found")
	}

	if usr.IsDeleted {
		return "", errors.New("user not found")
	}

	if usr.IsBlocked {
		return "", errors.New("user is blocked")
	}

	payload := jwt.JwtCustomClaims{
		UserId: usr.ID,
		Email:  usr.Email,
	}

	token, err := jwt.GenerateJWT([]byte(config.ENV.JWTSecret), payload)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (uc *userUsecase) GetUserById(id string) (*domain.User, error) {
	user, err := uc.userRepo.GetUserById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
