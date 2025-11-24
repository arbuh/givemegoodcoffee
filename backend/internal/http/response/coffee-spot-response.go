// Package response contains DTO's for endpoint responses
package response

type CoffeeSpotResponse struct {
	ID       string           `json:"id"`
	Name     string           `json:"name"`
	Type     string           `json:"type"`
	Location LocationResponse `json:"location"`
}

type LocationResponse struct {
	CountryCode string `json:"countryCode"`
	Address     string `json:"address"`
}
