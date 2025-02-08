package analyzers

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/chromedp/chromedp"
	"github.com/sura2k/go-web-analyzer/models"
	"github.com/sura2k/go-web-analyzer/services/utils"

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

	// Connect to the url using chromedp
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var doctype, htmlContent string

	err = chromedp.Run(ctx,
		chromedp.Navigate(targetUrl),
		chromedp.Evaluate(`document.doctype ? new XMLSerializer().serializeToString(document.doctype) : ''`, &doctype), // Fetch DOCTYPE seprately as OuterHTML does not return DOCTYPE
		chromedp.OuterHTML("html", &htmlContent),
	)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to the given url. error: %w", err)
	}

	// Parse html content into a node tree
	htmlDoc, err := html.Parse(strings.NewReader(doctype + htmlContent)) //DOCTYPE is amended to htmlContent
	if err != nil {
		return nil, fmt.Errorf("parsing html content failed. error: %w", err)
	}

	// Derive base url
	baseUrl, err := utils.DeriveBaseUrl(targetUrl)
	if err != nil {
		return nil, fmt.Errorf("deriving base url failed. error: %w", err)
	}

	host, err := utils.DeriveHost(targetUrl)
	if err != nil {
		return nil, fmt.Errorf("deriving host failed. error: %w", err)
	}

	log.Println("AnalyzerPreprocessor: Input preprocessing completed")

	return &models.AnalyzerInput{
		TargetUrl: targetUrl,
		BaseUrl:   baseUrl,
		Host:      host,
		HtmlDoc:   htmlDoc,
	}, nil
}
