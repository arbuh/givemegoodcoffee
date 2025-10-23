// Package model contains domain models
package model

type CoffeeSpot struct {
	Name     string         `json:"name"`
	Type     CoffeeSpotType `json:"type"`
	Location Location       `json:"location"`
}

type CoffeeSpotType string

const (
	Cafe        CoffeeSpotType = "cafe"
	CoffeeShop  CoffeeSpotType = "coffeeShop"
	Restraurant CoffeeSpotType = "restaurant"
	Bakery      CoffeeSpotType = "bakery"
	GasStation  CoffeeSpotType = "gasStation"
	Other       CoffeeSpotType = "other"
)

// TODO: add a unit-test
func (t CoffeeSpotType) IsValid() bool {
	switch t {
	case Cafe, CoffeeShop, Restraurant, Bakery, GasStation, Other:
		return true
	default:
		return false
	}
}
