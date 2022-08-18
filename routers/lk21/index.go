package auth

import (
	"fmt"
	"net/http"

	"moviedl/controllers"
	"moviedl/helpers"
	scrap "moviedl/scrapers/lk21"
)

func (serve *RouterLK21) Index() {
	path := fmt.Sprintf("%s/%s", serve.AbsPath, serve.Path)

	var scrapLK21 scrap.LK21 = scrap.CollectorImpl()

	var controller controllers.LK21 = controllers.ScrapImpl(scrapLK21)

	latest := fmt.Sprintf("%s/latest", path)
	helpers.CustomHandleFunc(serve.Mux, latest, controller.Latest, http.MethodGet)

	popular := fmt.Sprintf("%s/popular", path)
	helpers.CustomHandleFunc(serve.Mux, popular, controller.Popular, http.MethodGet)

	search := fmt.Sprintf("%s/search", path)
	helpers.CustomHandleFunc(serve.Mux, search, controller.SearchByTitle, http.MethodGet)

	listDownload := fmt.Sprintf("%s/list-download", path)
	helpers.CustomHandleFunc(serve.Mux, listDownload, controller.ListUrlDownload, http.MethodGet)
}

type RouterLK21 struct {
	Mux     *http.ServeMux
	AbsPath string
	Path    string
}
