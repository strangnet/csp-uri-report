package http

import (
	"context"
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type encoder struct {
	logger *log.Logger
}

func NewEncoder(logger *log.Logger) *encoder {
	return &encoder{
		logger: logger,
	}
}

func (e *encoder) Error(ctx context.Context, w http.ResponseWriter, err error) {
	e.logger.Error(err)

	resp := errorResponse{
		Message: err.Error(),
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	switch err {
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(resp)
}

func (e *encoder) Response(w http.ResponseWriter, resp interface{}) error {
	return e.StatusResponse(w, resp, http.StatusOK)
}

func (e *encoder) StatusResponse(w http.ResponseWriter, resp interface{}, statusCode int) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(resp)
}

type errorResponse struct {
	Message string `json:"message"`
}
