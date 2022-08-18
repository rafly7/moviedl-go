package lk21

import (
	"fmt"
	models "moviedl/models/lk21"
	"strings"

	"github.com/gocolly/colly/v2"
)

func (s *Collector) SearchByTitle(search string) *[]models.Movie {
	// var movies = []models.Movie{}

	arr := &[]models.Movie{}
	reGenre, err := reGenre()
	if err != nil {
		return arr
	}
	reGenreStr, err := reGenreStr()
	if err != nil {
		return arr
	}
	reSlugBypass, err := reSlugBypass()
	if err != nil {
		return arr
	}
	reImgSrc, err := reImgSrc()
	if err != nil {
		return arr
	}
	/**
	 */

	c := Scrap()
	c.OnHTML("div.search-item", func(h *colly.HTMLElement) {

		imgSrc := h.ChildAttr(".search-poster img", "src")
		imgSrc = findSrcImg(reImgSrc, imgSrc)
		fmt.Println(imgSrc)

		href := h.ChildAttr(".search-poster a", "href")
		href = findSlugBypass(reSlugBypass, href)
		fmt.Println(href)

		title := h.ChildAttr(".search-poster a", "title")
		fmt.Println(title)

		other := h.ChildAttrs(".search-content .cat-links a", "href")
		genres := &[]string{}
		for _, val := range other {
			if strings.HasPrefix(val, "/genre") {
				genre := reGenreStr.FindString(reGenre.FindString(val))
				*genres = append(*genres, genre)
			}
		}
		fmt.Println(other)

		fmt.Println()
		*arr = append(*arr, models.Movie{
			Title:      title,
			ImageSrc:   imgSrc,
			SlugBypass: href,
			Genres:     *genres,
		})
	})
	url := fmt.Sprintf("http://149.56.24.226?s=%s", search)
	c.Visit(url)
	return arr
}
