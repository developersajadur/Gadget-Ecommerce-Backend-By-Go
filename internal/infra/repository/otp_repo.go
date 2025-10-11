package repository

import (
	"ecommerce/internal/domain"
	"time"

	"github.com/jmoiron/sqlx"
)

type OtpRepository interface {
	CreateAndSendEmail(otp *domain.Otp) (*domain.Otp, error)
	VerifyOtp(code string) (*domain.Otp, error)
}

type otpRepository struct {
	db *sqlx.DB
}

func NewOtpRepository(db *sqlx.DB) OtpRepository {
	return &otpRepository{db: db}
}

func (r *otpRepository) CreateAndSendEmail(otp *domain.Otp) (*domain.Otp, error) {
	query := `
		INSERT INTO otps (user_id, email, code, expires_at, created_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	err := r.db.QueryRowx(query, otp.UserId, otp.Email, otp.Code, otp.ExpiresAt, otp.CreatedAt).Scan(&otp.ID)
	if err != nil {
		return nil, err
	}

	return otp, nil
}





func (r *otpRepository) VerifyOtp(code string) (*domain.Otp, error) {
	var otp domain.Otp

	// 1. Find the OTP
	query := `
		SELECT * FROM otps
		WHERE code = $1 AND expires_at > NOW() AND verified = FALSE
		LIMIT 1
	`
	err := r.db.Get(&otp, query, code)
	if err != nil {
		return nil, err
	}

	// Start transaction
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, err
	}

	// 2. Mark OTP as verified
	_, err = tx.Exec(`UPDATE otps SET verified = TRUE, verified_at = NOW() WHERE id = $1`, otp.ID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// 3. Update user as verified
	_, err = tx.Exec(`UPDATE users SET is_verified = TRUE WHERE id = $1 AND is_deleted = false AND is_blocked = false`, otp.UserId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	// 4. Update local struct
	now := time.Now()
	otp.Verified = true
	otp.VerifiedAt = &now

	return &otp, nil
}