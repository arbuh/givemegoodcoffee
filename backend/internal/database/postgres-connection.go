package database

import (
	"context"
	"fmt"
	"givemegoodcoffee/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresConnection struct {
	Pool *pgxpool.Pool
}

func NewPostgresConnection(ctx context.Context, config *config.Config) (*PostgresConnection, error) {

	pgxConfig, err := pgxpool.ParseConfig(config.Database.PostgresURL)
	if err != nil {
		return nil, fmt.Errorf("unable to parse connection string: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, pgxConfig)
	if err != nil {
		return nil, fmt.Errorf("unable create a database pool: %w", err)
	}

	if err = pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	connection := &PostgresConnection{Pool: pool}
	return connection, nil
}

func (c *PostgresConnection) Close() {
	c.Pool.Close()
}
