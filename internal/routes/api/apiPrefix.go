package api

import (
	v1 "app/internal/routes/api/v1"

	"github.com/gorilla/mux"
)

func RegisterPrefix(router *mux.Router) {
	// create the api prefix
	apiPrefix := router.PathPrefix("/api/").Subrouter()

	// register the v1 prefix
	v1.RegisterPrefix(apiPrefix)
}
