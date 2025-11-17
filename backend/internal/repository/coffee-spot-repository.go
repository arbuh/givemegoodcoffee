// Package repository contains repositories
package repository

import (
	"context"
	"givemegoodcoffee/internal/model"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/google/uuid"
)

type CoffeeSpotRepository interface {
	Save(spot model.CoffeeSpot)
	Get(id uuid.UUID)
}

type CoffeeSpotPostgresRepository struct {
	pool *pgxpool.Pool
}

// TODO: Create a separate instance for the pool
func NewCoffeeSpotRepository() (*CoffeeSpotRepository, error) {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	defer dbpool.Close()

	return &CoffeeSpotPostgresRepository{pool: dbpool}, nil
}

// Get implements CoffeeSpotRepository.
func (c CoffeeSpotPostgresRepository) Get(id uuid.UUID) {
	panic("unimplemented")
}

// Save implements CoffeeSpotRepository.
func (c CoffeeSpotPostgresRepository) Save(spot model.CoffeeSpot) {
	panic("unimplemented")
}
