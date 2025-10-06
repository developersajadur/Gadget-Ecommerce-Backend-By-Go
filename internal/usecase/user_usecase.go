package usecase

import (
	"ecommerce/internal/config"
	"ecommerce/internal/domain"
	"ecommerce/internal/infra/repository"
	"ecommerce/pkg/utils"
	"ecommerce/pkg/utils/jwt"
	"errors"
	"fmt"
)

type UserUsecase interface {
	Register(name, email, password string) (*domain.User, error)
	List() ([]*domain.User, error)
	Login(email, password string) (string, error)
	GetUserById(id string) (*domain.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo}
}

func (uc *userUsecase) Register(name, email, password string) (*domain.User, error) {
	existing, _ := uc.userRepo.FindByEmail(email)
	if existing != nil {
		fmt.Println(existing)
		return nil, errors.New("user already exists")
	}

	user := &domain.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return user, err
	}

	user.Password = hashedPassword

	err = uc.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *userUsecase) List() ([]*domain.User, error) {
	users, err := uc.userRepo.List()
	if err != nil {
		return nil, errors.New("Internal error")
	}
	return users, nil
}

func (uc *userUsecase) Login(email, password string) (string, error) {
	usr, err := uc.userRepo.Login(email, password)
	if err != nil {
		return "", errors.New("user not found or invalid credentials")
	}

	payload := jwt.JwtCustomClaims{
		UserId: usr.ID,
		Email:  usr.Email,
	}

	// Generate token
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
