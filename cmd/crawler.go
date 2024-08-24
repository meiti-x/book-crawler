package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/meiti-x/book_crawler/init/crawler"
	initdb "github.com/meiti-x/book_crawler/init/db"
	"github.com/meiti-x/book_crawler/internal/book_parser"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"strings"

	"github.com/gocolly/colly/v2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	client, collection, err := initdb.InitializeDatabase()
	ctx := context.Background()
	if err != nil {
		log.Fatal(err)
	}
	defer initdb.DisconnectDatabase(ctx, client)

	c := crawler.InitializeColly()

	visitedURLs := make(map[string]bool)

	c.OnHTML("[class^='bookPage_bookPageContainer']", func(e *colly.HTMLElement) {
		book := book_parser.ParseDom(e)

		filter := bson.M{"book_id": book.BookID}
		count, err := collection.CountDocuments(ctx, filter)
		if err != nil {
			log.Fatalf("Failed to check if document exists: %v", err)
		}

		if count == 0 {
			_, err := collection.InsertOne(ctx, book)
			if err != nil {
				log.Fatalf("Failed to insert document: %v", err)
			}
			fmt.Printf("======== Book %s has Been Inserted\n", book.Title)
		} else {
			fmt.Printf("======== Book %s already exists in the database\n", book.Title)
		}
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

	c.Visit("https://taaghche.com/")

}
