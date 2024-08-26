package book

import (
	"fmt"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/gocolly/colly/v2"
	digitformatter "github.com/meiti-x/book_crawler/internal/formater"
	"strings"
)

type Book struct {
	Title       string   `bson:"title"`
	Author      string   `bson:"author"`
	Translator  string   `bson:"translator"`
	Publication string   `bson:"publication"`
	Categories  []string `bson:"categories"`
	Rate        string   `bson:"rate"`
	TotalRate   string   `bson:"total_rate"`
	CoverImage  string   `bson:"cover_image"`
	Description string   `bson:"description"`
	URL         string   `bson:"url"`
	BookID      string   `bson:"book_id"`
}

// ParseDom is extract required data from data
func ParseDom(e *colly.HTMLElement) *Book {
	fmt.Println(e.ChildText("[class^='moreInfo_info']:nth-child(3) [class^='moreInfo_value']"))
	converter := md.NewConverter("", true, nil)

	title := e.ChildText("h1")
	authorName := e.ChildText("[class^='bookHeader_bookInfo_']:nth-child(1) span:nth-child(2)")
	translator := e.ChildText("[class^='bookHeader_bookInfo_']:nth-child(2) span:nth-child(2) span")
	publication := e.ChildText("[class^='bookHeader_bookInfo_']:nth-child(3) span:nth-child(2) span")
	categories := strings.Split(e.ChildText("[class^='categories_categoriesGroup_']"), "،")
	coverImage := e.ChildAttr("img", "src")
	rate := digitformatter.ConvertPersianDigitsToEnglish(e.ChildText("[class^='rate_rate_'] span:nth-child(1)"))
	totalRate := digitformatter.ConvertPersianDigitsToEnglish(e.ChildText("[class^='rate_rate_'] span:nth-child(2)"))
	bookID, _ := ExtractIDFromURL(e.Request.URL)

	var descriptionMarkdown string
	e.ForEach("#book-description", func(_ int, el *colly.HTMLElement) {
		rawHTML, er := el.DOM.Html()
		if er == nil {
			descriptionMarkdown, _ = converter.ConvertString(rawHTML)
		}
	})

	book := &Book{
		Title:       title,
		Author:      authorName,
		Translator:  translator,
		Publication: publication,
		Categories:  categories,
		Rate:        rate,
		TotalRate:   totalRate,
		CoverImage:  coverImage,
		Description: descriptionMarkdown,
		BookID:      bookID,
		URL:         e.Request.URL.String(),
	}

	return book
}
