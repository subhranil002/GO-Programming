package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

// RequestLogger is a middleware that logs incoming HTTP requests to the terminal.
func RequestLogger(log *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Process the request
			next.ServeHTTP(w, r)

			// Log request details after processing
			log.Info("request completed",
				"method", r.Method,
				"path", r.URL.Path,
				"duration", time.Since(start).String(),
			)
		})
	}
}
