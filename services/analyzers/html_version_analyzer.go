package analyzers

import (
	"log"
	"strings"

	"github.com/sura2k/go-web-analyzer/models"

	"golang.org/x/net/html"
)

// HtmlVersionAnalyzer struct
type HtmlVersionAnalyzer struct{}

// Analyze method updates the relevant field
func (a HtmlVersionAnalyzer) Analyze(analyzerInput *models.AnalyzerInput, arm *AnalyzerResultManager) {
	log.Println("HtmlVersionAnalyzer: Started")

	htmlVersion := getHtmlVersion(analyzerInput)
	arm.SetHtmlVersion(htmlVersion)

	log.Println("HtmlVersionAnalyzer: Completed")
}

// Analyze and returns the HTML version using <!DOCTYPE> element
// Assumptions
//   - Assumes as HTML5, if no <!DOCTYPE> element is presented
func getHtmlVersion(analyzerInput *models.AnalyzerInput) string {
	doctype := analyzerInput.HtmlDoc.FirstChild
	if doctype != nil && doctype.Type == html.DoctypeNode && strings.ToUpper(doctype.Data) == "HTML" {
		for _, attr := range doctype.Attr {
			if strings.ToUpper(attr.Key) == "PUBLIC" {
				switch attr.Val {
				case "-//W3C//DTD HTML 4.01//EN":
					return "HTML 4.01 Strict"
				case "-//W3C//DTD HTML 4.01 Transitional//EN":
					return "HTML 4.01 Transitional"
				case "-//W3C//DTD HTML 4.01 Frameset//EN":
					return "HTML 4.01 Frameset"
				case "-//W3C//DTD XHTML 1.0 Strict//EN":
					return "XHTML 1.0 Strict"
				case "-//W3C//DTD XHTML 1.0 Transitional//EN":
					return "XHTML 1.0 Transitional"
				case "-//W3C//DTD XHTML 1.0 Frameset//EN":
					return "XHTML 1.0 Frameset"
				case "-//W3C//DTD XHTML 1.1//EN":
					return "XHTML 1.1"
				default:
					return "Unknown Version: " + attr.Val
				}
			}
		}

		//Assumes as HTML5, if no attributes are presented
		return "HTML5"
	}

	//Assumes as HTML5, if no <!DOCTYPE> element is presented
	return "HTML5"
}
