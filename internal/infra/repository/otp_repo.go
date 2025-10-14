package repository

import (
	"ecommerce/internal/models"
	"time"
	"errors"
	"gorm.io/gorm"
)

type OtpRepository interface {
	CreateAndSendEmail(otp *models.Otp) (*models.Otp, error)
	VerifyOtp(code string) (*models.Otp, error)
}

type otpRepository struct {
	db *gorm.DB
}

func NewOtpRepository(db *gorm.DB) OtpRepository {
	return &otpRepository{db: db}
}

func (r *otpRepository) CreateAndSendEmail(otp *models.Otp) (*models.Otp, error) {
	if err := r.db.Create(otp).Error; err != nil {
		return nil, err
	}
	return otp, nil
}

func (r *otpRepository) VerifyOtp(code string) (*models.Otp, error) {
	var otp models.Otp

	err := r.db.Where("code = ? AND expires_at > ? AND verified = false", code, time.Now()).First(&otp).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	// Transaction: mark OTP verified and update user
	err = r.db.Transaction(func(tx *gorm.DB) error {
		now := time.Now()

		if err := tx.Model(&models.Otp{}).Where("id = ?", otp.ID).
			Updates(map[string]interface{}{
				"verified":    true,
				"verified_at": &now,
			}).Error; err != nil {
			return err
		}

		if err := tx.Model(&models.User{}).Where("id = ? AND is_deleted = false AND is_blocked = false", otp.UserID).
			Update("is_verified", true).Error; err != nil {
			return err
		}

		// update struct for return
		otp.Verified = true
		otp.VerifiedAt = &now

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &otp, nil
}
