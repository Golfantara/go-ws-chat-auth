// infrastructure/router.go
package infrastructure

import (
	"log"
	live "net/http"
	"ws/chat-system/delivery/http"
	"ws/chat-system/repository"
	"ws/chat-system/usecase"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func InitRouter() *mux.Router {
    db := InitDB()

    chatRepo := repository.NewChatRepository(db)
    otpRepo := repository.NewOTPRepository(db)
    userRepo := repository.NewUserRepository(db)

    chatUC := usecase.NewChatUsecase(chatRepo)
    otpUC := usecase.NewOTPUsecase(otpRepo)
    userUC := usecase.NewUserUsecase(userRepo)

    chatHandler := http.NewChatHandler(chatUC)
    otpHandler := http.NewOTPHandler(otpUC)
    userHandler := http.NewUserHandler(userUC)

    r := mux.NewRouter()
    r.HandleFunc("/chat/send", chatHandler.SendMessage).Methods("POST")
    r.HandleFunc("/otp/generate", otpHandler.GenerateOTP).Methods("POST")
    r.HandleFunc("/otp/verify", otpHandler.VerifyOTP).Methods("POST")
    r.HandleFunc("/user/register", userHandler.Register).Methods("POST")
    r.HandleFunc("/user/login", userHandler.Login).Methods("POST")
	r.HandleFunc("/ws", chatHandler.HandleWebSocket)

    return r
}

func StartServer() {
    r := InitRouter()
	handler := cors.New(cors.Options{
        AllowedOrigins: []string{"http://127.0.0.1:5500"},
        AllowedMethods: []string{"GET", "POST", "OPTIONS"},
        AllowedHeaders: []string{"Content-Type"},
    }).Handler(r)
    log.Fatal(live.ListenAndServe(":8080", handler))
}
