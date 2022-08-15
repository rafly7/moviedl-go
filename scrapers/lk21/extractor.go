package lk21

import (
	"log"
	"moviedl/scrapers"
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
)

func Scrap() *colly.Collector {
	return scrapers.Scrap()
}

func reImgSrc() (*regexp.Regexp, error) {
	reImgSrc, err := regexp.Compile(`([A-Za-z0-9\.\-]+)([a-z\/])`)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return reImgSrc, nil
}

func reSlugBypass() (*regexp.Regexp, error) {
	reSlug, err := regexp.Compile(`([a-z0-9\-]+)`)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return reSlug, nil
}

func reGenre() (*regexp.Regexp, error) {
	reGenre, err := regexp.Compile(`([a-z]+)([\/]$)`)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return reGenre, nil
}

func reGenreStr() (*regexp.Regexp, error) {
	reGenreStr, err := regexp.Compile(`([a-z]+)`)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return reGenreStr, nil
}

func findSrcImg(regex *regexp.Regexp, imgSrc string) string {
	finds := regex.FindAllString(imgSrc, -1)
	return strings.Join(finds[:], "")
}

func findSlugBypass(regex *regexp.Regexp, url string) string {
	finds := regex.FindAllString(url, -1)
	i := len(finds)
	return finds[i-1]
}
