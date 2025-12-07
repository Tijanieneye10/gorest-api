package routes

import (
	"net/http"

	"github.com/tijanieneye10/restapi/internal/handlers"
)

func SetupHealthRoute(mux *http.ServeMux, handler *handlers.Handler) {
	mux.HandleFunc("GET /health", handler.HealthHandler())
}
