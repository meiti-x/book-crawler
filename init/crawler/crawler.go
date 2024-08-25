package crawler

import (
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
)

// InitializeColly is for initialize web crawler
func InitializeColly() *colly.Collector {
	const crawlingMaxDepth = 2
	const threadCount = 3

	c := colly.NewCollector(
		colly.AllowedDomains("taaghche.com"),
		colly.MaxDepth(crawlingMaxDepth),
		colly.Async(true),
	)

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: threadCount,
	})

	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	return c
}
