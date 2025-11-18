// Package repository contains repositories
package repository

import (
	"givemegoodcoffee/internal/database"
	"givemegoodcoffee/internal/model"

	"github.com/google/uuid"
)

type CoffeeSpotRepository interface {
	Save(spot *model.CoffeeSpot)
	Get(id uuid.UUID) *model.CoffeeSpot
}

type CoffeeSpotPostgresRepository struct {
	connection *database.PostgresConnection
}

func NewCoffeeSpotRepository(connection *database.PostgresConnection) CoffeeSpotRepository {
	return &CoffeeSpotPostgresRepository{connection}
}

// Get implements CoffeeSpotRepository.
func (c CoffeeSpotPostgresRepository) Get(id uuid.UUID) *model.CoffeeSpot {
	panic("unimplemented")
}

// Save implements CoffeeSpotRepository.
func (c CoffeeSpotPostgresRepository) Save(spot *model.CoffeeSpot) {
	panic("unimplemented")
}
