package handler

import (
	"encoding/json"
	"givemegoodcoffee/internal/http/mapper"
	"net/http"
)

type ErrorHander struct{ errorMapper *mapper.ErrorMapper }

func NewErrorHander() *ErrorHander {
	errorMapper := mapper.NewErrorMapper()
	return &ErrorHander{errorMapper: errorMapper}
}

func (h ErrorHander) WriteClientError(w http.ResponseWriter, errorString string, code int) {
	w.Header().Del("Content-Length")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	response := h.errorMapper.ToErrorResponse(errorString)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		h.WriteServerError(w, err.Error())
	}
}

// TODO: add logging
func (h ErrorHander) WriteServerError(w http.ResponseWriter, errorString string) {
	w.Header().Del("Content-Length")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	// We don't leak details about internal errors
	w.Write([]byte(`{}`))
}
