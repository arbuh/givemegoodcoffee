// Package middleware contains middleware for processing HTTP requests
package middleware

import (
	"context"
	"net/http"

	httpctx "givemegoodcoffee/internal/http/context"

	"github.com/google/uuid"
)

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-ID")

		if requestID == "" {
			requestID = uuid.New().String()
		}

		w.Header().Set("X-Request-ID", requestID)

		ctx := context.WithValue(r.Context(), httpctx.RequestIDKey, requestID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
