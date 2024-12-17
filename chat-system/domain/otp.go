// domain/otp.go
package domain

import "time"

type OTP struct {
    Key       string    `gorm:"primaryKey"`
    UserID    string    `gorm:"not null"`
    CreatedAt time.Time `gorm:"not null"`
    ExpiredAt time.Time `gorm:"not null"`
}

type OTPRepository interface {
    CreateOTP(otp OTP) error
    VerifyOTP(key string, userID string) (bool, error)
}
