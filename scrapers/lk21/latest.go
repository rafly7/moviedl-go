package lk21

import (
	"fmt"
	"log"
	models "moviedl/models/lk21"
	"strings"

	"github.com/gocolly/colly/v2"
)

func (s *Collector) Latest(page int) *[]models.Latest {
	arr := &[]models.Latest{}

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
	c := Scrap()
	c.OnHTML("section .grid-archive #grid-wrapper", func(h *colly.HTMLElement) {
		// fmt.Println(h.Text)
		h.ForEach(".mega-item", func(i int, h *colly.HTMLElement) {
			href := h.ChildAttr("figure a", "href")
			href = findSlugBypass(reSlugBypass, href)
			// fmt.Println(href)

			imgSrc := h.ChildAttr("figure a img", "src")
			imgSrc = findSrcImg(reImgSrc, imgSrc)
			// fmt.Println(imgSrc)

			rating := h.ChildText(".rating")
			// fmt.Println(rating)

			title := h.ChildText(".grid-header h1 a")
			spt := strings.Split(title, " Film Subtitle Indonesia Streaming Movie Download")
			title = strings.Split(spt[0], "Nonton ")[1]
			// fmt.Println(title)

			genres := &[]string{}
			h.ForEach(".grid-categories", func(i int, sc *colly.HTMLElement) {
				ca := sc.ChildAttrs("a", "href")
				for _, val := range ca {
					if strings.HasPrefix(val, "/genre") {
						genre := reGenreStr.FindString(reGenre.FindString(val))
						*genres = append(*genres, genre)
					}
				}
				// fmt.Println(ca)
			})
			// fmt.Println(*genres)
			*arr = append(*arr, models.Latest{
				Title:      title,
				ImageSrc:   imgSrc,
				SlugBypass: href,
				Rating:     rating,
				Genres:     *genres,
			})
		})
	})
	c.OnError(func(r *colly.Response, err error) {
		log.Print(err)
	})
	c.OnRequest(func(r *colly.Request) {
		log.Println(r.URL.String())
	})
	url := fmt.Sprintf("https://lk21.コム/latest/page/%d", page)
	c.Visit(url)
	// c.Wait()
	return arr
}
