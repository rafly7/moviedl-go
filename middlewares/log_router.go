package middlewares

import (
	"log"
	"net/http"
	"strings"

	"moviedl/configs"
)

func LogRouter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.Header.Get("X-Forwarded-For")
		if ip == "" {
			ip = strings.Split(r.RemoteAddr, ":")[0]
		}
		if configs.LoadEnv("LOG_LEVEL") == "info" {
			log.Printf("[INFO-ROUTE] [%s] %s", ip, r.UserAgent())
		}
		next.ServeHTTP(w, r)
	})
}
