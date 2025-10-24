// Package mapper contains mappers between models and endpoint requests/responses
package mapper

import (
	"givemegoodcoffee/internal/http/response"
	"givemegoodcoffee/internal/model"
)

type CoffeeSpotMapper struct{}

func NewCoffeeSpotMapper() *CoffeeSpotMapper {
	return &CoffeeSpotMapper{}
}

func (m CoffeeSpotMapper) ToResponse(spot *model.CoffeeSpot) *response.CoffeeSpotResponse {
	return &response.CoffeeSpotResponse{
		Name: spot.Name,
		Type: string(spot.Type),
		Location: response.AddressResponse{
			CountryCode: spot.Location.Address.CountryCode,
			Address:     spot.Location.Address.FormattedAddress,
		},
	}
}
