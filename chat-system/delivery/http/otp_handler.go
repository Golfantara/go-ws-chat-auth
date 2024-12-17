// delivery/http/otp_handler.go
package http

import (
	"encoding/json"
	"net/http"
	"ws/chat-system/usecase"
)

type OTPHandler struct {
    otpUC *usecase.OTPUsecase
}

func NewOTPHandler(otpUC *usecase.OTPUsecase) *OTPHandler {
    return &OTPHandler{otpUC: otpUC}
}

// Generate OTP baru
func (h *OTPHandler) GenerateOTP(w http.ResponseWriter, r *http.Request) {
    var req struct {
        UserID string `json:"user_id"`
    }

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    otpKey, err := h.otpUC.GenerateOTP(req.UserID)
    if err != nil {
        http.Error(w, "Failed to generate OTP", http.StatusInternalServerError)
        return
    }

    response := struct {
        OTPKey string `json:"otp_key"`
    }{
        OTPKey: otpKey,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

// Verifikasi OTP
func (h *OTPHandler) VerifyOTP(w http.ResponseWriter, r *http.Request) {
    var req struct {
        OTPKey  string `json:"otp_key"`
        UserID  string `json:"user_id"`
    }

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    valid, err := h.otpUC.VerifyOTP(req.OTPKey, req.UserID)
    if err != nil {
        http.Error(w, "Failed to verify OTP", http.StatusInternalServerError)
        return
    }

    if !valid {
        http.Error(w, "Invalid OTP or OTP expired", http.StatusUnauthorized)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("OTP valid"))
}
