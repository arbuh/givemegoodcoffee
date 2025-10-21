// Package handler contains handlers for endpoints' requests
package handler

type Handlers struct {
	HealthHandler *HealthHandler
}

func NewHandlers() *Handlers {
	healthHandler := NewHealthHandler()

	return &Handlers{HealthHandler: healthHandler}
}
