package analyzers

import (
	"golang.org/x/net/html"
)

// GetHtmlVersion returns the HTML version in the given page body
func GetHtmlVersion(htmlDoc *html.Node) string {
	//TODO
	return "HTML5"
}
