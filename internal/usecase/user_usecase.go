package usecase

import (
	"errors"
	"fmt"

	"ecommerce/internal/config"
	"ecommerce/internal/infra/repository"
	"ecommerce/internal/models"
	"ecommerce/pkg/utils"
	"ecommerce/pkg/utils/jwt"
)

type UserUsecase interface {
	Create(name, email, password string) (*models.User, error)
	List(page, limit, search string) ([]*models.User, error)
	Login(email, password string) (string, error)
	GetUserById(id string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	GetMyUserDetails(id string) (*models.User, error)
	BlockUserByAdmin(id string) error
	UnblockUserByAdmin(id string) error
}

type userUsecase struct {
	userRepo repository.UserRepository
	otpUC    OtpUsecase
}

func NewUserUsecase(userRepo repository.UserRepository, otpUC OtpUsecase) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
		otpUC:    otpUC,
	}
}

func (uc *userUsecase) Create(name, emailAddr, password string) (*models.User, error) {
	existing, _ := uc.userRepo.FindByEmail(emailAddr)
	if existing != nil {
		return nil, errors.New("user already exists")
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:      name,
		Email:     emailAddr,
		Password:  hashedPassword,
	}

	if err := uc.userRepo.Create(user); err != nil {
		return nil, err
	}

	// Asynchronously create OTP and send email
	go func() {
		_, err := uc.otpUC.CreateAndSendEmail(user.ID, user.Name, user.Email)
		if err != nil {
			fmt.Println("OTP creation and send failed:", err)
		}
	}()

	return user, nil
}

func (uc *userUsecase) List(page, limit, search string) ([]*models.User, error) {
	return uc.userRepo.List(page, limit, search, map[string]interface{}{})
}

func (uc *userUsecase) Login(email, password string) (string, error) {
	user, err := uc.userRepo.Login(email, password)
	if err != nil || user == nil {
		return "", errors.New("invalid credentials")
	}

	if user.IsDeleted {
		return "", errors.New("user not found")
	}
	if user.IsBlocked {
		return "", errors.New("user is blocked")
	}
	if !user.IsVerified {
		return "", errors.New("user is not verified")
	}

	payload := jwt.JwtCustomClaims{
		UserId: user.ID,
		Email:  user.Email,
		Role:   user.Role,
	}

	token, err := jwt.GenerateJWT([]byte(config.ENV.JWTSecret), payload)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (uc *userUsecase) GetUserById(id string) (*models.User, error) {
	return uc.userRepo.GetUserById(id)
}

func (uc *userUsecase) FindByEmail(email string) (*models.User, error) {
	return uc.userRepo.FindByEmail(email)
}

func (uc *userUsecase) GetMyUserDetails(id string) (*models.User, error) {
	return uc.userRepo.GetMyUserDetails(id)
}

func (uc *userUsecase) BlockUserByAdmin(id string) error {
	user, err := uc.userRepo.GetUserById(id)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}
	if user.IsBlocked {
		return errors.New("user is already blocked")
	}

	return uc.userRepo.BlockUserByAdmin(id)
}

func (uc *userUsecase) UnblockUserByAdmin(id string) error {
	user, err := uc.userRepo.GetUserById(id)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}
	if !user.IsBlocked {
		return errors.New("user is already unblocked")
	}

	return uc.userRepo.UnblockUserByAdmin(id)
}
