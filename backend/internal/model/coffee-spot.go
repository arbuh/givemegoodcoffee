// Package model contains domain models
package model

import (
	"github.com/google/uuid"
)

type CoffeeSpot struct {
	ID       uuid.UUID
	Name     string
	Type     CoffeeSpotType
	Location Location
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

/*
We don't need this functionality yet
// TODO: add a unit-test
func (t CoffeeSpotType) IsValid() bool {
	switch t {
	case Cafe, CoffeeShop, Restraurant, Bakery, GasStation, Other:
		return true
	default:
		return false
	}
}
*/
