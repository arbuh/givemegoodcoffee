// Package model contains domain models
package model

import (
	"fmt"

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

var allTypes = []CoffeeSpotType{Cafe, CoffeeShop, Restraurant, Bakery, GasStation, Other}

// TODO: add a unit-test
func (t CoffeeSpotType) IsValid() bool {
	switch t {
	case Cafe, CoffeeShop, Restraurant, Bakery, GasStation, Other:
		return true
	default:
		return false
	}
}

func CoffeeSpotTypeFromString(s string) (CoffeeSpotType, error) {
	t := CoffeeSpotType(s)

	if t.IsValid() {
		return t, nil
	} else {
		return "", UnsupportedCoffeeSpotTypeError{unsupportedType: s}
	}
}

type UnsupportedCoffeeSpotTypeError struct {
	unsupportedType string
}

func (e UnsupportedCoffeeSpotTypeError) Error() string {
	return fmt.Sprintf("Unsupported coffee spot type '%s', it should be one of %v", e.unsupportedType, allTypes)
}
