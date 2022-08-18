package controllers

import (
	"moviedl/helpers"
	"moviedl/utils"
	"net/http"
	"regexp"
	"strconv"
)

func (c *service) Latest(w http.ResponseWriter, r *http.Request) {
	reNum := regexp.MustCompile(`^[1-9]+$`)
	page := r.URL.Query().Get("page")
	if reNum.MatchString(page) {
		page, err := strconv.Atoi(page)
		if err != nil {
			helpers.Response(w, utils.M{"message": "Something went wrong"}, http.StatusInternalServerError, nil)
			return
		}
		res := c.scrap.Latest(page)
		if len(*res) > 0 {
			helpers.Response(w, utils.M{"results": *res}, http.StatusOK, nil)
			return
		}
	}
	helpers.Response(w, utils.M{"message": "Failed"}, http.StatusBadRequest, nil)
}

func (c *service) SearchByTitle(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	if title != "" {
		res := c.scrap.SearchByTitle(title)
		if len(*res) > 0 {
			helpers.Response(w, utils.M{"results": *res}, http.StatusOK, nil)
			return
		}
	}
	helpers.Response(w, utils.M{"message": "Failed"}, http.StatusBadRequest, nil)
}

func (c *service) ListUrlDownload(w http.ResponseWriter, r *http.Request) {
	slug := r.URL.Query().Get("name")
	if slug != "" {
		res := c.scrap.Download(slug)
		if len(*res) > 0 {
			helpers.Response(w, utils.M{"results": *res}, http.StatusOK, nil)
			return
		}
	}
	helpers.Response(w, utils.M{"message": "Failed"}, http.StatusBadRequest, nil)
}

func (c *service) Popular(w http.ResponseWriter, r *http.Request) {
	res := c.scrap.Popular()
	if len(*res) > 0 {
		helpers.Response(w, utils.M{"results": *res}, http.StatusOK, nil)
		return
	}
	helpers.Response(w, utils.M{"message": "Failed"}, http.StatusBadRequest, nil)
}
