// Package database contains connections to databases
package database

import "context"

type Connection interface {
	Close()
	RunMigrations(ctx context.Context) error
}
