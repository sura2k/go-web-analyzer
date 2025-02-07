package analyzers

import (
	"go-web-analyzer/models"
	"log"

	"golang.org/x/net/html"
)

// HeadingTagsAnalyzer struct
type HeadingTagsAnalyzer struct{}

// Analyze method updates the relevant field
func (a HeadingTagsAnalyzer) Analyze(analyzerInput *models.AnalyzerInput, arm *AnalyzerResultManager) {
	log.Println("HeadingTagsAnalyzer: Started")

	headings := getHeadingTags(analyzerInput)
	arm.SetHeadings(headings)

	log.Println("HeadingTagsAnalyzer: Completed")
}

// Analyze and return heading levels and their counts
// Assumptions:
//   - Hidden h tags are also considered
func getHeadingTags(analyzerInput *models.AnalyzerInput) *models.Headings {
	headingTags := make(map[string]int)

	// Inner recursive functions MUST be declared first since it calls itself within the same context
	var traverse func(*html.Node)

	// Traverse through the node tree and locate all H tags, and add to map with incrementing the count
	traverse = func(node *html.Node) {
		if node.Type == html.ElementNode {
			switch node.Data {
			case "h1", "h2", "h3", "h4", "h5", "h6":
				headingTags[node.Data] = headingTags[node.Data] + 1
			}
		}
		for elem := node.FirstChild; elem != nil; elem = elem.NextSibling {
			traverse(elem)
		}
	}

	traverse(analyzerInput.HtmlDoc)

	return &models.Headings{
		H1Count: headingTags["h1"], // Note: No need to check for nil since Go returns default int value i.e 0 if the key is not available
		H2Count: headingTags["h2"],
		H3Count: headingTags["h3"],
		H4Count: headingTags["h4"],
		H5Count: headingTags["h5"],
		H6Count: headingTags["h6"],
	}
}
