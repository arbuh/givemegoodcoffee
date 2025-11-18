// Package database contains connections to databases
package database

type Connection interface {
	Close()
}
