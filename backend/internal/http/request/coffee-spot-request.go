// Package request contains DTO's for endpoint requests
package request

type CoffeeSpotRequest struct {
	Name     string          `json:"name"`
	Type     string          `json:"type"`
	Location LocationRequest `json:"location"`
}

type LocationRequest struct {
	CountryCode string `json:"countryCode"`
	Address     string `json:"address"`
}
