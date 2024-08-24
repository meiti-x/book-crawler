package book_parser

import (
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/gocolly/colly/v2"
	digitformatter "github.com/meiti-x/book_crawler/internal/digit_formater"
	"github.com/meiti-x/book_crawler/internal/extract_book_id"
	"github.com/meiti-x/book_crawler/types"
	"strings"
)

// ParseDom is extract required data from data
func ParseDom(e *colly.HTMLElement) *types.Book {
	converter := md.NewConverter("", true, nil)

	title := e.ChildText("h1")
	authorName := e.ChildText("[class^='bookHeader_bookInfo_']:nth-child(1) span:nth-child(2)")
	translator := e.ChildText("[class^='bookHeader_bookInfo_']:nth-child(2) span:nth-child(2) span")
	publication := e.ChildText("[class^='bookHeader_bookInfo_']:nth-child(3) span:nth-child(2) span")
	categories := strings.Split(e.ChildText("[class^='categories_categoriesGroup_']"), "،")
	publishDate := digitformatter.ConvertPersianDigitsToEnglish(e.ChildText("[class^='moreInfo_info']:nth-child(3) p:nth-child(2)"))
	coverImage := e.ChildAttr("img", "src")
	rate := digitformatter.ConvertPersianDigitsToEnglish(e.ChildText("[class^='rate_rate'] span:nth-child(1)"))
	totalRate := digitformatter.ConvertPersianDigitsToEnglish(e.ChildText("[class^='rate_rate'] span:nth-child(2)"))
	bookID, _ := extract_book_id.ExtractIDFromURL(e.Request.URL)

	var descriptionMarkdown string
	e.ForEach("#book-description", func(_ int, el *colly.HTMLElement) {
		rawHTML, er := el.DOM.Html()
		if er == nil {
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
		BookID:      bookID,
		URL:         e.Request.URL.String(),
	}

	return book
}
