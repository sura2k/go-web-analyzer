package analyzers

import (
	"go-web-analyzer/models"

	"golang.org/x/net/html"
)

// GetHeadingTags returns heading levels and their counts for the given htmlDoc
// Concerns:
//   - Hidden h tags are also calculated
func GetHeadingTags(analyzerInfo *AnalyzerInfo) *models.Headings {
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

	traverse(analyzerInfo.htmlDoc)

	return &models.Headings{
		NumOfH1: headingTags["h1"], // Note: No need to check for nil since Go returns default int value i.e 0 if the key is not available
		NumOfH2: headingTags["h2"],
		NumOfH3: headingTags["h3"],
		NumOfH4: headingTags["h4"],
		NumOfH5: headingTags["h5"],
		NumOfH6: headingTags["h6"],
	}
}
