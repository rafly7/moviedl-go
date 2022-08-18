package middlewares

import (
	"log"
	"net/http"
	"strings"
)

func NotFoundRouter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.Header.Get("X-Forwarded-For")
		if ip == "" {
			ip = strings.Split(r.RemoteAddr, ":")[0]
		}
		log.Printf("[ROUTE-NOT-FOUND] [%s] %s [%s] [%s]", r.Method, r.RequestURI, ip, r.UserAgent())
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Router Not Found"))
	})
}
