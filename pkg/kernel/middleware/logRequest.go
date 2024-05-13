package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
)

func WithLogging(next http.Handler, logger *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info(fmt.Sprintf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI()))

		next.ServeHTTP(w, r)
	})
}
