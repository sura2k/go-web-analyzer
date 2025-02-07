package analyzers

import (
	"go-web-analyzer/models"
	"log"
	"strings"

	"golang.org/x/net/html"
)

// PageTitleAnalyzer struct
type PageTitleAnalyzer struct{}

// Analyze method updates the relevant field
func (a PageTitleAnalyzer) Analyze(analyzerInput *models.AnalyzerInput, arm *AnalyzerResultManager) {
	log.Println("PageTitleAnalyzer: Started")

	pageTitle := getPageTitle(analyzerInput)
	arm.SetPageTitle(pageTitle)

	log.Println("PageTitleAnalyzer: Completed")
}

// Analyze and return the page title - i.e. value of <title>
func getPageTitle(analyzerInput *models.AnalyzerInput) string {
	head := findHeadTag(analyzerInput.HtmlDoc)
	return findTitle(head)
}

// Find the <head> element
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

// Finds the <title> element inside the provided tag (i.e. <head> is expected) and return the <title> value
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
