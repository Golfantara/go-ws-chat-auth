// usecase/otp_usecase.go
package usecase

import (
	"time"
	"ws/chat-system/domain"

	"github.com/google/uuid"
)

type OTPUsecase struct {
    repo domain.OTPRepository
}

func NewOTPUsecase(repo domain.OTPRepository) *OTPUsecase {
    return &OTPUsecase{repo: repo}
}

func (u *OTPUsecase) GenerateOTP(userID string) (string, error) {
    otp := domain.OTP{
        Key:       uuid.NewString(),
        UserID:    userID,
        CreatedAt: time.Now(),
        ExpiredAt: time.Now().Add(5 * time.Minute),
    }
    err := u.repo.CreateOTP(otp)
    return otp.Key, err
}

func (u *OTPUsecase) VerifyOTP(key, userID string) (bool, error) {
    return u.repo.VerifyOTP(key, userID)
}
