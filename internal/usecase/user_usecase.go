package usecase

import (
	"ecommerce/internal/domain"
	"ecommerce/internal/infra/repository"
	"ecommerce/pkg/utils"
	"errors"
	"fmt"
)

type UserUsecase interface {
	Register(name, email, password string) (*domain.User, error)
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
