// repository/otp_repository.go
package repository

import (
	"time"
	"ws/chat-system/domain"

	"gorm.io/gorm"
)

type otpRepository struct {
    db *gorm.DB
}

func NewOTPRepository(db *gorm.DB) domain.OTPRepository {
    return &otpRepository{db: db}
}

func (r *otpRepository) CreateOTP(otp domain.OTP) error {
    return r.db.Create(&otp).Error
}

func (r *otpRepository) VerifyOTP(key string, userID string) (bool, error) {
    var otp domain.OTP
    result := r.db.Where("key = ? AND user_id = ? AND expired_at > ?", key, userID, time.Now()).First(&otp)
    if result.Error != nil || result.RowsAffected == 0 {
        return false, nil
    }
    return true, nil
}
