package analyzers

import (
	"fmt"
	"go-web-analyzer/models"
	"go-web-analyzer/services/utils"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// AnalyzerPreprocessor runs few vaidations and returns AnalyzerInput
type AnalyzerPreprocessor struct {
	targetUrl string
}

// NewAnalyzerPreprocessor initializes the AnalyzerPreprocessor with targetUrl
func NewAnalyzerPreprocessor(targetUrl string) *AnalyzerPreprocessor {
	return &AnalyzerPreprocessor{targetUrl: targetUrl}
}

// ExecutePreprocessor evaluates some basic validations and populate AnalyzerInput
func (aPreproc *AnalyzerPreprocessor) ExecutePreprocessor() (*models.AnalyzerInput, error) {
	log.Println("AnalyzerPreprocessor: Input preprocessing started")

	targetUrl := strings.TrimSpace(aPreproc.targetUrl)

	// If url is not vaid, immediately return with an error message
	_, err := utils.IsValidUrl(targetUrl)
	if err != nil {
		return nil, fmt.Errorf("invalid url. error: %w", err)
	}

	// Execute HTTP GET on the url
	resp, err := http.Get(targetUrl)
	if err != nil {
		return nil, fmt.Errorf("unable to connect with the url. error: %w", err)
	}

	// Close the resource to prevent leakages
	defer resp.Body.Close()

	// Extract response body (i.e. html content) as a node tree
	htmlDoc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("extracting html content failed. error: %w", err)
	}

	baseUrl, err := utils.DeriveBaseUrl(targetUrl)
	if err != nil {
		return nil, fmt.Errorf("deriving base url failed. error: %w", err)
	}

	log.Println("AnalyzerPreprocessor: Input preprocessing completed")

	return &models.AnalyzerInput{
		TargetUrl: targetUrl,
		BaseUrl:   baseUrl,
		HtmlDoc:   htmlDoc,
	}, nil
}
