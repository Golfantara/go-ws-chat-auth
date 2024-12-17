// repository/chat_repository.go
package repository

import (
	"ws/chat-system/domain"

	"gorm.io/gorm"
)

type chatRepository struct {
    db *gorm.DB
}

func NewChatRepository(db *gorm.DB) domain.ChatRepository {
    return &chatRepository{db: db}
}

func (r *chatRepository) SendMessage(chat domain.Chat) error {
    return r.db.Create(&chat).Error
}

func (r *chatRepository) GetMessages(userID string) ([]domain.Chat, error) {
    var chats []domain.Chat
    err := r.db.Where("sender_id = ? OR receiver_id = ?", userID, userID).Find(&chats).Error
    return chats, err
}
