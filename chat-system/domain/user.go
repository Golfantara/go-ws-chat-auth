// domain/user.go
package domain

type User struct {
    ID       string `gorm:"primaryKey"`
    Username string `gorm:"not null"`
    Password string `gorm:"not null"`
}

type UserRepository interface {
    GetUserByID(id string) (User, error)
    CreateUser(user User) error
    Authenticate(username, password string) (User, error)
}
