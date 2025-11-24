// Package handler contains handlers for endpoints' requests
package handler

import (
	"givemegoodcoffee/internal/repository"
	"log/slog"
)

type Handlers struct {
	HealthHandler     *HealthHandler
	CoffeeSpotHandler *CoffeeSpotHandler
}

func NewHandlers(coffeeSpotRepository repository.CoffeeSpotRepository, logger *slog.Logger) *Handlers {
	healthHandler := NewHealthHandler()

	errorHander := NewErrorHander(logger)
	coffeeSpotHandler := NewCoffeeSpotHandler(errorHander, coffeeSpotRepository)

	return &Handlers{HealthHandler: healthHandler, CoffeeSpotHandler: coffeeSpotHandler}
}
