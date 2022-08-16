package router

import (
	"github.com/cherifBurette1/rails-go-chat/tree/master/go-chat-creation-api/configs"
	"github.com/cherifBurette1/rails-go-chat/tree/master/go-chat-creation-api/internal/handlers"
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc(configs.ChatsRoute, handlers.CreateChat).Methods("POST")
	router.HandleFunc(configs.MessagesRoute, handlers.CreateMessage).Methods("POST")

	return router
}
