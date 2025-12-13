package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"error,omitempty"`
}

func SuccessResponse(w http.ResponseWriter, message string, Data any, statusCode int) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate")
	w.WriteHeader(statusCode)

	_ = json.NewEncoder(w).Encode(&Response{
		Status:  true,
		Message: message,
		Data:    Data,
	})
}

func ErrorResponse(w http.ResponseWriter, message string, errors any, statusCode int) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate")
	w.WriteHeader(statusCode)

	_ = json.NewEncoder(w).Encode(&Response{
		Status:  false,
		Message: message,
		Errors:  errors,
		Data:    nil,
	})
}
