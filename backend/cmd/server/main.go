package main

import (
	"givemegoodcoffee/internal/http/handler"
	"givemegoodcoffee/internal/http/router"
	"log"
	"log/slog"
	"net/http"
)

func main() {
	handlers := handler.NewHandlers()
	router := router.NewRouter(handlers)

	slog.Info("Server starting on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		// TODO: use structural logging when we run the application in a server
		log.Fatal(err)
	}
}
