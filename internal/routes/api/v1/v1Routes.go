package v1

import (
	"app/internal/routes/api/v1/forecast"
	"app/internal/routes/api/v1/temperature"
	"net/http"

	"github.com/gorilla/mux"
)

func registerRoutes(router *mux.Router) {
	// temperature handler
	router.HandleFunc("/temperature", temperature.Handler).Methods(http.MethodGet)

	// forecast handler
	router.HandleFunc("/forecast/{forecastPeriod}", forecast.Handler).Methods(http.MethodGet)
}
