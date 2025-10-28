package handler

import (
	"encoding/json"
	"givemegoodcoffee/internal/http/mapper"
	"net/http"
	"log/slog"
)

type ErrorHander struct{ errorMapper *mapper.ErrorMapper }

func NewErrorHander() *ErrorHander {
	errorMapper := mapper.NewErrorMapper()
	return &ErrorHander{errorMapper: errorMapper}
}

func (h ErrorHander) HandleClientError(w http.ResponseWriter, errorString string, code int) {
	w.Header().Del("Content-Length")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	response := h.errorMapper.ToErrorResponse(errorString)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		h.HandleServerError(w, err.Error())
	}
}

func (h ErrorHander) HandleServerError(w http.ResponseWriter, errorString string) {
	slog.Error(errorString)

	w.Header().Del("Content-Length")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	// We don't leak details about internal errors
	w.Write([]byte(`{}`))
}
