package analyzers

import (
	"go-web-analyzer/models"
	"go-web-analyzer/services/utils"
	"log"
	"strings"
	"sync"

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
	linkMap := make(map[string]bool) // For tracking purpose

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
						linkMap[href] = true
					} else if strings.Contains(href, ":") {
						// All non-hyperlinks are ignored. Ex: ftp://, mailto:
						links.NonHyperLinks.Total++
						return // Skip to the next element
					} else {
						// Internal link
						links.Internal.Total++
						directUrl := utils.DeriveDirectUrl(href, analyzerInput.BaseUrl)
						linkMap[directUrl] = false
					}
				}
			}
		}
		for elem := node.FirstChild; elem != nil; elem = elem.NextSibling {
			traverse(elem)
		}
	}

	traverse(analyzerInput.HtmlDoc)

	// Starts the link health checker goroutines
	startLinkHealthChecker(linkMap, links)

	return links
}

// Check the accessbility of all the link parallaly and update the inaccessible counts
//
// Major Concern:
//
//	Since this implementation start all the link processing parally
//	it leads to uncontrolled goroutine creation and then it will eatup the CPU and RAM
//	which eventually crashes your application
//
// Solution:
//
//	Execute links in batches
func startLinkHealthChecker(linkMap map[string]bool, links *models.Links) {
	log.Println("LinkHealthChecker: Started. numOfLinks: ", len(linkMap))

	var lock sync.Mutex
	var wg sync.WaitGroup

	// Start goroutines
	for key, val := range linkMap {
		wg.Add(1) // IMPORTANT: No issues even if increment the task count here since wg.Wait() will not be executed till the loop is completed
		go func(url string, isExternal bool) {
			defer wg.Done()

			// Check accessibility of the url
			// Note:
			//	This takes considerable amount of time to return and
			//	thats why additional goroutine is used to increase performance
			isAccessible := utils.IsUrlAccessible(url)

			// Lock
			lock.Lock()

			// Update the unsafe shared Links struct
			if !isAccessible {
				if isExternal {
					links.External.Inaccessible++
				} else {
					links.Internal.Inaccessible++
				}
			}

			// Unlock
			lock.Unlock()
		}(key, val)
	}

	// Blocks the main thread until all goroutines are completed
	wg.Wait()

	log.Println("LinkHealthChecker: Completed")
}
