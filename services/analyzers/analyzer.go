package analyzers

import (
	"go-web-analyzer/models"
	"log"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

// Returns an Empty PageAnalysis
func GetEmptyAnalyze() *models.AnalysisResult {
	return &models.AnalysisResult{}
}

// URL Analyzer Process
func Analyze(targetUrl string) *models.AnalysisResult {
	targetUrl = strings.TrimSpace(targetUrl)

	data := models.AnalysisResult{
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

	//Gather page body
	resp, err := http.Get(targetUrl)
	if err != nil {
		log.Println("ERROR: Fetching page failed:", err)
		data.Status = false
		data.StatusMessage = "Fetching page failed"
		return &data
	}
	defer resp.Body.Close() //Close the resource to prevent leakages

	htmlDoc, err := html.Parse(resp.Body) // Read full response body
	if err != nil {
		log.Println("ERROR: Reading page response body failed:", err)
		data.Status = false
		data.StatusMessage = "Reading page response body failed"
		return &data
	}

	//TODO: Actual Page Analyzer logic goes here
	data.HtmlVersion = GetHtmlVersion(htmlDoc)
	data.PageTitle = GetPageTitle(htmlDoc)
	data.Headings = models.Headings{
		NumOfH1: 0,
		NumOfH2: 0,
		NumOfH3: 0,
		NumOfH4: 0,
		NumOfH5: 0,
		NumOfH6: 0,
	}
	data.Links = models.Links{
		NumOfIntLinks:             0,
		NumOfExtLinks:             0,
		NumOfIntLinksInaccessible: 0,
		NumOfExtLinksInaccessible: 0,
	}
	data.HasLoginForm = true
	data.Status = true      //TODO: If anything goes wrong while analyzing, this must be set to false
	data.StatusMessage = "" //TODO: If anything goes wrong while analyzing, an error message should be set here

	return &data
}

// Check whether URL is valid
func isValidUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
