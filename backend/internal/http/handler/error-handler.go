package handler

import (
	"encoding/json"
	httpctx "givemegoodcoffee/internal/http/context"
	"givemegoodcoffee/internal/http/mapper"
	"log/slog"
	"net/http"
)

type ErrorHander struct{ errorMapper *mapper.ErrorMapper }

func NewErrorHander() *ErrorHander {
	errorMapper := mapper.NewErrorMapper()
	return &ErrorHander{errorMapper: errorMapper}
}

func (h ErrorHander) HandleClientError(w http.ResponseWriter, r *http.Request, errorString string, code int) {
	w.Header().Del("Content-Length")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	response := h.errorMapper.ToErrorResponse(errorString)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		h.HandleServerError(w, r, err.Error())
	}
}

func (h ErrorHander) HandleServerError(w http.ResponseWriter, r *http.Request, errorString string) {
	ctx := r.Context()
	requestID := ctx.Value(httpctx.RequestIDKey)

	// TODO: use structural logging when we run the application in a server
	slog.Error(errorString, slog.Any("requestID", requestID))

	w.Header().Del("Content-Length")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	// We don't leak details about internal errors
	w.Write([]byte(`{}`))
}
