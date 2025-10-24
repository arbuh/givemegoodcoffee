// Package handler contains handlers for endpoints' requests
package handler

type Handlers struct {
	HealthHandler     *HealthHandler
	CoffeeSpotHandler *CoffeeSpotHandler
}

func NewHandlers() *Handlers {
	healthHandler := NewHealthHandler()

	coffeeSpotHandler := NewCoffeeSpotHandler()

	return &Handlers{HealthHandler: healthHandler, CoffeeSpotHandler: coffeeSpotHandler}
}
