package repository

import (
	"ecommerce/internal/domain"

	"github.com/jmoiron/sqlx"
)

type OtpRepository interface {
	Create(otp *domain.Otp) (*domain.Otp, error)
}

type otpRepository struct {
	db *sqlx.DB
}



func NewOtpRepository(db *sqlx.DB) OtpRepository {
	return &otpRepository{db: db}
}

func (r *otpRepository) Create(otp *domain.Otp) (*domain.Otp, error) {
	query := `
		INSERT INTO otps (user_id, code, expires_at, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	err := r.db.QueryRowx(query, otp.UserId, otp.Code, otp.ExpiresAt, otp.CreatedAt).Scan(&otp.ID)
	if err != nil {
		return nil, err
	}

	return otp, nil
}