package scrapers

import (
	"log"
	"math/rand"
	"moviedl/utils"
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
)

func Scrap() *colly.Collector {
	randNum := len(utils.UserAgents)
	randAgent := utils.UserAgents[rand.Intn(randNum)]
	c := colly.NewCollector(func(c *colly.Collector) {
		c.UserAgent = randAgent
	})
	return c
}

func findSrcImg(imgSrc string) string {
	reImgSrc, err := regexp.Compile(`([A-Za-z0-9\.\-]+)([a-z\/])`)
	if err != nil {
		log.Panic(err)
	}
	finds := reImgSrc.FindAllString(imgSrc, -1)
	return strings.Join(finds[:], "")
}

func findSlugBypass(url string) string {
	reImgSrc, err := regexp.Compile(`([a-z0-9\-]+)`)
	if err != nil {
		log.Panic(err)
	}
	finds := reImgSrc.FindAllString(url, -1)
	i := len(finds)
	return finds[i-1]
}
