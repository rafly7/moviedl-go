package controllers

import (
	"net/http"

	lk21 "moviedl/scrapers/lk21"
)

type LK21 interface {
	Latest(w http.ResponseWriter, r *http.Request)
	SearchByTitle(w http.ResponseWriter, r *http.Request)
	Popular(w http.ResponseWriter, r *http.Request)
	ListUrlDownload(w http.ResponseWriter, r *http.Request)
}

type Controllers interface {
	LK21
}

func ScrapImpl(scr lk21.ScrapType) *service {
	return &service{
		scrap: scr,
	}
}

type service struct {
	scrap lk21.ScrapType
}
