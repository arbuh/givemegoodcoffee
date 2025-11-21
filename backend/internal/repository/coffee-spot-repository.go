// Package repository contains repositories
package repository

import (
	"context"
	"errors"
	"fmt"
	"givemegoodcoffee/internal/database"
	"givemegoodcoffee/internal/model"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type CoffeeSpotRepository interface {
	Save(ctx context.Context, spot *model.CoffeeSpot) error
	Get(ctx context.Context, id uuid.UUID) (*model.CoffeeSpot, error)
}

type CoffeeSpotPostgresRepository struct {
	connection *database.PostgresConnection
}

func NewCoffeeSpotRepository(connection database.Connection) CoffeeSpotRepository {
	pgConnection, _ := connection.(*database.PostgresConnection)
	return &CoffeeSpotPostgresRepository{pgConnection}
}

// Get implements CoffeeSpotRepository.
func (c CoffeeSpotPostgresRepository) Get(ctx context.Context, id uuid.UUID) (*model.CoffeeSpot, error) {
	query := "SELECT data from coffee_spots WHERE id = $1"

	var spot model.CoffeeSpot
	err := c.connection.Pool.QueryRow(ctx, query, id).Scan(&spot)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("failed to get coffee spot: %w", err)
	}

	return &spot, nil
}

// Save implements CoffeeSpotRepository.
func (c CoffeeSpotPostgresRepository) Save(ctx context.Context, spot *model.CoffeeSpot) error {
	query := "INSERT INTO coffee_spots (id, data) VALUES $1"

	_, err := c.connection.Pool.Exec(ctx, query, spot.ID, spot)
	if err != nil {
		return fmt.Errorf("failed to save coffee spot: %w", err)
	}

	return nil
}
