package analyzers

import (
	"go-web-analyzer/models"
	"log"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

// Starts the Analyzing Process
func StartAnalyzer(targetUrl string) *models.AnalyzerResult {
	targetUrl = strings.TrimSpace(targetUrl)

	analyzerResult := models.AnalyzerResult{
		TargetUrl:     targetUrl,
		Status:        true,
		StatusMessage: "",
	}

	// If URL is not vaid, immediately return with an error message
	_, err := IsValidUrl(targetUrl)
	if err != nil {
		log.Println("ERROR: Provided URL is invalid:", err)
		analyzerResult.Status = false
		analyzerResult.StatusMessage = "Provided URL is invalid"
		return &analyzerResult
	}

	//Gather page body
	resp, err := http.Get(targetUrl)
	if err != nil {
		log.Println("ERROR: Fetching page failed:", err)
		analyzerResult.Status = false
		analyzerResult.StatusMessage = "Fetching page failed"
		return &analyzerResult
	}
	defer resp.Body.Close() //Close the resource to prevent leakages

	htmlDoc, err := html.Parse(resp.Body) // Read full response body
	if err != nil {
		log.Println("ERROR: Reading page response body failed:", err)
		analyzerResult.Status = false
		analyzerResult.StatusMessage = "Reading page response body failed"
		return &analyzerResult
	}

	baseUrl, err := DeriveBaseUrl(targetUrl)
	if err != nil {
		log.Println("ERROR: Unable to extract the base url:", err)
		analyzerResult.Status = false
		analyzerResult.StatusMessage = "Extracting base url failed"
		return &analyzerResult
	}

	analyzerInput := &models.AnalyzerInput{
		TargetUrl: targetUrl,
		BaseUrl:   baseUrl,
		HtmlDoc:   htmlDoc,
	}

	// Populate AnalyzerResult fields
	// Note: Used pointers for child structs so that Go doesnt have to copy such structs
	analyzerResult.HtmlVersion = GetHtmlVersion(analyzerInput)
	analyzerResult.PageTitle = GetPageTitle(analyzerInput)
	analyzerResult.Headings = *GetHeadingTags(analyzerInput)
	analyzerResult.Links = *GetLinkSummary(analyzerInput)
	analyzerResult.HasLoginForm = HasLoginForm(analyzerInput)

	return &analyzerResult
}

// Check whether the url is valid
func IsValidUrl(rawUrl string) (bool, error) {
	u, err := url.Parse(rawUrl)

	if err != nil {
		return false, err
	}

	return u.Scheme != "" && u.Host != "", nil
}

// Derive base url from the given url
func DeriveBaseUrl(rawUrl string) (string, error) {
	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}

	// Construct base url using scheme and host
	baseUrl := parsedUrl.Scheme + "://" + parsedUrl.Host
	return baseUrl, nil
}
