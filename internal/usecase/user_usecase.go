package usecase

import (
	"database/sql"
	"ecommerce/internal/config"
	"ecommerce/internal/domain"
	"ecommerce/internal/infra/repository"
	"ecommerce/pkg/utils"
	"ecommerce/pkg/utils/jwt"
	"errors"
	"fmt"
	"time"
)

type UserUsecase interface {
	Create(name, email, password string) (*domain.User, error)
	List(page string, limit string, search string) ([]*domain.User, error)
	Login(email, password string) (string, error)
	GetUserById(id string) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	GetMyUserDetails(id string) (*domain.User, error)
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

func (uc *userUsecase) Create(name, emailAddr, password string) (*domain.User, error) {
	existing, err := uc.userRepo.FindByEmail(emailAddr)
	if err == nil && existing != nil {
		return nil, errors.New("user already exists")
	}
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Name:      name,
		Email:     emailAddr,
		Password:  hashedPassword,
		Role:      domain.RoleUser,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := uc.userRepo.Create(user); err != nil {
		return nil, err
	}


	// Try to create OTP and send email
	go func() {
		_, err := uc.otpUC.CreateAndSendEmail(user.ID, user.Name, user.Email)
		if err != nil {
			fmt.Println("OTP creation and send failed:", err)
		}
	}()

	return user, nil
}

func (uc *userUsecase) List(page string, limit string, search string) ([]*domain.User, error) {
	users, err := uc.userRepo.List(page, limit, search)
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
	if !usr.IsVerified {
		return "", errors.New("user is not verified")
	}

	payload := jwt.JwtCustomClaims{
		UserId: usr.ID,
		Email:  usr.Email,
		Role:   usr.Role,
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

func (uc *userUsecase) FindByEmail(email string) (*domain.User, error) {

	user, err := uc.userRepo.FindByEmail(email)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *userUsecase) GetMyUserDetails(id string) (*domain.User, error) {
	user, err := uc.userRepo.GetMyUserDetails(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *userUsecase) BlockUserByAdmin(id string) error {
	isExistingUser, err := uc.userRepo.GetUserById(id)
	if err != nil {
		return err
	}
	if isExistingUser == nil {
		return errors.New("user not found")
	}
	if isExistingUser.IsBlocked {
		return errors.New("user is already blocked")
	}
	err = uc.userRepo.BlockUserByAdmin(id)
	if err != nil {
		return err
	}
	return nil
}

func (uc *userUsecase) UnblockUserByAdmin(id string) error {
	isExistingUser, err := uc.userRepo.GetUserById(id)
	if err != nil {
		return err
	}
	if isExistingUser == nil {
		return errors.New("user not found")
	}
	if !isExistingUser.IsBlocked {
		return errors.New("user is already unblocked")
	}
	err = uc.userRepo.UnblockUserByAdmin(id)
	if err != nil {
		return err
	}
	return nil
}
