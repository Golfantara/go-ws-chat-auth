// usecase/chat_usecase.go
package usecase

import (
	"time"
	"ws/chat-system/domain"
)

type ChatUsecase struct {
    repo domain.ChatRepository
}

func NewChatUsecase(repo domain.ChatRepository) *ChatUsecase {
    return &ChatUsecase{repo: repo}
}

func (u *ChatUsecase) SendMessage(senderID, receiverID, message string) error {
    chat := domain.Chat{
        SenderID:  senderID,
        ReceiverID: receiverID,
        Message:   message,
        CreatedAt: time.Now(),
    }
    return u.repo.SendMessage(chat)
}

func (u *ChatUsecase) GetMessages(userID string) ([]domain.Chat, error) {
    return u.repo.GetMessages(userID)
}
