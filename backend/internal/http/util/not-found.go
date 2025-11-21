// Package util contains helpers for http package
package util

import "net/http"

func WriteNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{}`))
}
