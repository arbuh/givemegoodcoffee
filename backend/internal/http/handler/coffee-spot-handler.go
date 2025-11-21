package handler

import (
	"encoding/json"
	"fmt"
	"givemegoodcoffee/internal/http/mapper"
	"givemegoodcoffee/internal/http/request"
	"givemegoodcoffee/internal/http/util"
	"givemegoodcoffee/internal/model"
	"givemegoodcoffee/internal/repository"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type CoffeeSpotHandler struct {
	coffeeSpotMapper     *mapper.CoffeeSpotMapper
	coffeeSpotRepository repository.CoffeeSpotRepository
	errorHander          *ErrorHander
}

func NewCoffeeSpotHandler(errorHander *ErrorHander, coffeeSpotRepository repository.CoffeeSpotRepository) *CoffeeSpotHandler {
	coffeeSpotMapper := mapper.NewCoffeeSpotMapper()
	return &CoffeeSpotHandler{coffeeSpotMapper, coffeeSpotRepository, errorHander}
}

func (h CoffeeSpotHandler) PostCoffeeSpot(w http.ResponseWriter, r *http.Request) {
	var request request.CoffeeSpotRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		h.errorHander.HandleClientError(w, r, "Cannot parse the request: "+err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: move ID generation to a special service
	id, err := uuid.NewRandom()
	if err != nil {
		h.errorHander.HandleServerError(w, r, "Cannot generate ID for spot: "+err.Error())
		return
	}

	var spot *model.CoffeeSpot
	spot, err = h.coffeeSpotMapper.FromRequest(id, &request)
	if err != nil {
		h.errorHander.HandleClientError(w, r, "Cannot map the request: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = h.coffeeSpotRepository.Save(r.Context(), spot)
	if err != nil {
		h.errorHander.HandleServerError(w, r, "Cannot save coffee spot: "+err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	responseBody := fmt.Sprintf(`{"id":"%s"}`, id)
	w.Write([]byte(responseBody))
}

func (h CoffeeSpotHandler) GetCoffeeSpot(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	rawID := vars["id"]
	if rawID == "" {
		h.errorHander.HandleClientError(w, r, "The path parameter 'id' is mandatory", http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(rawID)
	if err != nil {
		h.errorHander.HandleClientError(w, r, "The path parameter 'id' must be a valid UUID", http.StatusBadRequest)
		return
	}

	spot, err := h.coffeeSpotRepository.Get(r.Context(), id)
	if err != nil {
		h.errorHander.HandleServerError(w, r, "Cannot get coffee spot: "+err.Error())
		return
	}

	// dummySpot := model.CoffeeSpot{
	// 	ID:   id,
	// 	Name: "Frappie-Lattie Cafe",
	// 	Type: model.CoffeeShop,
	// 	Location: model.Location{
	// 		GeoPoint: model.GeoPoint{
	// 			Lat: "0.0",
	// 			Lon: "0.0",
	// 		},
	// 		Address: model.Address{
	// 			CountryCode:      "nl",
	// 			FormattedAddress: "Nijnte pleintje 7, Utrecht",
	// 		},
	// 	},
	// }

	w.Header().Set("Content-Type", "application/json")

	if spot == nil {
		util.WriteNotFound(w)
		return
	}

	response := h.coffeeSpotMapper.ToResponse(spot)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		h.errorHander.HandleServerError(w, r, "Cannot serialize `CoffeeSpotResponse` to JSON")
		return
	}
}
