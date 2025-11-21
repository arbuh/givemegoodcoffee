// Package handler contains handlers for endpoints' requests
package handler

import "givemegoodcoffee/internal/repository"

type Handlers struct {
	HealthHandler     *HealthHandler
	CoffeeSpotHandler *CoffeeSpotHandler
}

func NewHandlers(coffeeSpotRepository repository.CoffeeSpotRepository) *Handlers {
	healthHandler := NewHealthHandler()

	errorHander := NewErrorHander()
	coffeeSpotHandler := NewCoffeeSpotHandler(errorHander, coffeeSpotRepository)

	return &Handlers{HealthHandler: healthHandler, CoffeeSpotHandler: coffeeSpotHandler}
}
