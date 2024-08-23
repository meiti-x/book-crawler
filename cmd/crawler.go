package main

import (
	"context"
	"fmt"
	"github.com/meiti-x/book_crawler/init/crawler"
	initdb "github.com/meiti-x/book_crawler/init/db"
	"github.com/meiti-x/book_crawler/internal/book_parser"
	"log"
	"strings"

	"github.com/gocolly/colly/v2"
)

func main() {
	client, collection, err := initdb.InitializeDatabase("mongodb://localhost:27017", "bookstore", "books")
	ctx := context.Background()
	if err != nil {
		log.Fatal(err)
	}
	defer initdb.DisconnectDatabase(ctx, client)

	c := crawler.InitializeColly()

	visitedURLs := make(map[string]bool)

	// Callback for processing the book page data
	c.OnHTML("[class^='bookPage_bookPageContainer']", func(e *colly.HTMLElement) {
		book := book_parser.ParseDom(e)

		_, err := collection.InsertOne(ctx, book)
		if err != nil {
			log.Fatalf("Failed to insert document: %v", err)
		}
		fmt.Sprintf("======== Book %s has Been Inserted\n", book.Title)
	})

	// Callback for finding links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))

		if strings.Contains(link, "/book") {
			if !visitedURLs[link] {
				visitedURLs[link] = true
				err := c.Visit(link)
				if err != nil {
					return
				}
			}
		}
	})
	
	err = c.Visit("https://taaghche.com/")
	if err != nil {
		return
	}

}
