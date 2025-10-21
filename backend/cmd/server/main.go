package main

import (
	"fmt"
	"givemegoodcoffee/internal/handler"
	"givemegoodcoffee/internal/router"
	"net/http"
)

func main() {
	handlers := handler.NewHandlers()
	router := router.NewRouter(handlers)

	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println("Server error", err)
	}
}
