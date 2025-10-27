package handler

import (
	"encoding/json"
	"givemegoodcoffee/internal/http/mapper"
	"givemegoodcoffee/internal/model"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type CoffeeSpotHandler struct {
	coffeeSpotMapper *mapper.CoffeeSpotMapper
	errorHander      *ErrorHander
}

func NewCoffeeSpotHandler(errorHander *ErrorHander) *CoffeeSpotHandler {
	coffeeSpotMapper := mapper.NewCoffeeSpotMapper()
	return &CoffeeSpotHandler{coffeeSpotMapper, errorHander}
}

func (h CoffeeSpotHandler) GetCoffeeSpot(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	rawID := vars["id"]
	if rawID == "" {
		h.errorHander.WriteClientError(w, "The path parameter 'id' is mandatory", http.StatusBadRequest)
		return
	}

	id, error := uuid.Parse(rawID)
	if error != nil {
		h.errorHander.WriteClientError(w, "The path parameter 'id' must be a valid UUID", http.StatusBadRequest)
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

	response := h.coffeeSpotMapper.ToResponse(&dummySpot)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		h.errorHander.WriteServerError(w, "Cannot serialize `CoffeeSpotResponse` to JSON")
		return
	}
}
