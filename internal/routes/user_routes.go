package routes

import (
	"net/http"

	"github.com/tijanieneye10/restapi/internal/handlers"
)

func SetupUserRoutes(mux *http.ServeMux, h *handlers.Handler) {
	mux.HandleFunc("POST /register", h.CreateUser())
}
