package routes

import (
	"net/http"

	"github.com/tijanieneye10/restapi/internal/handlers"
)

func SetupRoutes(mux *http.ServeMux, handlers *handlers.Handler) {
	SetupHealthRoute(mux, handlers)
	SetupUserRoutes(mux, handlers)
}
