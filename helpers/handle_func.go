package helpers

import (
	"net/http"

	"moviedl/middlewares"
)

var allMethods = &[]string{http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodConnect, http.MethodOptions, http.MethodTrace}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func CustomHandleFunc(mux *http.ServeMux, url string, controller func(http.ResponseWriter, *http.Request), methods ...string) {
	var handler http.Handler = mux
	handler = middlewares.NotFoundRouter(mux)
	mux.Handle(url, checkMethod(handler, controller, methods...))
}

func checkMethod(next http.Handler, fn func(http.ResponseWriter, *http.Request), methods ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if contains(methods, http.MethodGet) && !contains(methods, http.MethodHead) {
			methods = append(methods, http.MethodHead)
		}
		if len(methods) == 0 {
			methods = *allMethods
		}

		for _, m := range *allMethods {
			for _, n := range methods {
				if m == n && m == r.Method {
					fn(w, r)
					return
				}
			}
		}
		next.ServeHTTP(w, r)
	})
}
