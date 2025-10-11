package usecase

import (
	"ecommerce/internal/domain"
	"ecommerce/internal/infra/repository"
	"ecommerce/pkg/utils/email"
	"ecommerce/pkg/utils/otp"
	"fmt"
	"time"
)

type OtpUsecase interface {
	CreateAndSendEmail(userID string, userName string, userEmail string) (*domain.Otp, error)
}

type otpUsecase struct {
	otpRepo repository.OtpRepository
}

func NewOtpUsecase(repo repository.OtpRepository) OtpUsecase {
	return &otpUsecase{repo}
}

func (uc *otpUsecase) CreateAndSendEmail(userID string, userName string, userEmail string) (*domain.Otp, error) {
	otpEntry := &domain.Otp{
		UserId:    userID,
		Email:     userEmail,
		Code:      otp.GenerateOTP(),
		ExpiresAt: time.Now().Add(5 * time.Minute),
		CreatedAt: time.Now(),
	}

	createdOtp, err := uc.otpRepo.CreateAndSendEmail(otpEntry)
	if err != nil {
		return nil, err
	}

	// Send OTP email (non-blocking, log failure)
	emailData := map[string]string{
		"Name": userName,
		"OTP":  otpEntry.Code,
	}

	if err := email.SendEmail(userEmail, "Verify Your Account", "templates/otp.html", emailData); err != nil {
		fmt.Println("Failed to send OTP email:", err)
	}

	return createdOtp, nil
}
