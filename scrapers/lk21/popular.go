package lk21

import (
	models "moviedl/models/lk21"

	"github.com/gocolly/colly/v2"
)

func Popular() *[]models.Popular {
	slice := &[]models.Popular{}
	reSlugBypass, err := reSlugBypass()
	if err != nil {
		return slice
	}
	reImgSrc, err := reImgSrc()
	if err != nil {
		return slice
	}
	c := Scrap()
	c.OnHTML("section.hot-block .slider", func(h *colly.HTMLElement) {
		h.ForEach("figure", func(i int, sc *colly.HTMLElement) {
			href := sc.ChildAttr("a", "href")
			href = findSlugBypass(reSlugBypass, href)
			// fmt.Println(href)

			title := sc.ChildText("a h3")
			// fmt.Println(title)

			rating := sc.ChildText(".rating")
			// fmt.Println(rating)

			imgSrc := sc.ChildAttr("a img", "src")
			imgSrc = findSrcImg(reImgSrc, imgSrc)
			// fmt.Println(imgSrc)
			// fmt.Println()
			*slice = append(*slice, models.Popular{
				Title:      title,
				ImageSrc:   imgSrc,
				SlugBypass: href,
				Rating:     rating,
			})
		})
		// s := h.ChildAttr("img", "src")
		// fmt.Println(s)
	})
	c.Visit("http://149.56.24.226")
	return slice
}
