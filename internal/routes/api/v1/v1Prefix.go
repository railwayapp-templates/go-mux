package v1

import (
	"github.com/gorilla/mux"
)

func RegisterPrefix(router *mux.Router) {
	// create the v1 path prefix
	v1Subrouter := router.PathPrefix("/v1/").Subrouter()

	registerRoutes(v1Subrouter)
}
