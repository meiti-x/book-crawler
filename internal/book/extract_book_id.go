package book

import (
	"fmt"
	"net/url"
	"strings"
)

// ExtractIDFromURL is only return the book id in the url
func ExtractIDFromURL(inputURL *url.URL) (string, error) {
	// Parse the URL
	parsedURL, err := url.Parse(inputURL.String())
	if err != nil {
		return "", err
	}

	// Split the path and get the part that contains the ID
	pathSegments := strings.Split(parsedURL.Path, "/")

	// Check if we have at least two parts in the path
	if len(pathSegments) < 2 {
		return "", fmt.Errorf("URL path doesn't contain an ID")
	}

	// The ID should be the second last part (assuming a consistent URL structure)
	id := pathSegments[len(pathSegments)-2]

	return id, nil
}
