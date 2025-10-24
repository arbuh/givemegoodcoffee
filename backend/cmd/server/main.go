package main

import (
	"givemegoodcoffee/internal/http/handler"
	"givemegoodcoffee/internal/http/router"
	"log"
	"net/http"
)

func main() {
	handlers := handler.NewHandlers()
	router := router.NewRouter(handlers)

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
