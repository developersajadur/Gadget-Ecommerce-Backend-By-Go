package usecase

import (
	"ecommerce/internal/domain"
	"ecommerce/internal/infra/repository"
	"ecommerce/pkg/utils/otp"
	"time"
)


type OtpUsecase interface {
	Create(user_id string) (*domain.Otp, error)
}

type otpUsecase struct {
	otpRepo repository.OtpRepository
}

func NewOtpUsecase(repo repository.OtpRepository) OtpUsecase {
	return &otpUsecase{repo}
}


func (uc *otpUsecase) Create(userID string) (*domain.Otp, error) {
	otpEntry := &domain.Otp{
		UserId:    userID,
		Code:      otp.GenerateOTP(),
		ExpiresAt: time.Now().Add(5 * time.Minute), // valid for 5 minutes
		CreatedAt: time.Now(),
	}

	createdOtp, err := uc.otpRepo.Create(otpEntry)
	if err != nil {
		return nil, err
	}

	return createdOtp, nil
}
