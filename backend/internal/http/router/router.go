// Package router contain definitions for HTTP requests routing
package router

import (
	"givemegoodcoffee/internal/http/handler"
	"givemegoodcoffee/internal/http/middleware"

	"github.com/gorilla/mux"
)

func NewRouter(handlers *handler.Handlers) *mux.Router {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()

	// Health endpoints
	api.HandleFunc("/health", handlers.HealthHandler.HealthCheck).Methods("GET")

	// Coffee spot endpoints
	api.HandleFunc("/coffeespot/{id}", handlers.CoffeeSpotHandler.GetCoffeeSpot).Methods("GET")

	return api
}
