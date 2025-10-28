// Package context contains helpers and constants for HTTP requests context
package context

type contextKey string

const (
	RequestIDKey contextKey = "requestId"
)
