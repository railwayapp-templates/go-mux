package middleware

import (
	"app/internal/logger"
	"net/http"

	"github.com/felixge/httpsnoop"
	"golang.org/x/exp/slog"
)

// logger middleware for access logs
func Logger() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// gathers metrics from the upstream handlers
			metrics := httpsnoop.CaptureMetrics(h, w, r)

			//prints log and metrics
			logger.Stdout.Info(
				"access_log",
				slog.String("method", r.Method),
				slog.String("uri", r.URL.RequestURI()),
				slog.String("user_agent", r.Header.Get("User-Agent")),
				slog.String("ip", r.RemoteAddr),
				slog.Int("code", metrics.Code),
				slog.Int64("bytes", metrics.Written),
				slog.Duration("request_time", metrics.Duration),
			)
		})
	}
}
