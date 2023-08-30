package temperature

import (
	"app/internal/logger"
	"app/internal/responder"
	"net/http"
)

// mock api response
func Handler(w http.ResponseWriter, _ *http.Request) {
	msg := map[string]string{"temp_c": "18.8", "wind_speed_km": "10", "humidity_percent": "70"}

	if err := responder.JSONPretty(w, msg, http.StatusOK); err != nil {
		logger.StderrWithSource.Error(err.Error())
	}
}
