package handler

import (
	"encoding/json"
	"givemegoodcoffee/internal/http/mapper"
	"givemegoodcoffee/internal/model"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type CoffeeSpotHandler struct{ CoffeeSpotMapper *mapper.CoffeeSpotMapper }

func NewCoffeeSpotHandler() *CoffeeSpotHandler {
	coffeeSpotMapper := mapper.NewCoffeeSpotMapper()
	return &CoffeeSpotHandler{coffeeSpotMapper}
}

func (h CoffeeSpotHandler) GetCoffeeSpot(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	rawID := vars["id"]
	if rawID == "" {
		http.Error(w, "The path parameter 'id' is mandatory", http.StatusBadRequest)
		return
	}

	id, error := uuid.Parse(rawID)
	if error != nil {
		http.Error(w, "The path parameter 'id' must be a valid UUID", http.StatusBadRequest)
		return
	}

	dummySpot := model.CoffeeSpot{
		ID:   id,
		Name: "Frappie-Lattie Cafe",
		Type: model.CoffeeShop,
		Location: model.Location{
			GeoPoint: model.GeoPoint{
				Lat: "0.0",
				Lon: "0.0",
			},
			Address: model.Address{
				CountryCode:      "nl",
				FormattedAddress: "Nijnte pleintje 7, Utrecht",
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")

	response := h.CoffeeSpotMapper.ToResponse(&dummySpot)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Cannot serialize a model to JSON", http.StatusInternalServerError)
		return
	}
}
