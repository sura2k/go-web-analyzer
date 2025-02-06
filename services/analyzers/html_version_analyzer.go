package analyzers

import (
	"strings"

	"golang.org/x/net/html"
)

// GetHtmlVersion returns the HTML version for the given htmlDoc
func GetHtmlVersion(analyzerInfo *AnalyzerInfo) string {
	doctype := analyzerInfo.htmlDoc.FirstChild
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

		//Assumes as HTML5 if no attributes are presented
		return "HTML5"
	}

	return "Not Found"
}
