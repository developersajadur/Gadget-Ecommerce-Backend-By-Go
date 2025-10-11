package repository

import (
	"ecommerce/internal/domain"
	"github.com/jmoiron/sqlx"
)

type OtpRepository interface {
	CreateAndSendEmail(otp *domain.Otp) (*domain.Otp, error)
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
