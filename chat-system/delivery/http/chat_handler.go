// delivery/http/chat_handler.go
package http

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"ws/chat-system/usecase"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

type ChatHandler struct {
    chatUC *usecase.ChatUsecase
}

func NewChatHandler(chatUC *usecase.ChatUsecase) *ChatHandler {
    return &ChatHandler{chatUC: chatUC}
}

func (h *ChatHandler) SendMessage(w http.ResponseWriter, r *http.Request) {
    var req struct {
        SenderID   string `json:"sender_id"`
        ReceiverID string `json:"receiver_id"`
        Message    string `json:"message"`
    }

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    err := h.chatUC.SendMessage(req.SenderID, req.ReceiverID, req.Message)
    if err != nil {
        http.Error(w, "Failed to send message", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Message sent"))
}


type MessageEvent struct {
	Type    string `json:"type"`
	Payload struct {
		Message string `json:"message"`
		From    string `json:"from"`
		Sent    time.Time  `json:"sent"`
	} `json:"payload"`
}

func (h *ChatHandler) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("Error upgrading to WebSocket:", err)
        return
    }
    defer conn.Close()


    for {
        _, msg, err := conn.ReadMessage()
        if err != nil {
            log.Println("Error reading message:", err)
            break
        }

        log.Printf("Received message: %s", msg)

		var wsMessage MessageEvent
        if err := json.Unmarshal(msg, &wsMessage); err != nil {
            log.Println("Error unmarshaling message:", err)
            continue
        }

		messsageContent := wsMessage.Payload.Message
		sender := wsMessage.Payload.From
		sentTimeStamp := wsMessage.Payload.Sent

		sentTimeStamp = time.Now()
		log.Printf("Message: %s", messsageContent)
        log.Printf("Sender: %s", sender)
        log.Printf("Sent timestamp: %v", sentTimeStamp)
        // Create the response message structure
        response := MessageEvent{
            Type: "new_message",
            Payload: struct {
                Message string `json:"message"`
                From    string `json:"from"`
                Sent    time.Time  `json:"sent"`
            }{
                Message: messsageContent,
                From:    sender,
                Sent:    sentTimeStamp,
            },
        }

        // Convert the event to JSON
        responseBytes, err := json.Marshal(response)
        if err != nil {
            log.Println("Error marshaling event:", err)
            break
        }

        // Send the JSON message over WebSocket
        err = conn.WriteMessage(websocket.TextMessage, responseBytes)
        if err != nil {
            log.Println("Error sending message:", err)
            break
        }
    }
}
