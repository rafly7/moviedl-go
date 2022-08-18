package scrapers

import (
	"math/rand"
	"moviedl/utils"

	"github.com/gocolly/colly/v2"
)

func Scrap() *colly.Collector {
	randNum := len(utils.UserAgents)
	randAgent := utils.UserAgents[rand.Intn(randNum)]
	return colly.NewCollector(func(c *colly.Collector) {
		c.UserAgent = randAgent
	})
}
