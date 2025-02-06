package analyzers

import (
	"strings"

	"golang.org/x/net/html"
)

// GetPageTitle returns the page title i.e. value of <title> for the given htmlDoc
func GetPageTitle(htmlDoc *html.Node) string {
	head := findHeadTag(htmlDoc)
	return findTitle(head)
}

// Finds the <head> element in the HTML document
func findHeadTag(htmlDoc *html.Node) *html.Node {
	if htmlDoc.Type == html.ElementNode && htmlDoc.Data == "head" {
		return htmlDoc
	}
	for elem := htmlDoc.FirstChild; elem != nil; elem = elem.NextSibling {
		head := findHeadTag(elem)
		if head != nil {
			return head
		}
	}
	return nil
}

// Finds the <title> element inside the <head> tag and return the value
func findTitle(head *html.Node) string {
	if head == nil {
		return ""
	}
	for elem := head.FirstChild; elem != nil; elem = elem.NextSibling {
		if elem.Type == html.ElementNode && elem.Data == "title" && elem.FirstChild != nil {
			return strings.TrimSpace(elem.FirstChild.Data)
		}
	}
	return ""
}
