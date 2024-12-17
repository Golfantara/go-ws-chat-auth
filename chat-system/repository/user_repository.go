// repository/user_repository.go
package repository

import (
	"ws/chat-system/domain"

	"gorm.io/gorm"
)

type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
    return &userRepository{db: db}
}

func (r *userRepository) GetUserByID(id string) (domain.User, error) {
    var user domain.User
    err := r.db.First(&user, "id = ?", id).Error
    return user, err
}

func (r *userRepository) CreateUser(user domain.User) error {
    return r.db.Create(&user).Error
}

func (r *userRepository) Authenticate(username, password string) (domain.User, error) {
    var user domain.User
    err := r.db.Where("username = ? AND password = ?", username, password).First(&user).Error
    return user, err
}
