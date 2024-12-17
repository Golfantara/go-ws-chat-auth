// delivery/http/user_handler.go
package http

import (
	"encoding/json"
	"net/http"
	"ws/chat-system/usecase"
)

type UserHandler struct {
    userUC *usecase.UserUsecase
}

func NewUserHandler(userUC *usecase.UserUsecase) *UserHandler {
    return &UserHandler{userUC: userUC}
}

// Register user baru
func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    user, err := h.userUC.Register(req.Username, req.Password)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    response := struct {
        ID       string `json:"id"`
        Username string `json:"username"`
    }{
        ID:       user.ID,
        Username: user.Username,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

// Login user dengan username dan password
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    user, err := h.userUC.Login(req.Username, req.Password)
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }

    response := struct {
        ID       string `json:"id"`
        Username string `json:"username"`
    }{
        ID:       user.ID,
        Username: user.Username,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}
