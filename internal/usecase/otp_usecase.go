package usecase

import (
	"ecommerce/internal/models"
	"ecommerce/internal/infra/repository"
	"ecommerce/pkg/utils/email"
	"ecommerce/pkg/utils/otp"
	"fmt"
	"time"
)

type OtpUsecase interface {
	CreateAndSendEmail(userID, userName, userEmail string) (*models.Otp, error)
	VerifyOtp(code string) (*models.Otp, error)
}

type otpUsecase struct {
	otpRepo repository.OtpRepository
}

func NewOtpUsecase(repo repository.OtpRepository) OtpUsecase {
	return &otpUsecase{otpRepo: repo}
}

// CreateAndSendEmail generates an OTP, saves it to DB, and sends it via email.
func (uc *otpUsecase) CreateAndSendEmail(userID, userName, userEmail string) (*models.Otp, error) {
	otpEntry := &models.Otp{
		UserID:    userID,
		Email:     userEmail,
		Code:      otp.GenerateOTP(),
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}

	createdOtp, err := uc.otpRepo.CreateAndSendEmail(otpEntry)
	if err != nil {
		return nil, err
	}

	// Send OTP email asynchronously
	emailData := map[string]string{
		"Name": userName,
		"OTP":  createdOtp.Code,
	}
	go func() {
		if err := email.SendEmail(userEmail, "Verify Your Account", "templates/otp.html", emailData); err != nil {
			fmt.Println("Failed to send OTP email:", err)
		}
	}()

	return createdOtp, nil
}

// VerifyOtp validates the OTP code, marks it verified, and updates the user's verified status.
func (uc *otpUsecase) VerifyOtp(code string) (*models.Otp, error) {
	return uc.otpRepo.VerifyOtp(code)
}
