// Package response contains DTO's for endpoint responses
package response

type CoffeeSpotResponse struct {
	Name     string          `json:"name"`
	Type     string          `json:"type"`
	Location AddressResponse `json:"location"`
}

type AddressResponse struct {
	CountryCode      string `json:"countryCode"`
	Address string `json:"address"`
}
