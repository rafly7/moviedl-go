package routers

import (
	"fmt"
	"net/http"

	auth "moviedl/routers/lk21"
)

func (serve *Serve) Index(handler *http.Handler) {
	serve.versionAPI = "v1"
	serve.api(handler)
}

func (serve *Serve) api(next *http.Handler) {
	pathAPI := fmt.Sprintf("/api/%s", serve.versionAPI)
	routeAuth := &auth.RouterLK21{Mux: serve.Mux, AbsPath: pathAPI, Path: "lk21"}
	routeAuth.Index()
}

type Serve struct {
	Mux        *http.ServeMux
	versionAPI string
}
