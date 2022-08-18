package lk21

import (
	"fmt"
	models "moviedl/models/lk21"
	"net/http"
	"strings"

	"github.com/gocolly/colly/v2"
)

func (s *Collector) Download(slug string) *[]models.LinkBypass {
	arr := &[]models.LinkBypass{}
	c := Scrap()
	c.OnHTML("tbody", func(h *colly.HTMLElement) {
		h.ForEach("tr", func(i int, h *colly.HTMLElement) {
			typeLink := h.ChildText("strong")
			link := h.ChildAttr("a", "href")
			*arr = append(*arr, models.LinkBypass{
				Type: typeLink,
				Link: link,
			})
			fmt.Println(typeLink)
			fmt.Println(link)
			fmt.Println()
		})
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println(r.Method, "data with url:", r.URL.String())
	})
	payload := fmt.Sprintf("slug=%s", slug)
	c.Request("POST", "http://dl.sharemydrive.xyz/verifying.php", strings.NewReader(payload), nil, http.Header{"Content-Type": []string{"application/x-www-form-urlencoded"}})
	return arr
}
