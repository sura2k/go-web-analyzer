package models

import "golang.org/x/net/html"

// AnalyzerInput struct used wrap analyzer input data
type AnalyzerInput struct {
	TargetUrl string
	BaseUrl   string
	Host      string
	HtmlDoc   *html.Node
}
