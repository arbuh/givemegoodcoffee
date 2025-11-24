package database

import (
	"context"
	"fmt"
	"givemegoodcoffee/internal/config"
	"log/slog"
	"os"
	"path/filepath"
	"sort"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	migrationDir string = "migration"
)

type PostgresConnection struct {
	Pool   *pgxpool.Pool
	logger *slog.Logger
}

func NewPostgresConnection(ctx context.Context, config *config.Config, logger *slog.Logger) (*PostgresConnection, error) {

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

	connection := &PostgresConnection{Pool: pool, logger: logger}
	return connection, nil
}

func (c *PostgresConnection) Close() {
	c.Pool.Close()
}

func (c *PostgresConnection) RunMigrations(ctx context.Context) error {
	files, err := filepath.Glob(filepath.Join(migrationDir, "*.sql"))
	if err != nil {
		return fmt.Errorf("cannot find migration files: %w", err)
	}

	sort.Strings(files)

	for _, file := range files {

		c.logger.Info("Running migration: " + filepath.Base(file))

		sql, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read %s: %w", file, err)
		}

		if _, err := c.Pool.Exec(ctx, string(sql)); err != nil {
			return fmt.Errorf("failed to execute %s: %w", file, err)
		}
	}

	return nil
}
