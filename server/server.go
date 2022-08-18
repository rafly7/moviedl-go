package server

import (
	"moviedl/middlewares"
	"moviedl/routers"
	"net/http"
)

func HandleServer() (*http.ServeMux, http.Handler) {
	r := http.NewServeMux()
	var handler http.Handler = r
	handler = middlewares.LogRouter(handler)
	s := &routers.Serve{Mux: r}
	s.Index(&handler)
	r.Handle("/", middlewares.NotFoundRouter(r))
	return r, handler
}
