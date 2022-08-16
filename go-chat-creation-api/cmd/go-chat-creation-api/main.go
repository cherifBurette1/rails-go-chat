package main

import (
	"log"
	"net/http"

	"github.com/cherifBurette1/rails-go-chat/tree/master/go-chat-creation-api/internal/router"
)

func main() {
	r := router.InitRouter()

	log.Println("Listening on 8080 ......")
	log.Fatal(http.ListenAndServe(":8080", r))
}
