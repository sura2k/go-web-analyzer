package analyzers

import (
	"go-web-analyzer/models"
	"net/url"
)

// Local aliases for simplicity, otherwise models.PageAnalysis etc should be used
type PageAnalysis = models.PageAnalysis
type Headings = models.Headings
type Links = models.Links

// Returns and Empty PageAnalysis
func GetEmptyAnalyze() *PageAnalysis {
	return &PageAnalysis{}
}

// URL Analyzer Process
func Analyze(targetUrl string) *PageAnalysis {
	data := PageAnalysis{
		TargetUrl:     targetUrl,
		Status:        true,
		StatusMessage: "",
	}

	// If URL is not vaid, immediately returns with an error message
	if !isValidUrl(targetUrl) {
		data.Status = false
		data.StatusMessage = "Provided URL is invalid"
		return &data
	}

	//TODO: Actual Page Analyzer logic goes here
	data.HtmlVersion = "HTML5"
	data.PageTitle = "Demo"
	data.Headings = Headings{
		NumOfH1: 0,
		NumOfH2: 0,
		NumOfH3: 0,
		NumOfH4: 0,
		NumOfH5: 0,
		NumOfH6: 0,
	}
	data.Links = Links{
		NumOfIntLinks:             0,
		NumOfExtLinks:             0,
		NumOfIntLinksInaccessible: 0,
		NumOfExtLinksInaccessible: 0,
	}
	data.HasLoginForm = true
	data.RedirectsToLogin = false
	data.Status = true      //TODO: If anything goes wrong while analyzing, this must be set to false
	data.StatusMessage = "" //TODO: If anything goes wrong while analyzing, an error message should be set here

	return &data
}

// Check whether URL is valid
func isValidUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
