package crawler

import "github.com/gocolly/colly/v2"

// InitializeColly is for initialize web crawler
func InitializeColly() *colly.Collector {
	const crawlingMaxDepth = 2

	c := colly.NewCollector(
		colly.AllowedDomains("taaghche.com"),
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36"),
		colly.MaxDepth(crawlingMaxDepth),
	)
	return c
}
