package analyzers

import (
	"go-web-analyzer/models"
)

// Starts the Analyzing Process
func StartAnalyzer(targetUrl string) (*models.AnalyzerResult, error) {

	analyzerInput, err := PreprocessInput(targetUrl)
	if err != nil {
		return nil, err
	}

	// Populate AnalyzerResult fields
	// Note: Used pointers for child structs so that Go doesnt have to copy such structs
	analyzerResult := models.AnalyzerResult{
		TargetUrl:    analyzerInput.TargetUrl,
		HtmlVersion:  GetHtmlVersion(analyzerInput),
		PageTitle:    GetPageTitle(analyzerInput),
		Headings:     *GetHeadingTags(analyzerInput),
		Links:        *GetLinkSummary(analyzerInput),
		HasLoginForm: HasLoginForm(analyzerInput),
	}

	return &analyzerResult, nil
}
