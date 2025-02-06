package analyzers

import (
	"fmt"
	"go-web-analyzer/models"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// GetLinkSummary returns internal, external, inactive link counts for the given htmlDoc
func GetLinkSummary(htmlDoc *html.Node) *models.Links {
	links := &models.Links{}

	var traverse func(*html.Node)

	// Traverse through the node tree and locate all H tags, and add to map with incrementing the count
	traverse = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, attr := range node.Attr {
				if attr.Key == "href" {
					href := attr.Val

					// Ignore empty or # links
					if href == "" || href == "#" {
						return // Skip to the next element
					}

					// Assumes if the href starts with "http://", "https://" or "//" are external links
					if strings.HasPrefix(href, "http://") || strings.HasPrefix(href, "https://") || strings.HasPrefix(href, "//") {
						// External link
						links.NumOfExtLinks++
						if !isUrlAccessible(href) {
							links.NumOfExtLinksInaccessible++
						}
					} else if strings.Contains(href, ":") {
						// All non-hyperlinks are ignored. Ex: ftp://, mailto:
						return // Skip to the next element
					} else {
						// Internal link
						links.NumOfIntLinks++
						if !isUrlAccessible(deriveDirectUrl(href, "TODO_BASE_URL")) {
							links.NumOfIntLinksInaccessible++
						}
					}
				}
			}
		}
		for elem := node.FirstChild; elem != nil; elem = elem.NextSibling {
			traverse(elem)
		}
	}

	traverse(htmlDoc)
	return links
}

// Derive direct URL for relative URLs
func deriveDirectUrl(relativeUrl string, baseUrl string) string {
	//TODO
	return relativeUrl
}

// Check if the URL is accessible
func isUrlAccessible(url string) bool {
	fmt.Println("Check: ", url)
	client := &http.Client{}

	resp, err := client.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	// Check the response status code is 2xx
	if resp.StatusCode/100 == 2 {
		return true
	} else {
		return false
	}
}
