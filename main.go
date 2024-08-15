package main

import (
	"fmt"
	"log"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/gocolly/colly/v2"
	"github.com/meiti-x/book_crawler/utils"
)

func main() {
	converter := md.NewConverter("", true, nil)

	c := colly.NewCollector(
		colly.AllowedDomains("taaghche.com"),
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36"),
	)

	fmt.Println(utils.ConvertPersianDigitsToEnglish("tset 1234"))

	c.OnHTML("#book-description", func(e *colly.HTMLElement) {
		fmt.Print("\n \n")
		fmt.Println("=========")
		// Get the raw HTML of the element
		rawHTML, err := e.DOM.Html()
		if err != nil {
			fmt.Errorf(err.Error())
		}

		// Convert HTML to Markdown
		markdown, err := converter.ConvertString(rawHTML)
		if err != nil {
			log.Fatalf("Error converting HTML to Markdown: %v", err)
		}

		// Print or use the Markdown content
		fmt.Println(markdown)
		fmt.Println("=========")
		fmt.Print("\n \n")
	})

	c.OnHTML("[class^='bookPage_bookPageContainer']", func(e *colly.HTMLElement) {
		title := e.ChildText("h1")
		authorName := e.ChildText("[class^='bookHeader_bookInfo_']:nth-child(1) span:nth-child(2)")
		translator := e.ChildText("[class^='bookHeader_bookInfo_']:nth-child(2) span:nth-child(2) span")
		publication := e.ChildText("[class^='bookHeader_bookInfo_']:nth-child(3) span:nth-child(2) span")
		categories := e.ChildText("[class^='categories_categoriesGroup_']")
		rate := e.ChildText("[class^='rate_rate'] span:nth-child(1)")
		totalRate := e.ChildText("[class^='rate_rate'] span:nth-child(2)")
		// shortDescription := e.ChildText("[class^='rate_rate'] span:nth-child(2)")
		coverImage := e.ChildAttr("img", "src")

		fmt.Print("\n \n")
		fmt.Println("=========")
		fmt.Println("title", title)
		fmt.Println("authorName", authorName)
		fmt.Println("translator", translator)
		fmt.Println("coverImage", coverImage)
		fmt.Println("publication", publication)
		fmt.Println("publication", strings.Split(categories, "،"))
		fmt.Println("rate", utils.ConvertPersianDigitsToEnglish(rate))
		fmt.Println("totalRate", utils.ConvertPersianDigitsToEnglish(totalRate))

		fmt.Println("=========")
		fmt.Print("\n \n")
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.Contains(link, "/book") {
			c.Visit(e.Request.AbsoluteURL(link))
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://taaghche.com/")
}
