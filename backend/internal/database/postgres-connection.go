package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresConnection struct {
	Pool *pgxpool.Pool
}

func NewPostgresConnection(ctx context.Context, connString string) (*PostgresConnection, error) {
	pool, err := pgxpool.New(ctx, connString)
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
