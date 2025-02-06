package analyzers

import (
	"golang.org/x/net/html"
)

// GetHeadingTags returns a map of all heading levels availabes and counts for the given htmlDoc
func GetHeadingTags(htmlDoc *html.Node) map[string]int {
	headings := make(map[string]int)

	// Inner recursive functions MUST be declared first since it calls itself within the same context
	var traverse func(*html.Node)

	// Traverse through the node tree and locate all H tags, and add to map with incrementing the count
	traverse = func(node *html.Node) {
		if node.Type == html.ElementNode {
			switch node.Data {
			case "h1", "h2", "h3", "h4", "h5", "h6":
				headings[node.Data] = headings[node.Data] + 1
			}
		}
		for elem := node.FirstChild; elem != nil; elem = elem.NextSibling {
			traverse(elem)
		}
	}

	traverse(htmlDoc)
	return headings
}
