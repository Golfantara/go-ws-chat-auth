// domain/chat.go
package domain

import "time"

type Chat struct {
    ID        string    `gorm:"primaryKey"`
    SenderID  string    `gorm:"not null"`
    ReceiverID string   `gorm:"not null"`
    Message   string    `gorm:"not null"`
    CreatedAt time.Time `gorm:"not null"`
}

type ChatRepository interface {
    SendMessage(chat Chat) error
    GetMessages(userID string) ([]Chat, error)
}
