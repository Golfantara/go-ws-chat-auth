// usecase/user_usecase.go
package usecase

import (
	"errors"
	"ws/chat-system/domain"

	"github.com/google/uuid"
)

type UserUsecase struct {
    repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) *UserUsecase {
    return &UserUsecase{repo: repo}
}

// Register user baru
func (u *UserUsecase) Register(username, password string) (domain.User, error) {
    // Validasi username dan password (contoh sederhana)
    if username == "" || password == "" {
        return domain.User{}, errors.New("username dan password tidak boleh kosong")
    }

    user := domain.User{
        Username: username,
        Password: password, 
		ID: uuid.NewString(),// Perhatikan, untuk produksi gunakan hashing password!
    }

    err := u.repo.CreateUser(user)
    if err != nil {
        return domain.User{}, err
    }

    return user, nil
}

// Login user dengan username dan password
func (u *UserUsecase) Login(username, password string) (domain.User, error) {
    user, err := u.repo.Authenticate(username, password)
    if err != nil {
        return domain.User{}, err
    }

    if user.ID == "" {
        return domain.User{}, errors.New("username atau password salah")
    }

    return user, nil
}
