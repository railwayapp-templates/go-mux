package base

import (
	"app/internal/routes/base/health"
	"app/internal/routes/base/root"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	baseRouter := router.NewRoute().Subrouter()

	// root handler
	baseRouter.HandleFunc("/", root.Handler).Methods(http.MethodGet)

	// health handler
	baseRouter.HandleFunc("/health", health.Handler).Methods(http.MethodGet, http.MethodHead)
}
