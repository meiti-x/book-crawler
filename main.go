package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/gocolly/colly/v2"
	"github.com/meiti-x/book_crawler/types"
	"github.com/meiti-x/book_crawler/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// MongoDB setup
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("bookstore").Collection("books")

	converter := md.NewConverter("", true, nil)

	c := colly.NewCollector(
		colly.AllowedDomains("taaghche.com"),
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36"),
		colly.MaxDepth(2),
	)

	visitedURLs := make(map[string]bool)
	c.OnHTML("[class^='bookPage_bookPageContainer']", func(e *colly.HTMLElement) {
		title := e.ChildText("h1")
		authorName := e.ChildText("[class^='bookHeader_bookInfo_']:nth-child(1) span:nth-child(2)")
		translator := e.ChildText("[class^='bookHeader_bookInfo_']:nth-child(2) span:nth-child(2) span")
		publication := e.ChildText("[class^='bookHeader_bookInfo_']:nth-child(3) span:nth-child(2) span")
		categories := strings.Split(e.ChildText("[class^='categories_categoriesGroup_']"), "،")
		publishDate := utils.ConvertPersianDigitsToEnglish(e.ChildText("[class^='more_info_']:nth-child(2) p:nth-child(2)"))
		coverImage := e.ChildAttr("img", "src")
		rate := utils.ConvertPersianDigitsToEnglish(e.ChildText("[class^='rate_rate'] span:nth-child(1)"))
		totalRate := utils.ConvertPersianDigitsToEnglish(e.ChildText("[class^='rate_rate'] span:nth-child(2)"))
		if err != nil {
			fmt.Println("Error:", err)
		}

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
		if err != nil {
			log.Fatalf("Failed to insert document: %v", err)
		}
		fmt.Printf("\n======== Book %s has Been Inserted\n", title)
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.Contains(link, "/book") {
			c.Visit(e.Request.AbsoluteURL(link))
		}
	})

	c.OnRequest(func(r *colly.Request) {
		if visitedURLs[r.URL.String()] {
			r.Abort()
		}
		visitedURLs[r.URL.String()] = true
	})

	c.Visit("https://taaghche.com/")

	fmt.Println("Data written to MongoDB successfully.")
}
