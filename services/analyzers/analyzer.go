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

	baseUrl, err := getBaseURL(targetUrl)
	if err != nil {
		log.Println("ERROR: Unable to extract the base url:", err)
		data.Status = false
		data.StatusMessage = "Extracting base url failed"
		return &data
	}

	analyzerInfo := &AnalyzerInfo{
		baseUrl: baseUrl,
		htmlDoc: htmlDoc,
	}

	// Populate AnalysisResult
	data.HtmlVersion = GetHtmlVersion(analyzerInfo)
	data.PageTitle = GetPageTitle(analyzerInfo)
	data.Headings = *GetHeadingTags(analyzerInfo) //Rather than letting Go to copy the data, it could be better to use a pointer
	data.Links = *GetLinkSummary(analyzerInfo)
	data.HasLoginForm = HasLoginForm(analyzerInfo)

	return &data
}

// Check whether URL is valid
func isValidUrl(rawURL string) bool {
	u, err := url.Parse(rawURL)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func getBaseURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	// Construct base URL using scheme and host
	baseURL := parsedURL.Scheme + "://" + parsedURL.Host
	return baseURL, nil
}

// AnalyzerInfo struct for passing raw data to various analyzers
type AnalyzerInfo struct {
	baseUrl string
	htmlDoc *html.Node
}
