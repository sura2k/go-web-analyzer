package analyzers

import (
	"go-web-analyzer/models"
	"log"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

// LinksAnalyzer struct
type LinksAnalyzer struct{}

// Analyze method updates the relevant field
func (a LinksAnalyzer) Analyze(analyzerInput *models.AnalyzerInput, arm *AnalyzerResultManager) {
	log.Println("LinksAnalyzer: Started")

	links := getLinkDetails(analyzerInput)
	arm.SetLinks(links)

	log.Println("LinksAnalyzer: Completed")
}

// Analyze and return internal, external and inactive link counts
// Assumptions:
//   - Hidden <a> tags are also considered
func getLinkDetails(analyzerInput *models.AnalyzerInput) *models.Links {
	links := &models.Links{}

	var traverse func(*html.Node)

	// Traverse through the node tree and locate all H tags, and add to map with incrementing the count
	traverse = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, attr := range node.Attr {
				if attr.Key == "href" {
					href := attr.Val

					// Ignore # links since they are valid anchors
					if href == "#" {
						return // Skip to the next element
					}

					// Count empty links
					if href == "" {
						links.EmptyLinks.Total++
						return // Skip to the next element
					}

					// Assumes if the href starts with "http://", "https://" or "//" are external links
					if strings.HasPrefix(href, "http://") || strings.HasPrefix(href, "https://") || strings.HasPrefix(href, "//") {
						// External link
						links.External.Total++
						if !isUrlAccessible(href) {
							links.External.Inaccessible++
						}
					} else if strings.Contains(href, ":") {
						// All non-hyperlinks are ignored. Ex: ftp://, mailto:
						links.NonHyperLinks.Total++
						return // Skip to the next element
					} else {
						// Internal link
						links.Internal.Total++
						if !isUrlAccessible(deriveDirectUrl(href, analyzerInput.BaseUrl)) {
							links.Internal.Inaccessible++
						}
					}
				}
			}
		}
		for elem := node.FirstChild; elem != nil; elem = elem.NextSibling {
			traverse(elem)
		}
	}

	traverse(analyzerInput.HtmlDoc)
	return links
}

// Derive direct url for relative urls
func deriveDirectUrl(relativeUrl string, baseUrl string) string {
	parsedBaseUrl, err := url.Parse(baseUrl)
	if err != nil {
		return ""
	}
	// Resolve relative url
	return parsedBaseUrl.ResolveReference(&url.URL{Path: relativeUrl}).String()
}

// Check if the url is accessible - whether returns HTTP 2xx
func isUrlAccessible(url string) bool {
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
