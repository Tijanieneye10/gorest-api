package handlers

import (
	"encoding/json"
	"net/http"
)

func healthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := map[string]string{
			"message": "Server is up and running",
		}
		_ = json.NewEncoder(w).Encode(response)
	}
}
