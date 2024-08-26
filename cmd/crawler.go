package main

import (
	"context"
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/joho/godotenv"
	"github.com/meiti-x/book_crawler/init/crawler"
	"github.com/meiti-x/book_crawler/init/db"
	"github.com/meiti-x/book_crawler/internal/book"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"strings"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client, collection, err := db.InitializeDatabase()

	ctx := context.Background()

	if err != nil {
		log.Fatal(err)
	}

	defer db.DisconnectDatabase(ctx, client)

	err = db.CreateBookIDIndex(ctx, collection)
	if err != nil {
		log.Fatalf("Failed to create index: %v", err)
	}
	c := crawler.InitializeColly()

	visitedURLs := make(map[string]bool)

	c.OnHTML("[class^='bookPage_bookPageContainer']", func(e *colly.HTMLElement) {
		book := book.ParseDom(e)

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
	c.Wait()

}
