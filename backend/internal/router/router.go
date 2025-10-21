// Package router contain definitions for HTTP requests routing
package router

import (
	"givemegoodcoffee/internal/handler"

	"github.com/gorilla/mux"
)

func NewRouter(handlers *handler.Handlers) *mux.Router {
	r := mux.NewRouter()

	api := r.PathPrefix("api/v1").Subrouter()

	api.HandleFunc("health", handlers.HealthHandler.GetHealth).Methods("GET")

	return r
}
