package main

import (
	"fmt"
	"log"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/gocolly/colly/v2"
	init_db "github.com/meiti-x/book_crawler/init"
	digit_formatter "github.com/meiti-x/book_crawler/internal"
	"github.com/meiti-x/book_crawler/types"
)

func main() {
	// Initialize MongoDB connection
	client, collection, ctx, err := init_db.InitializeDatabase("mongodb://localhost:27017", "bookstore", "books")
	if err != nil {
		log.Fatal(err)
	}
	defer init_db.DisconnectDatabase(client, ctx)

	converter := md.NewConverter("", true, nil)

	c := colly.NewCollector(
		colly.AllowedDomains("taaghche.com"),
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36"),
		colly.MaxDepth(2),
	)

	visitedURLs := make(map[string]bool)

	// Callback for processing the book page data
	c.OnHTML("[class^='bookPage_bookPageContainer']", func(e *colly.HTMLElement) {
		fmt.Println("Crawling the page data")
		title := e.ChildText("h1")
		authorName := e.ChildText("[class^='bookHeader_bookInfo_']:nth-child(1) span:nth-child(2)")
		translator := e.ChildText("[class^='bookHeader_bookInfo_']:nth-child(2) span:nth-child(2) span")
		publication := e.ChildText("[class^='bookHeader_bookInfo_']:nth-child(3) span:nth-child(2) span")
		categories := strings.Split(e.ChildText("[class^='categories_categoriesGroup_']"), "،")
		publishDate := digit_formatter.ConvertPersianDigitsToEnglish(e.ChildText("[class^='more_info_']:nth-child(2) p:nth-child(2)"))
		coverImage := e.ChildAttr("img", "src")
		rate := digit_formatter.ConvertPersianDigitsToEnglish(e.ChildText("[class^='rate_rate'] span:nth-child(1)"))
		totalRate := digit_formatter.ConvertPersianDigitsToEnglish(e.ChildText("[class^='rate_rate'] span:nth-child(2)"))

		var descriptionMarkdown string
		e.ForEach("#book-description", func(_ int, el *colly.HTMLElement) {
			rawHTML, err := el.DOM.Html()
			if err == nil {
				descriptionMarkdown, _ = converter.ConvertString(rawHTML)
			}
		})

		book := &types.Book{
			Title:       title,
			Author:      authorName,
			Translator:  translator,
			Publication: publication,
			Categories:  categories,
			Rate:        rate,
			TotalRate:   totalRate,
			CoverImage:  coverImage,
			Description: descriptionMarkdown,
			PublishDate: publishDate,
		}

		_, err := collection.InsertOne(ctx, book)
		fmt.Println(collection)
		if err != nil {
			log.Fatalf("Failed to insert document: %v", err)
		}
		fmt.Printf("======== Book %s has Been Inserted\n", title)
	})

	// Callback for finding links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))

		if strings.Contains(link, "/book") {
			if !visitedURLs[link] {
				visitedURLs[link] = true
				c.Visit(link)
			}
		}
	})

	// Callback for requests
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL.String())
	})

	c.Visit("https://taaghche.com/")

	fmt.Println("Data written to MongoDB successfully.")
}
