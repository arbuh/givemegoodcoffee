package main

import (
	"context"
	"givemegoodcoffee/internal/config"
	"givemegoodcoffee/internal/database"
	"givemegoodcoffee/internal/http/handler"
	"givemegoodcoffee/internal/http/router"
	"givemegoodcoffee/internal/repository"
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

	// TODO: Create logger depending on env, use a structural logger for deployments
	logger := slog.Default()

	slog.Info("Starting application", "env", env)

	var err error

	var cfg *config.Config
	cfg, err = config.LoadConfig(env)
	if err != nil {
		exitOnError(err, logger)
	}

	slog.Info("Connecting to database")
	var connection database.Connection
	connection, err = database.NewPostgresConnection(ctx, cfg, logger)
	if err != nil {
		exitOnError(err, logger)
	}
	defer connection.Close()

	err = connection.RunMigrations(ctx)
	if err != nil {
		exitOnError(err, logger)
	}

	coffeeSpotRepository := repository.NewCoffeeSpotRepository(connection)
	handlers := handler.NewHandlers(coffeeSpotRepository, logger)
	router := router.NewRouter(handlers)

	slog.Info("Server starting on :8080")
	if err = http.ListenAndServe(":8080", router); err != nil {
		exitOnError(err, logger)
	}
}

func exitOnError(err error, logger *slog.Logger) {
	logger.Error("application failed", "error", err)
	os.Exit(1)
}
