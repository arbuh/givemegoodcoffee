package handler

import (
	"encoding/json"
	"givemegoodcoffee/internal/model"
	"net/http"

	"github.com/gorilla/mux"
)

type CoffeeSpotHandler struct{}

func NewCoffeeSpotHandler() *CoffeeSpotHandler {
	return &CoffeeSpotHandler{}
}

func (h CoffeeSpotHandler) GetCoffeeSpot(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	if id == "" {
		http.Error(w, "The path parameter 'id' is mandatory", http.StatusBadRequest)
		return
	}

	dummySpot := model.CoffeeSpot{
		Id:   id,
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

	err := json.NewEncoder(w).Encode(dummySpot)
	if err != nil {
		http.Error(w, "Cannot serialize a model to JSON", http.StatusInternalServerError)
		return
	}
}
