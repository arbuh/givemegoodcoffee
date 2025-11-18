// Package config contains types for configurations and functions to initialize them
package config

import (
	"fmt"
	"os"
)

const (
	DEV string = "dev"
)

type Config struct {
	Database *DatabaseConfig
}

type DatabaseConfig struct {
	PostgresURL string
}

func LoadConfig(env string) (*Config, error) {
	databaseConfig, err := loadDatabaseConfig(env)
	if err != nil {
		return nil, err
	}

	config := &Config{
		Database: databaseConfig,
	}
	return config, nil
}

func loadDatabaseConfig(env string) (*DatabaseConfig, error) {
	url := os.Getenv("DATABASE_URL")

	if url == "" && env == DEV {
		// TODO: Define the local connection for both the application and the DB container
		url = "postgres://givemegoodcoffee:mypassword@localhost:5432/givemegoodcoffee_dev"
	}

	if url == "" {
		return nil, fmt.Errorf("DATABASE_URL environment variable is required")
	}

	return &DatabaseConfig{url}, nil
}
