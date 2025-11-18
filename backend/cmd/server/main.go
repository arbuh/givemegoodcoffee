package main

import (
	"context"
	"givemegoodcoffee/internal/config"
	"givemegoodcoffee/internal/database"
	"givemegoodcoffee/internal/http/handler"
	"givemegoodcoffee/internal/http/router"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	ctx := context.Background()

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = config.DEV
	}

	slog.Info("Starting application", "env", env)

	var err error

	var cfg *config.Config
	cfg, err = config.LoadConfig(env)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	_, err = database.NewPostgresConnection(ctx, cfg)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	handlers := handler.NewHandlers()
	router := router.NewRouter(handlers)

	slog.Info("Server starting on :8080")
	if err = http.ListenAndServe(":8080", router); err != nil {
		// TODO: use structural logging when we run the application in a server
		log.Fatal(err)
		os.Exit(1)
	}
}
