// Package mapper contains mappers between models and endpoint requests/responses
package mapper

import (
	"givemegoodcoffee/internal/http/request"
	"givemegoodcoffee/internal/http/response"
	"givemegoodcoffee/internal/model"

	"github.com/google/uuid"
)

type CoffeeSpotMapper struct{}

func NewCoffeeSpotMapper() *CoffeeSpotMapper {
	return &CoffeeSpotMapper{}
}

func (m CoffeeSpotMapper) ToResponse(spot *model.CoffeeSpot) *response.CoffeeSpotResponse {
	return &response.CoffeeSpotResponse{
		Name: spot.Name,
		Type: string(spot.Type),
		Location: response.LocationResponse{
			CountryCode: spot.Location.Address.CountryCode,
			Address:     spot.Location.Address.FormattedAddress,
		},
	}
}

func (m CoffeeSpotMapper) FromRequest(request *request.CoffeeSpotRequest) (*model.CoffeeSpot, error) {
	t, err := model.CoffeeSpotTypeFromString(request.Type)
	if err != nil {
		return nil, err
	}

	spot := &model.CoffeeSpot{
		// TODO: implement a special UUID service
		ID:   uuid.UUID{},
		Name: request.Name,
		Type: t,
		Location: model.Location{
			Address: model.Address{
				CountryCode:      request.Location.CountryCode,
				FormattedAddress: request.Location.Address,
			},
		},
	}

	return spot, nil
}
